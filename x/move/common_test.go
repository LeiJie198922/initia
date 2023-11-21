package move_test

import (
	"bytes"
	"encoding/binary"
	"testing"

	initiaapp "github.com/initia-labs/initia/app"
	customdistrtypes "github.com/initia-labs/initia/x/distribution/types"
	"github.com/initia-labs/initia/x/move/types"
	stakingtypes "github.com/initia-labs/initia/x/mstaking/types"
	vmtypes "github.com/initia-labs/initiavm/types"

	"github.com/stretchr/testify/require"

	"cosmossdk.io/math"

	"github.com/cosmos/cosmos-sdk/crypto/keys/ed25519"
	"github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	cryptotypes "github.com/cosmos/cosmos-sdk/crypto/types"
	testutilsims "github.com/cosmos/cosmos-sdk/testutil/sims"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"

	abci "github.com/cometbft/cometbft/abci/types"
	tmproto "github.com/cometbft/cometbft/proto/tendermint/types"
)

// Bond denom should be set for staking test
const bondDenom = initiaapp.BondDenom
const secondBondDenom = "ulp"

var (
	priv1 = secp256k1.GenPrivKey()
	addr1 = sdk.AccAddress(priv1.PubKey().Address())
	priv2 = secp256k1.GenPrivKey()
	addr2 = sdk.AccAddress(priv2.PubKey().Address())

	valKey = ed25519.GenPrivKey()

	commissionRates = stakingtypes.NewCommissionRates(math.LegacyZeroDec(), math.LegacyZeroDec(), math.LegacyZeroDec())

	genCoins       = sdk.NewCoins(sdk.NewCoin(bondDenom, sdk.NewInt(5000000))).Sort()
	bondCoin       = sdk.NewCoin(bondDenom, sdk.NewInt(1_000_000))
	secondBondCoin = sdk.NewCoin(secondBondDenom, sdk.NewInt(1_000_000))
)

func checkBalance(t *testing.T, app *initiaapp.InitiaApp, addr sdk.AccAddress, balances sdk.Coins) {
	ctxCheck := app.BaseApp.NewContext(true, tmproto.Header{})
	require.True(t, balances.IsEqual(app.BankKeeper.GetAllBalances(ctxCheck, addr)))
}

func createApp(t *testing.T) *initiaapp.InitiaApp {
	baseCoin := sdk.NewCoin(bondDenom, sdk.NewInt(1_000_000_000_000))
	quoteCoin := sdk.NewCoin("uusdc", sdk.NewInt(2_500_000_000_000))
	dexCoins := sdk.NewCoins(baseCoin, quoteCoin)

	app := initiaapp.SetupWithGenesisAccounts(nil, authtypes.GenesisAccounts{
		&authtypes.BaseAccount{Address: addr1.String()},
		&authtypes.BaseAccount{Address: addr2.String()},
		&authtypes.BaseAccount{Address: types.StdAddr.String()},
	},
		banktypes.Balance{Address: addr1.String(), Coins: genCoins},
		banktypes.Balance{Address: addr2.String(), Coins: genCoins},
		banktypes.Balance{Address: types.StdAddr.String(), Coins: dexCoins},
	)

	checkBalance(t, app, addr1, genCoins)
	checkBalance(t, app, addr2, genCoins)
	checkBalance(t, app, types.StdAddr, dexCoins)

	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	ctx := app.BaseApp.NewContext(false, tmproto.Header{})
	createDexPool(t, ctx, app, baseCoin, quoteCoin, sdk.NewDecWithPrec(8, 1), sdk.NewDecWithPrec(2, 1))

	// set reward weight
	distrParams := customdistrtypes.DefaultParams()
	distrParams.RewardWeights = []customdistrtypes.RewardWeight{
		{Denom: bondDenom, Weight: sdk.OneDec()},
	}
	app.DistrKeeper.SetParams(ctx, distrParams)
	app.StakingKeeper.SetBondDenoms(ctx, []string{bondDenom, secondBondDenom})

	// fund second bond coin
	app.BankKeeper.SendCoins(ctx, types.StdAddr, addr1, sdk.NewCoins(secondBondCoin))
	app.EndBlock(abci.RequestEndBlock{})
	app.Commit()

	// create validator
	description := stakingtypes.NewDescription("foo_moniker", "", "", "", "")
	createValidatorMsg, err := stakingtypes.NewMsgCreateValidator(
		sdk.ValAddress(addr1), valKey.PubKey(), sdk.NewCoins(bondCoin, secondBondCoin), description, commissionRates,
	)
	require.NoError(t, err)

	err = executeMsgs(t, app, []sdk.Msg{createValidatorMsg}, []uint64{0}, []uint64{0}, priv1)
	require.NoError(t, err)

	checkBalance(t, app, addr1, genCoins.Sub(bondCoin))

	header = tmproto.Header{Height: app.LastBlockHeight() + 1}
	app.BeginBlock(abci.RequestBeginBlock{Header: header})

	return app
}

func executeMsgs(t *testing.T, app *initiaapp.InitiaApp, msgs []sdk.Msg, accountNum []uint64, sequenceNum []uint64, priv ...cryptotypes.PrivKey) error {
	txGen := initiaapp.MakeEncodingConfig().TxConfig
	header := tmproto.Header{Height: app.LastBlockHeight() + 1}
	_, _, err := testutilsims.SignCheckDeliver(t, txGen, app.BaseApp, header, msgs, "", accountNum, sequenceNum, true, true, priv...)
	return err
}

func decToVmArgument(t *testing.T, val math.LegacyDec) []byte {
	bz := val.BigInt().Bytes()
	diff := 16 - len(bz)
	require.True(t, diff >= 0)
	if diff > 0 {
		bz = append(bytes.Repeat([]byte{0}, diff), bz...)
	}

	high := binary.BigEndian.Uint64(bz[:8])
	low := binary.BigEndian.Uint64(bz[8:16])

	// serialize to uint128
	bz, err := vmtypes.SerializeUint128(high, low)
	require.NoError(t, err)

	return bz
}

func createDexPool(
	t *testing.T, ctx sdk.Context, app *initiaapp.InitiaApp,
	baseCoin sdk.Coin, quoteCoin sdk.Coin,
	weightBase sdk.Dec, weightQuote sdk.Dec,
) (metadataLP vmtypes.AccountAddress) {
	metadataBase, err := types.MetadataAddressFromDenom(baseCoin.Denom)
	require.NoError(t, err)

	metadataQuote, err := types.MetadataAddressFromDenom(quoteCoin.Denom)
	require.NoError(t, err)

	denomLP := "ulp"

	//
	// prepare arguments
	//

	name, err := vmtypes.SerializeString("LP Coin")
	require.NoError(t, err)

	symbol, err := vmtypes.SerializeString(denomLP)
	require.NoError(t, err)

	// 0.003 == 0.3%
	swapFeeBz := decToVmArgument(t, math.LegacyNewDecWithPrec(3, 3))
	weightBaseBz := decToVmArgument(t, weightBase)
	weightQuoteBz := decToVmArgument(t, weightQuote)

	baseAmount, err := vmtypes.SerializeUint64(baseCoin.Amount.Uint64())
	require.NoError(t, err)

	quoteAmount, err := vmtypes.SerializeUint64(quoteCoin.Amount.Uint64())
	require.NoError(t, err)

	err = app.MoveKeeper.ExecuteEntryFunction(
		ctx,
		vmtypes.StdAddress,
		vmtypes.StdAddress,
		"dex",
		"create_pair_script",
		[]vmtypes.TypeTag{},
		[][]byte{
			name,
			symbol,
			swapFeeBz,
			weightBaseBz,
			weightQuoteBz,
			metadataBase[:],
			metadataQuote[:],
			baseAmount,
			quoteAmount,
		},
	)
	require.NoError(t, err)

	return types.NamedObjectAddress(vmtypes.StdAddress, denomLP)
}

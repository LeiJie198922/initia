package keeper_test

import (
	"testing"

	"github.com/cometbft/cometbft/crypto/secp256k1"
	"github.com/stretchr/testify/require"

	dbm "github.com/cosmos/cosmos-db"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/initia-labs/initia/x/ibc/perm/types"
)

func Test_ExportGenesis(t *testing.T) {
	ctx, k := _createTestInput(t, dbm.NewMemDB())

	channelA := "channel-123"
	channelB := "channel-456"
	portA := "port-123"
	portB := "port-456"

	pubKeyA := secp256k1.GenPrivKey().PubKey()
	pubKeyB := secp256k1.GenPrivKey().PubKey()

	addrA := sdk.AccAddress(pubKeyA.Address())
	addrB := sdk.AccAddress(pubKeyB.Address())

	genState := types.NewGenesisState(
		[]types.PermissionedRelayer{
			{
				PortId:    portA,
				ChannelId: channelA,
				Relayer:   addrA.String(),
			},
			{
				PortId:    portB,
				ChannelId: channelB,
				Relayer:   addrB.String(),
			},
		},
	)
	k.InitGenesis(ctx, *genState)

	exportedState := k.ExportGenesis(ctx)
	require.Equal(t, genState, exportedState)
}

package keeper

import (
	"fmt"
	"time"

	gogotypes "github.com/cosmos/gogoproto/types"

	"github.com/initia-labs/initia/x/mstaking/types"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

// get a single validator
func (k Keeper) GetValidator(ctx sdk.Context, addr sdk.ValAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	value := store.Get(types.GetValidatorKey(addr))
	if value == nil {
		return validator, false
	}

	return types.MustUnmarshalValidator(k.cdc, value), true
}

func (k Keeper) mustGetValidator(ctx sdk.Context, addr sdk.ValAddress) types.Validator {
	validator, found := k.GetValidator(ctx, addr)
	if !found {
		panic(fmt.Sprintf("validator record not found for address: %X\n", addr))
	}

	return validator
}

// get a single validator by consensus address
func (k Keeper) GetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) (validator types.Validator, found bool) {
	store := ctx.KVStore(k.storeKey)

	opAddr := store.Get(types.GetValidatorByConsAddrKey(consAddr))
	if opAddr == nil {
		return validator, false
	}

	return k.GetValidator(ctx, opAddr)
}

func (k Keeper) mustGetValidatorByConsAddr(ctx sdk.Context, consAddr sdk.ConsAddress) types.Validator {
	validator, found := k.GetValidatorByConsAddr(ctx, consAddr)
	if !found {
		panic(fmt.Errorf("validator with consensus-Address %s not found", consAddr))
	}

	return validator
}

// set the main record holding validator details
func (k Keeper) SetValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	bz := types.MustMarshalValidator(k.cdc, &validator)
	store.Set(types.GetValidatorKey(validator.GetOperator()), bz)
}

// validator index
func (k Keeper) SetValidatorByConsAddr(ctx sdk.Context, validator types.Validator) error {
	consPk, err := validator.GetConsAddr()
	if err != nil {
		return err
	}
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorByConsAddrKey(consPk), validator.GetOperator())
	return nil
}

// validator index
func (k Keeper) SetValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorsByPowerIndexKey(validator, k.PowerReduction(ctx)), validator.GetOperator())
}

// validator index
func (k Keeper) DeleteValidatorByPowerIndex(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorsByPowerIndexKey(validator, k.PowerReduction(ctx)))
}

// Update the tokens of an existing validator.
// NOTE: validators power index key not updated here
func (k Keeper) AddValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	tokensToAdd sdk.Coins) (valOut types.Validator, addedShares sdk.DecCoins) {
	validator, addedShares = validator.AddTokensFromDel(tokensToAdd)

	// add a validator to whitelist group
	if !k.IsWhitelist(ctx, validator) {
		votingPower := k.VotingPower(ctx, validator.Tokens)
		if votingPower.GTE(k.MinVotingPower(ctx)) {
			k.AddWhitelistValidator(ctx, validator)
		}
	}

	k.SetValidator(ctx, validator)
	return validator, addedShares
}

// Update the tokens of an existing validator.
// NOTE: validators power index key not updated here
func (k Keeper) RemoveValidatorTokensAndShares(ctx sdk.Context, validator types.Validator,
	sharesToRemove sdk.DecCoins) (valOut types.Validator, removedTokens sdk.Coins) {
	validator, removedTokens = validator.RemoveDelShares(sharesToRemove)

	k.SetValidator(ctx, validator)
	return validator, removedTokens
}

// Update the tokens of an existing validator.
// NOTE: validators power index key not updated here
func (k Keeper) RemoveValidatorTokens(ctx sdk.Context,
	validator types.Validator, tokensToRemove sdk.Coins) types.Validator {
	validator = validator.RemoveTokens(tokensToRemove)

	k.SetValidator(ctx, validator)
	return validator
}

// UpdateValidatorCommission attempts to update a validator's commission rate.
// An error is returned if the new commission rate is invalid.
func (k Keeper) UpdateValidatorCommission(ctx sdk.Context,
	validator types.Validator, newRate sdk.Dec) (types.Commission, error) {
	commission := validator.Commission
	blockTime := ctx.BlockHeader().Time

	if err := commission.ValidateNewRate(newRate, blockTime); err != nil {
		return commission, err
	}

	if minCommission := k.MinCommissionRate(ctx); newRate.LT(minCommission) {
		return commission, fmt.Errorf("cannot set validator commission to less than minimum rate of %s", minCommission)
	}

	commission.Rate = newRate
	commission.UpdateTime = blockTime

	return commission, nil
}

// remove the validator record and associated indexes
// except for the bonded validator index which is only handled in ApplyAndReturnTendermintUpdates
// TODO, this function panics, and it's not good.
func (k Keeper) RemoveValidator(ctx sdk.Context, address sdk.ValAddress) {
	// first retrieve the old validator record
	validator, found := k.GetValidator(ctx, address)
	if !found {
		return
	}

	if !validator.IsUnbonded() {
		panic("cannot call RemoveValidator on bonded or unbonding validators")
	}

	if !validator.Tokens.IsZero() {
		panic("attempting to remove a validator which still contains tokens")
	}

	valConsAddr, err := validator.GetConsAddr()
	if err != nil {
		panic(err)
	}

	// delete the old validator record
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorKey(address))
	store.Delete(types.GetValidatorByConsAddrKey(valConsAddr))
	store.Delete(types.GetValidatorsByPowerIndexKey(validator, k.PowerReduction(ctx)))

	// call hooks
	if err := k.Hooks().AfterValidatorRemoved(ctx, valConsAddr, validator.GetOperator()); err != nil {
		k.Logger(ctx).Error("error in after validator removed hook", "error", err)
	}
}

// get groups of validators

// get the set of all validators with no limits, used during genesis dump
func (k Keeper) GetAllValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		validators = append(validators, types.MustUnmarshalValidator(k.cdc, iterator.Value()))
	}

	return validators
}

// return a given amount of all the validators
func (k Keeper) GetValidators(ctx sdk.Context, maxRetrieve uint32) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)
	validators = make([]types.Validator, maxRetrieve)

	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorsKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxRetrieve); iterator.Next() {
		validators[i] = types.MustUnmarshalValidator(k.cdc, iterator.Value())
		i++
	}

	return validators[:i] // trim if the array length < maxRetrieve
}

// get the current group of bonded validators sorted by power-rank
func (k Keeper) GetBondedValidatorsByPower(ctx sdk.Context) []types.Validator {
	maxValidators := k.MaxValidators(ctx)
	validators := make([]types.Validator, maxValidators)

	iterator := k.ValidatorsPowerStoreIterator(ctx)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid() && i < int(maxValidators); iterator.Next() {
		address := iterator.Value()
		validator := k.mustGetValidator(ctx, address)

		if validator.IsBonded() {
			validators[i] = validator
			i++
		}
	}

	return validators[:i] // trim
}

// returns an iterator for the current validator power store
func (k Keeper) ValidatorsPowerStoreIterator(ctx sdk.Context) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return sdk.KVStoreReversePrefixIterator(store, types.ValidatorsByPowerIndexKey)
}

// Last Validator Index

// Load the last validator power.
// Returns zero if the operator was not a validator last block.
func (k Keeper) GetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) (power int64) {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetLastValidatorPowerKey(operator))
	if bz == nil {
		return 0
	}

	intV := gogotypes.Int64Value{}
	k.cdc.MustUnmarshal(bz, &intV)

	return intV.GetValue()
}

// Set the last validator power.
func (k Keeper) SetLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress, power int64) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&gogotypes.Int64Value{Value: power})
	store.Set(types.GetLastValidatorPowerKey(operator), bz)
}

// Delete the last validator power.
func (k Keeper) DeleteLastValidatorPower(ctx sdk.Context, operator sdk.ValAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetLastValidatorPowerKey(operator))
}

// returns an iterator for the consensus validators in the last block
func (k Keeper) LastValidatorsIterator(ctx sdk.Context) (iterator sdk.Iterator) {
	store := ctx.KVStore(k.storeKey)
	iterator = sdk.KVStorePrefixIterator(store, types.LastValidatorPowerKey)

	return iterator
}

// Iterate over last validator powers.
func (k Keeper) IterateLastValidatorPowers(ctx sdk.Context, handler func(operator sdk.ValAddress, power int64) (stop bool)) {
	store := ctx.KVStore(k.storeKey)

	iter := sdk.KVStorePrefixIterator(store, types.LastValidatorPowerKey)
	defer iter.Close()

	for ; iter.Valid(); iter.Next() {
		addr := sdk.ValAddress(types.AddressFromLastValidatorPowerKey(iter.Key()))
		intV := &gogotypes.Int64Value{}

		k.cdc.MustUnmarshal(iter.Value(), intV)

		if handler(addr, intV.GetValue()) {
			break
		}
	}
}

// get the group of the bonded validators
func (k Keeper) GetLastValidators(ctx sdk.Context) (validators []types.Validator) {
	store := ctx.KVStore(k.storeKey)

	// add the actual validator power sorted store
	maxValidators := k.MaxValidators(ctx)
	validators = make([]types.Validator, maxValidators)

	iterator := sdk.KVStorePrefixIterator(store, types.LastValidatorPowerKey)
	defer iterator.Close()

	i := 0
	for ; iterator.Valid(); iterator.Next() {
		// sanity check
		if i >= int(maxValidators) {
			panic("more validators than maxValidators found")
		}

		address := types.AddressFromLastValidatorPowerKey(iterator.Key())
		validator := k.mustGetValidator(ctx, address)

		validators[i] = validator
		i++
	}

	return validators[:i] // trim
}

// GetUnbondingValidators returns a slice of mature validator addresses that
// complete their unbonding at a given time and height.
func (k Keeper) GetUnbondingValidators(ctx sdk.Context, endTime time.Time, endHeight int64) []string {
	store := ctx.KVStore(k.storeKey)

	bz := store.Get(types.GetValidatorQueueKey(endTime, endHeight))
	if bz == nil {
		return []string{}
	}

	addrs := types.ValAddresses{}
	k.cdc.MustUnmarshal(bz, &addrs)

	return addrs.Addresses
}

// SetUnbondingValidatorsQueue sets a given slice of validator addresses into
// the unbonding validator queue by a given height and time.
func (k Keeper) SetUnbondingValidatorsQueue(ctx sdk.Context, endTime time.Time, endHeight int64, addrs []string) {
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&types.ValAddresses{Addresses: addrs})
	store.Set(types.GetValidatorQueueKey(endTime, endHeight), bz)
}

// InsertUnbondingValidatorQueue inserts a given unbonding validator address into
// the unbonding validator queue for a given height and time.
func (k Keeper) InsertUnbondingValidatorQueue(ctx sdk.Context, val types.Validator) {
	addrs := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	addrs = append(addrs, val.OperatorAddress)
	k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, addrs)
}

// DeleteValidatorQueueTimeSlice deletes all entries in the queue indexed by a
// given height and time.
func (k Keeper) DeleteValidatorQueueTimeSlice(ctx sdk.Context, endTime time.Time, endHeight int64) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorQueueKey(endTime, endHeight))
}

// DeleteValidatorQueue removes a validator by address from the unbonding queue
// indexed by a given height and time.
func (k Keeper) DeleteValidatorQueue(ctx sdk.Context, val types.Validator) {
	addrs := k.GetUnbondingValidators(ctx, val.UnbondingTime, val.UnbondingHeight)
	newAddrs := []string{}

	for _, addr := range addrs {
		if addr != val.OperatorAddress {
			newAddrs = append(newAddrs, addr)
		}
	}

	if len(newAddrs) == 0 {
		k.DeleteValidatorQueueTimeSlice(ctx, val.UnbondingTime, val.UnbondingHeight)
	} else {
		k.SetUnbondingValidatorsQueue(ctx, val.UnbondingTime, val.UnbondingHeight, newAddrs)
	}
}

// ValidatorQueueIterator returns an interator ranging over validators that are
// unbonding whose unbonding completion occurs at the given height and time.
func (k Keeper) ValidatorQueueIterator(ctx sdk.Context, endTime time.Time, endHeight int64) sdk.Iterator {
	store := ctx.KVStore(k.storeKey)
	return store.Iterator(types.ValidatorQueueKey, sdk.InclusiveEndBytes(types.GetValidatorQueueKey(endTime, endHeight)))
}

// UnbondAllMatureValidators unbonds all the mature unbonding validators that
// have finished their unbonding period.
func (k Keeper) UnbondAllMatureValidators(ctx sdk.Context) {
	blockTime := ctx.BlockTime()
	blockHeight := ctx.BlockHeight()

	// unbondingValIterator will contains all validator addresses indexed under
	// the ValidatorQueueKey prefix. Note, the entire index key is composed as
	// ValidatorQueueKey | timeBzLen (8-byte big endian) | timeBz | heightBz (8-byte big endian),
	// so it may be possible that certain validator addresses that are iterated
	// over are not ready to unbond, so an explicit check is required.
	unbondingValIterator := k.ValidatorQueueIterator(ctx, blockTime, blockHeight)
	defer unbondingValIterator.Close()

	for ; unbondingValIterator.Valid(); unbondingValIterator.Next() {
		key := unbondingValIterator.Key()
		keyTime, keyHeight, err := types.ParseValidatorQueueKey(key)
		if err != nil {
			panic(fmt.Errorf("failed to parse unbonding key: %w", err))
		}

		// All addresses for the given key have the same unbonding height and time.
		// We only unbond if the height and time are less than the current height
		// and time.
		if keyHeight <= blockHeight && (keyTime.Before(blockTime) || keyTime.Equal(blockTime)) {
			addrs := types.ValAddresses{}
			k.cdc.MustUnmarshal(unbondingValIterator.Value(), &addrs)

			for _, valAddr := range addrs.Addresses {
				addr, err := sdk.ValAddressFromBech32(valAddr)
				if err != nil {
					panic(err)
				}
				val, found := k.GetValidator(ctx, addr)
				if !found {
					panic("validator in the unbonding queue was not found")
				}

				if !val.IsUnbonding() {
					panic("unexpected validator in unbonding queue; status was not unbonding")
				}

				if val.UnbondingOnHoldRefCount == 0 {
					k.DeleteUnbondingIndex(ctx, val.UnbondingId)

					// TODO - sdk is not calling k.SetValidator
					val.UnbondingId = 0
					val = k.UnbondingToUnbonded(ctx, val)

					if val.GetDelegatorShares().IsZero() {
						k.RemoveValidator(ctx, val.GetOperator())
					}

					// remove validator from queue
					k.DeleteValidatorQueue(ctx, val)
				}
			}
		}
	}
}

func (k Keeper) IsValidatorJailed(ctx sdk.Context, addr sdk.ConsAddress) bool {
	v, ok := k.GetValidatorByConsAddr(ctx, addr)
	if !ok {
		return false
	}

	return v.Jailed
}

// IsWhitelist return whether a validator is already in whitelist
func (k Keeper) IsWhitelist(ctx sdk.Context, validator types.Validator) bool {
	store := ctx.KVStore(k.storeKey)
	return store.Has(types.GetValidatorWhitelistKey(validator.GetOperator()))
}

// AddWhitelistValidator add validator to power update whitelist
func (k Keeper) AddWhitelistValidator(ctx sdk.Context, validator types.Validator) {
	// jailed validators are not kept in the power update whitelist
	if validator.Jailed {
		return
	}

	store := ctx.KVStore(k.storeKey)
	store.Set(types.GetValidatorWhitelistKey(validator.GetOperator()), []byte{1})
}

// RemoveWhitelistValidator remove validator from power update whitelist
func (k Keeper) RemoveWhitelistValidator(ctx sdk.Context, validator types.Validator) {
	store := ctx.KVStore(k.storeKey)
	store.Delete(types.GetValidatorWhitelistKey(validator.GetOperator()))

	// remove power index
	k.DeleteValidatorByPowerIndex(ctx, validator)
}

// IterateWhitelistValidator iterate all whitelist validators
func (k Keeper) IterateWhitelistValidator(ctx sdk.Context, handler func(validator types.Validator) bool) {
	store := ctx.KVStore(k.storeKey)
	iterator := sdk.KVStorePrefixIterator(store, types.ValidatorWhitelistKey)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		valAddr := sdk.ValAddress(types.AddressFromValidatorWhitelistKey(iterator.Key()))
		validator := k.mustGetValidator(ctx, valAddr)
		if handler(validator) {
			break
		}
	}
}

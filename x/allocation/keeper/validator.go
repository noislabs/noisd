package keeper

import (
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

func (k Keeper) SetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress, amount int64) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	bz := sdk.Uint64ToBigEndian(uint64(amount))
	prefixStore.Set(operator, bz)
}

func (k Keeper) DeleteValidatorRewards(ctx sdk.Context, operator sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	prefixStore.Delete(operator)
}

func (k Keeper) GetValidatorRewards(ctx sdk.Context, operator sdk.AccAddress) int64 {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	bz := prefixStore.Get(operator)
	if bz == nil {
		return 0
	}
	return int64(sdk.BigEndianToUint64(bz))
}

// ValidatorRewards returns the rewards for all validators used for genesis export
func (k Keeper) ValidatorRewards(ctx sdk.Context) []types.ValidatorReward {
	store := ctx.KVStore(k.storeKey)
	prefixStore := prefix.NewStore(store, types.ValidatorRewardsPrefix)
	iterator := prefixStore.Iterator(nil, nil)
	defer iterator.Close()
	validatorRewards := []types.ValidatorReward{}
	denom := k.stakingKeeper.BondDenom(ctx)
	for ; iterator.Valid(); iterator.Next() {
		address := sdk.AccAddress(iterator.Key())
		amount := int64(sdk.BigEndianToUint64(iterator.Value()))
		reward := types.ValidatorReward{
			Address: address.String(),
			Rewards: sdk.NewCoins(sdk.NewInt64Coin(denom, amount)),
		}
		validatorRewards = append(validatorRewards, reward)
	}
	return validatorRewards
}

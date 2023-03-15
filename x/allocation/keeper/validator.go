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

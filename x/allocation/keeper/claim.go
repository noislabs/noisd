package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) ClaimRewards(ctx sdk.Context, operator sdk.AccAddress) (sdk.Coins, error) {
	return sdk.Coins{}, nil
}

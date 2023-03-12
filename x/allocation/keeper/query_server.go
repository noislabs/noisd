package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the allocation module.
func (k Keeper) Params(c context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	params := k.GetParams(ctx)

	return &types.QueryParamsResponse{Params: params}, nil
}

func (k Keeper) ClaimableRewards(c context.Context, _ *types.QueryClaimableRewardsRequest) (*types.QueryClaimableRewardsResponse, error) {
	return &types.QueryClaimableRewardsResponse{}, nil
}

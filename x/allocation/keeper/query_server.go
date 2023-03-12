package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

var _ types.QueryServer = Keeper{}

// Params returns params of the allocation module.
func (k Keeper) Params(goCtx context.Context, _ *types.QueryParamsRequest) (*types.QueryParamsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	params := k.GetParams(ctx)
	return &types.QueryParamsResponse{Params: params}, nil
}

// ClaimableRewards returns the claimable amount for a validator.
func (k Keeper) ClaimableRewards(goCtx context.Context, req *types.QueryClaimableRewardsRequest) (*types.QueryClaimableRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	address, err := sdk.AccAddressFromBech32(req.Address)
	if err != nil {
		return nil, err
	}
	validatorRewards := k.GetValidatorRewards(ctx, address)
	return &types.QueryClaimableRewardsResponse{Coins: validatorRewards.Rewards}, nil
}

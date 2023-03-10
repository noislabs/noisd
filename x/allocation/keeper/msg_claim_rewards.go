package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/types"
)

func (server msgServer) ClaimRewards(goCtx context.Context, msg *types.MsgClaimRewards) (*types.MsgClaimRewardsResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}
	coins, err := server.Keeper.ClaimRewards(ctx, sender)
	if err != nil {
		return nil, err
	}
	return &types.MsgClaimRewardsResponse{
		ClaimedRewards: coins,
	}, nil
}

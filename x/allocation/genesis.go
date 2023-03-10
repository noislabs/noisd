package allocation

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/noislabs/noisd/x/allocation/keeper"
	"github.com/noislabs/noisd/x/allocation/types"
)

// InitGenesis initializes the alloc module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the alloc module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	return &types.GenesisState{
		Params: k.GetParams(ctx),
	}
}

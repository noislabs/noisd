package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type StoreCodeDecorator struct {
}

func (scd StoreCodeDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	return next(ctx, tx, simulate)
}

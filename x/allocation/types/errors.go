package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/allocation module sentinel errors
// errors must start at 2 to avoid conflict with the default errors by the sdk
var (
	ErrNoRewards = sdkerrors.Register(ModuleName, 2, "no rewards available for claiming")
)

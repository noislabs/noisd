package types

import (
	errorsmod "cosmossdk.io/errors"
)

// x/allocation module sentinel errors
// errors must start at 2 to avoid conflict with the default errors by the sdk
var (
	ErrNoRewards = errorsmod.Register(ModuleName, 2, "no rewards available for claiming")
)

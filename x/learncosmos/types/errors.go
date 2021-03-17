package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/learncosmos module sentinel errors
var (
	ErrSample         = sdkerrors.Register(ModuleName, 1100, "sample error")
	ErrInvalidVersion = sdkerrors.Register(ModuleName, 8, "invalid ICS20 version")
)

package keeper

import (
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

var _ types.QueryServer = Keeper{}

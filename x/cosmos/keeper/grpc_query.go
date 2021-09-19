package keeper

import (
	"github.com/aaely/cosmos/x/cosmos/types"
)

var _ types.QueryServer = Keeper{}

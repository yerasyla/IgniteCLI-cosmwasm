package keeper

import (
	"planet/x/planet/types"
)

var _ types.QueryServer = Keeper{}

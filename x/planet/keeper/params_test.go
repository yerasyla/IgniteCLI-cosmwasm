package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	testkeeper "planet/testutil/keeper"
	"planet/x/planet/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.PlanetKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

package keeper_test

import (
	"testing"

	testkeeper "github.com/mastervectormaster/bluechip/testutil/keeper"
	"github.com/mastervectormaster/bluechip/x/bluechip/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.BluechipKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}

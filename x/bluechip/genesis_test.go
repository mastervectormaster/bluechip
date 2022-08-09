package bluechip_test

import (
	"testing"

	keepertest "github.com/mastervectormaster/bluechip/testutil/keeper"
	"github.com/mastervectormaster/bluechip/testutil/nullify"
	"github.com/mastervectormaster/bluechip/x/bluechip"
	"github.com/mastervectormaster/bluechip/x/bluechip/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.BluechipKeeper(t)
	bluechip.InitGenesis(ctx, *k, genesisState)
	got := bluechip.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}

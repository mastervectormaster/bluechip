package keeper_test

import (
	"context"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	keepertest "github.com/mastervectormaster/bluechip/testutil/keeper"
	"github.com/mastervectormaster/bluechip/x/bluechip/keeper"
	"github.com/mastervectormaster/bluechip/x/bluechip/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.BluechipKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/Stride-Labs/stride/testutil/keeper"
	"github.com/Stride-Labs/stride/x/records/keeper"
	"github.com/Stride-Labs/stride/x/records/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.RecordsKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}

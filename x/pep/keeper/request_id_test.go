package keeper_test

import (
	"context"
	"strconv"
	"testing"

	keepertest "github.com/Fairblock/fairyring/testutil/keeper"
	"github.com/Fairblock/fairyring/testutil/nullify"
	"github.com/Fairblock/fairyring/x/pep/keeper"
	"github.com/Fairblock/fairyring/x/pep/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNRequestId(keeper keeper.Keeper, ctx context.Context, n int) []types.RequestId {
	items := make([]types.RequestId, n)
	for i := range items {
		items[i].Creator = strconv.Itoa(i)

		keeper.SetRequestId(ctx, items[i])
	}
	return items
}

func TestRequestIdGet(t *testing.T) {
	keeper, ctx := keepertest.PepKeeper(t)
	items := createNRequestId(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetRequestId(ctx,
			item.Creator,
			"",
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}

// func TestRequestIdRemove(t *testing.T) {
// 	keeper, ctx := keepertest.PepKeeper(t)
// 	items := createNRequestId(keeper, ctx, 10)
// 	for _, item := range items {
// 		keeper.RemoveRequestId(ctx,
// 			item.Creator,
// 		)
// 		_, found := keeper.GetRequestId(ctx,
// 			item.Creator,
// 		)
// 		require.False(t, found)
// 	}
// }

func TestRequestIdGetAll(t *testing.T) {
	keeper, ctx := keepertest.PepKeeper(t)
	items := createNRequestId(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllRequestId(ctx)),
	)
}
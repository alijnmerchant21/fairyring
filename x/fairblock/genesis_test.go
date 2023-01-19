package fairblock_test

import (
	"testing"

	keepertest "fairyring/testutil/keeper"
	"fairyring/testutil/nullify"
	"fairyring/x/fairblock"
	"fairyring/x/fairblock/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		EncryptedTxList: []types.EncryptedTx{
			{
				TargetHeight: 0,
				Index:        0,
			},
			{
				TargetHeight: 1,
				Index:        1,
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.FairblockKeeper(t)
	fairblock.InitGenesis(ctx, *k, genesisState)
	got := fairblock.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	require.ElementsMatch(t, genesisState.EncryptedTxList, got.EncryptedTxList)
	// this line is used by starport scaffolding # genesis/test/assert
}
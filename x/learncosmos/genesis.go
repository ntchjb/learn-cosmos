package learncosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/keeper"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// this line is used by starport scaffolding # genesis/module/init
	k.SetGoldPool(ctx, genState)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()

	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}

package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

type (
	Keeper struct {
		cdc      codec.Marshaler
		storeKey sdk.StoreKey
		memKey   sdk.StoreKey
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey) *Keeper {
	return &Keeper{
		cdc:      cdc,
		storeKey: storeKey,
		memKey:   memKey,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetGoldPool set origin amount of gold in pool
func (k Keeper) SetGoldPool(ctx sdk.Context, genesisState types.GenesisState) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoldPoolKey))

	key := types.KeyPrefix(types.GoldPoolKey)
	value := k.cdc.MustMarshalBinaryBare(&genesisState)

	store.Set(key, value)
}

// GetGoldPool get origin amount of gold in pool
func (k Keeper) GetGoldPool(ctx sdk.Context) types.GenesisState {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoldPoolKey))

	key := types.KeyPrefix(types.GoldPoolKey)
	var genesisState types.GenesisState
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &genesisState)

	return genesisState
}

package keeper

import (
	"context"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) GoldPool(c context.Context, req *types.QueryGoldPoolRequest) (*types.QueryGoldPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var goldPool types.GenesisState
	ctx := sdk.UnwrapSDKContext(c)

	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoldPoolKey))
	k.cdc.MustUnmarshalBinaryBare(store.Get(types.KeyPrefix(types.GoldPoolKey)), &goldPool)

	return &types.QueryGoldPoolResponse{GoldAmount: goldPool.GoldAmount, GoldPricePerUnit: goldPool.GoldUnitPriceUusd}, nil
}

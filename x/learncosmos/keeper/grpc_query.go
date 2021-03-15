package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) GoldPool(c context.Context, req *types.QueryGoldPoolRequest) (*types.QueryGoldPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	goldPool := k.GetGoldPool(ctx)

	return &types.QueryGoldPoolResponse{
		GoldAmount: goldPool.Amount,
	}, nil
}

func (k Keeper) OwnedGold(c context.Context, req *types.QueryOwnedGoldRequest) (*types.QueryOwnedGoldResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	ownedGold := k.GetOwnedGold(ctx, req.Owner)

	return &types.QueryOwnedGoldResponse{Amount: ownedGold.Amount}, nil
}

func (k Keeper) Order(c context.Context, req *types.QueryPoolOrderRequest) (*types.QueryPoolOrderResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	order := k.GetOrder(ctx, req.Id)

	return &types.QueryPoolOrderResponse{
		Id:           order.Id,
		Type:         order.Type,
		UserAddr:     order.UserAddr,
		PricePerUnit: order.PricePerUnit,
		Amount:       order.Amount,
		Status:       order.Status,
		StatusReason: order.StatusReason,
	}, nil
}

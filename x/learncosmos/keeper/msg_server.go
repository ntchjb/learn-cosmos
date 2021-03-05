package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

type msgServer struct {
	Keeper
}

func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) BuyGold(goCtx context.Context, msg *types.MsgBuyGold) (*types.MsgBuyGoldResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.BuyGoldFromPool(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgBuyGoldResponse{}, nil
}

func (k msgServer) SellGold(goCtx context.Context, msg *types.MsgSellGold) (*types.MsgSellGoldResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.SellGoldToPool(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgSellGoldResponse{}, nil
}

func (k msgServer) TransferGold(goCtx context.Context, msg *types.MsgTransferGold) (*types.MsgTransferGoldResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	if err := k.Keeper.TransferGold(ctx, *msg); err != nil {
		return nil, err
	}

	return &types.MsgTransferGoldResponse{}, nil
}

package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
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

	buyerAddr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid buyer address: %w", err)
	}
	orderID := k.CreatePendingOrder(ctx, types.OrderType_BUY, msg.Amount, buyerAddr)
	if err := k.RequestGoldPrice(ctx, msg.IbcChannel, orderID, msg.OracleScriptId); err != nil {
		return nil, err
	}

	return &types.MsgBuyGoldResponse{}, nil
}

func (k msgServer) SellGold(goCtx context.Context, msg *types.MsgSellGold) (*types.MsgSellGoldResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	sellerAddr, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "invalid seller address: %w", err)
	}
	orderID := k.CreatePendingOrder(ctx, types.OrderType_SELL, msg.Amount, sellerAddr)
	if err := k.RequestGoldPrice(ctx, msg.IbcChannel, orderID, msg.OracleScriptId); err != nil {
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

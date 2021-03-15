package keeper

import (
	"context"
	"strings"

	"github.com/bandprotocol/chain/pkg/obi"
	oracle "github.com/bandprotocol/chain/x/oracle/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
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

func (k msgServer) ProcessIBCPacket(goCtx context.Context, msg *channeltypes.MsgRecvPacket) (*channeltypes.MsgRecvPacketResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// Unmarshal data in MsgRecvPacket
	var oracleResult oracle.OracleResponsePacketData
	if err := k.cdc.UnmarshalBinaryBare(msg.Packet.Data, &oracleResult); err != nil {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unable to unmarshal oracle result: %w", err)
	}

	// Get order ID from client ID of the MsgRecvPacket
	clientIDList := strings.Split(oracleResult.ClientID, ":")
	if len(clientIDList) != 2 {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "unable to get order ID from oracle's client ID")
	}
	orderID := strings.Split(oracleResult.ClientID, ":")[1]

	var obiOutput GoldPriceOBIOutput
	obi.MustDecode(oracleResult.Result, &obiOutput)
	if err := k.ProcessOrder(ctx, orderID, obiOutput.Price); err != nil {
		return nil, err
	}

	return &channeltypes.MsgRecvPacketResponse{}, nil
}

package learncosmos

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/keeper"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

func handleMsgBuyGold(ctx sdk.Context, k keeper.Keeper, msg *types.MsgBuyGold) (*sdk.Result, error) {
	if err := k.BuyGoldFromPool(ctx, *msg); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgSellGold(ctx sdk.Context, k keeper.Keeper, msg *types.MsgSellGold) (*sdk.Result, error) {
	if err := k.SellGoldToPool(ctx, *msg); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

func handleMsgTransferGold(ctx sdk.Context, k keeper.Keeper, msg *types.MsgTransferGold) (*sdk.Result, error) {
	if err := k.TransferGold(ctx, *msg); err != nil {
		return nil, err
	}

	return &sdk.Result{Events: ctx.EventManager().ABCIEvents()}, nil
}

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
		cdc        codec.Marshaler
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		bankKeeper types.BankKeeper
	}
)

func NewKeeper(cdc codec.Marshaler, storeKey, memKey sdk.StoreKey, bk types.BankKeeper) *Keeper {
	return &Keeper{
		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		bankKeeper: bk,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

// SetGoldPool set origin amount of gold in pool
func (k Keeper) SetGoldPool(ctx sdk.Context, goldPool types.GoldPool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoldPoolKey))

	key := types.KeyPrefix(types.GoldPoolKey)
	value := k.cdc.MustMarshalBinaryBare(&goldPool)

	store.Set(key, value)
}

// GetGoldPool get origin amount of gold in pool
func (k Keeper) GetGoldPool(ctx sdk.Context) types.GoldPool {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.GoldPoolKey))

	key := types.KeyPrefix(types.GoldPoolKey)
	var goldPool types.GoldPool
	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &goldPool)

	return goldPool
}

func (k Keeper) GetOwnedGold(ctx sdk.Context, owner string) types.OwnedGold {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnedGoldKey))

	key := types.KeyPrefix(types.OwnedGoldKey + owner)
	var ownedGold types.OwnedGold
	if !store.Has(key) {
		return types.OwnedGold{
			Owner: owner,
		}
	}

	k.cdc.MustUnmarshalBinaryBare(store.Get(key), &ownedGold)

	return ownedGold
}

func (k Keeper) SetOwnedGold(ctx sdk.Context, ownedGold types.OwnedGold) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.OwnedGoldKey))

	key := types.KeyPrefix(types.OwnedGoldKey + ownedGold.Owner)
	value := k.cdc.MustMarshalBinaryBare(&ownedGold)

	store.Set(key, value)
}

func (k Keeper) BuyGoldFromPool(ctx sdk.Context, msg types.MsgBuyGold) error {
	goldPool := k.GetGoldPool(ctx)
	ownedGold := k.GetOwnedGold(ctx, msg.Buyer)

	fmt.Println("goldPool is", goldPool, "and owned gold is", ownedGold)

	if msg.Amount > goldPool.Amount {
		return fmt.Errorf("gold amount in pool is not sufficient")
	}

	buyerAddr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return fmt.Errorf("buyer address is invalid")
	}

	currentBalance := k.bankKeeper.GetBalance(ctx, buyerAddr, "uusd")
	fmt.Println("current balance is", currentBalance)

	payAmount := msg.Amount * goldPool.PricePerUnit
	if currentBalance.Amount.Uint64() < payAmount {
		return fmt.Errorf("current buyer balance is not sufficient")
	}

	goldPool.Amount -= msg.Amount
	k.SetGoldPool(ctx, goldPool)

	ownedGold.Amount += msg.Amount
	k.SetOwnedGold(ctx, ownedGold)

	fmt.Println("new goldPool is", goldPool, "and new owned gold is", ownedGold)

	if err := k.bankKeeper.SendCoinsFromAccountToModule(
		ctx,
		buyerAddr,
		types.ModuleName,
		sdk.NewCoins(sdk.NewCoin("uusd", sdk.NewIntFromUint64(payAmount))),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) SellGoldToPool(ctx sdk.Context, msg types.MsgSellGold) error {
	goldPool := k.GetGoldPool(ctx)
	ownedGold := k.GetOwnedGold(ctx, msg.Seller)

	sellerAddr, _ := sdk.AccAddressFromBech32(msg.Seller)

	if msg.Amount > ownedGold.Amount {
		return fmt.Errorf("gold amount owned by sender is not sufficient")
	}

	goldPool.Amount += msg.Amount
	k.SetGoldPool(ctx, goldPool)
	ownedGold.Amount -= msg.Amount
	k.SetOwnedGold(ctx, ownedGold)

	if err := k.bankKeeper.SendCoinsFromModuleToAccount(
		ctx,
		types.ModuleName,
		sellerAddr,
		sdk.NewCoins(sdk.NewCoin("uusd", sdk.NewIntFromUint64(goldPool.PricePerUnit*msg.Amount))),
	); err != nil {
		return err
	}

	return nil
}

func (k Keeper) TransferGold(ctx sdk.Context, msg types.MsgTransferGold) error {
	senderOwnedGold := k.GetOwnedGold(ctx, msg.Sender)
	receiverOwnedGold := k.GetOwnedGold(ctx, msg.Receiver)

	if senderOwnedGold.Amount < msg.Amount {
		return fmt.Errorf("gold amount owned by sender is not sufficient")
	}

	senderOwnedGold.Amount -= msg.Amount
	k.SetOwnedGold(ctx, senderOwnedGold)
	receiverOwnedGold.Amount += msg.Amount
	k.SetOwnedGold(ctx, receiverOwnedGold)

	return nil
}

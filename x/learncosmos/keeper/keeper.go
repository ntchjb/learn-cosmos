package keeper

import (
	"fmt"
	"time"

	"github.com/bandprotocol/chain/pkg/obi"
	oracle "github.com/bandprotocol/chain/x/oracle/types"
	"github.com/tendermint/tendermint/libs/log"

	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	clienttypes "github.com/cosmos/cosmos-sdk/x/ibc/core/02-client/types"
	ibcChannelKeeper "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/keeper"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	host "github.com/cosmos/cosmos-sdk/x/ibc/core/24-host"
	"github.com/ntchjb/learn-cosmos/x/learncosmos/types"
)

type (
	Keeper struct {
		cdc           codec.Marshaler
		storeKey      sdk.StoreKey
		memKey        sdk.StoreKey
		bankKeeper    types.BankKeeper
		channelKeeper ibcChannelKeeper.Keeper
		scopedKeeper  capabilitykeeper.ScopedKeeper
	}

	GoldPriceOBI struct {
		Multiplier uint64
	}
)

func NewKeeper(cdc codec.Marshaler,
	storeKey, memKey sdk.StoreKey,
	bk types.BankKeeper,
	chk ibcChannelKeeper.Keeper,
	sck capabilitykeeper.ScopedKeeper,
) *Keeper {
	return &Keeper{
		cdc:           cdc,
		storeKey:      storeKey,
		memKey:        memKey,
		bankKeeper:    bk,
		channelKeeper: chk,
		scopedKeeper:  sck,
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

	if msg.Amount > goldPool.Amount {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "gold amount in pool is not sufficient")
	}

	buyerAddr, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "buyer address is invalid")
	}

	currentBalance := k.bankKeeper.GetBalance(ctx, buyerAddr, "uusd")
	fmt.Println("current balance is", currentBalance)

	payAmount := msg.Amount * goldPool.PricePerUnit
	if currentBalance.Amount.Uint64() < payAmount {
		return sdkerrors.Wrapf(sdkerrors.ErrInsufficientFunds, "current buyer balance is not sufficient")
	}

	goldPool.Amount -= msg.Amount
	k.SetGoldPool(ctx, goldPool)

	ownedGold.Amount += msg.Amount
	k.SetOwnedGold(ctx, ownedGold)

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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "gold amount owned by sender is not sufficient")
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
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidRequest, "gold amount owned by sender is not sufficient")
	}

	senderOwnedGold.Amount -= msg.Amount
	k.SetOwnedGold(ctx, senderOwnedGold)
	receiverOwnedGold.Amount += msg.Amount
	k.SetOwnedGold(ctx, receiverOwnedGold)

	return nil
}

func (k Keeper) RequestGoldPrice(ctx sdk.Context, ibcChannelID string, orderID string) error {
	sourcePort := types.ModuleName
	sourceChannel := ibcChannelID
	sourceChannelEnd, found := k.channelKeeper.GetChannel(ctx, sourcePort, sourceChannel)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown channel %s port %s",
			sourceChannel, sourcePort,
		)
	}
	destinationPort := sourceChannelEnd.Counterparty.PortId
	destinationChannel := sourceChannelEnd.Counterparty.ChannelId
	sequence, found := k.channelKeeper.GetNextSequenceSend(
		ctx, sourcePort, sourceChannel,
	)
	if !found {
		return sdkerrors.Wrapf(
			sdkerrors.ErrUnknownRequest,
			"unknown sequence number for channel %s port %s",
			sourceChannel, sourcePort,
		)
	}
	clientID := fmt.Sprintf("Order:%v", orderID)
	oracleScriptID := oracle.OracleScriptID(33)
	callData := obi.MustEncode(GoldPriceOBI{
		Multiplier: 100,
	})
	askCount := uint64(4)
	minCount := uint64(3)

	packet := oracle.NewOracleRequestPacketData(
		clientID, oracleScriptID, callData,
		askCount, minCount,
	)

	channelCap, ok := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(sourcePort, sourceChannel))
	if !ok {
		return sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
	}

	if err := k.channelKeeper.SendPacket(ctx, channelCap, channeltypes.NewPacket(
		packet.GetBytes(),
		sequence,
		sourcePort,
		sourceChannel,
		destinationPort,
		destinationChannel,
		clienttypes.NewHeight(0, 100),
		uint64(1*time.Minute),
	),
	); err != nil {
		return err
	}

	return nil
}

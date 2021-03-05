package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	// this line is used by starport scaffolding # 1
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgBuyGold{}, "learncosmos/BuyGold", nil)
	cdc.RegisterConcrete(&MsgSellGold{}, "learncosmos/SellGold", nil)
	cdc.RegisterConcrete(&MsgTransferGold{}, "learncosmos/TransferGold", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgBuyGold{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgSellGold{})
	registry.RegisterImplementations((*sdk.Msg)(nil), &MsgTransferGold{})
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

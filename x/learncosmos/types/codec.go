package types

import (
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/msgservice"
	channeltypes "github.com/cosmos/cosmos-sdk/x/ibc/core/04-channel/types"
	// this line is used by starport scaffolding # 1
)

func RegisterCodec(cdc *codec.LegacyAmino) {
	// this line is used by starport scaffolding # 2
	cdc.RegisterConcrete(&MsgBuyGold{}, "learncosmos/BuyGold", nil)
	cdc.RegisterConcrete(&MsgSellGold{}, "learncosmos/SellGold", nil)
	cdc.RegisterConcrete(&MsgTransferGold{}, "learncosmos/TransferGold", nil)
	cdc.RegisterConcrete(&channeltypes.MsgRecvPacket{}, "ibc/core/v1/MsgRecvPacket", nil)
}

func RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	// this line is used by starport scaffolding # 3
	registry.RegisterImplementations((*sdk.Msg)(nil),
		&MsgBuyGold{},
		&MsgSellGold{},
		&MsgTransferGold{},
		&channeltypes.MsgRecvPacket{},
	)

	msgservice.RegisterMsgServiceDesc(registry, &_Msg_serviceDesc)
}

var (
	amino     = codec.NewLegacyAmino()
	ModuleCdc = codec.NewAminoCodec(amino)
)

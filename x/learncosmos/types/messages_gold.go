package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgBuyGold{}

func NewMsgBuyGold(buyer string, amount uint64) *MsgBuyGold {
	return &MsgBuyGold{
		Buyer:  buyer,
		Amount: amount,
	}
}

func (msg MsgBuyGold) Route() string {
	return RouterKey
}

func (msg MsgBuyGold) Type() string {
	return "BuyGold"
}

func (msg *MsgBuyGold) GetSigners() []sdk.AccAddress {
	buyer, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{buyer}
}

func (msg *MsgBuyGold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgBuyGold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Buyer)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid buyer address")
	}

	return nil
}

var _ sdk.Msg = &MsgSellGold{}

func NewMsgSellGold(seller string, amount uint64) *MsgSellGold {
	return &MsgSellGold{
		Seller: seller,
		Amount: amount,
	}
}

func (msg MsgSellGold) Route() string {
	return RouterKey
}

func (msg MsgSellGold) Type() string {
	return "SellGold"
}

func (msg *MsgSellGold) GetSigners() []sdk.AccAddress {
	seller, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{seller}
}

func (msg *MsgSellGold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgSellGold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Seller)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid seller address")
	}

	return nil
}

var _ sdk.Msg = &MsgTransferGold{}

func NewMsgTransferGold(sender string, receiver string, amount uint64) *MsgTransferGold {
	return &MsgTransferGold{
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
	}
}

func (msg MsgTransferGold) Route() string {
	return RouterKey
}

func (msg MsgTransferGold) Type() string {
	return "TransferGold"
}

func (msg *MsgTransferGold) GetSigners() []sdk.AccAddress {
	seller, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{seller}
}

func (msg *MsgTransferGold) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgTransferGold) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address")
	}
	_, err = sdk.AccAddressFromBech32(msg.Receiver)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid receiver address")
	}

	return nil
}

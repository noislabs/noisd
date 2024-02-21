package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgClaimRewards{}

// msg types
const (
	TypeMsgClaimRewards = "claim_rewards"
)

func NewMsgClaimRewards(sender string) *MsgClaimRewards {
	return &MsgClaimRewards{
		Sender: sender,
	}
}

func (msg MsgClaimRewards) Route() string {
	return RouterKey
}

func (msg MsgClaimRewards) Type() string {
	return TypeMsgClaimRewards
}

func (msg MsgClaimRewards) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

func (msg MsgClaimRewards) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

func (msg MsgClaimRewards) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}

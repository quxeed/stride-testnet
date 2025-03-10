package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	utils "github.com/Stride-Labs/stride/utils"
)

const TypeMsgClaimUndelegatedTokens = "claim_undelegated_tokens"

var _ sdk.Msg = &MsgClaimUndelegatedTokens{}

func NewMsgClaimUndelegatedTokens(creator string, hostZone string, epoch uint64, sender string) *MsgClaimUndelegatedTokens {
	return &MsgClaimUndelegatedTokens{
		Creator:   creator,
		HostZoneId:  hostZone,
		Epoch: epoch,
		Sender: sender,
	}
}

func (msg *MsgClaimUndelegatedTokens) Route() string {
	return RouterKey
}

func (msg *MsgClaimUndelegatedTokens) Type() string {
	return TypeMsgClaimUndelegatedTokens
}

func (msg *MsgClaimUndelegatedTokens) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgClaimUndelegatedTokens) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgClaimUndelegatedTokens) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	// epoch much be greater than 0
	if msg.Epoch < 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "epoch must be positive")
	}
	// sender must be a valid stride address
	_, err = utils.AccAddressFromBech32(msg.Sender, "stride")
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}
	return nil
}
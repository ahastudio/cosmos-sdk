package types

import (
	crypto "github.com/tendermint/go-crypto"
)

type Msg interface {

	// Return the message type.
	// Must be alphanumeric or empty.
	Type() string

	// Get some property of the Msg.
	Get(key interface{}) (value interface{})

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []crypto.Address
}

type Tx interface {

	// Gets the Msg.
	GetMsg() Msg

	// The address that pays the base fee for this message.  The fee is
	// deducted before the Msg is processed.
	GetFeePayer() crypto.Address

	// Signatures returns the signature of signers who signed the Msg.
	// CONTRACT: Length returned is same as length of
	// pubkeys returned from MsgKeySigners, and the order
	// matches.
	// CONTRACT: If the signature is missing (ie the Msg is
	// invalid), then the corresponding signature is
	// .Empty().
	GetSignatures() []StdSignature
}

var _ Tx = (*StdTx)(nil)

type StdTx struct {
	Msg
	Signatures []StdSignature
}

func (tx StdTx) GetMsg() Msg                   { return tx.Msg }
func (tx StdTx) GetFeePayer() crypto.Address   { return tx.Signatures[0].PubKey.Address() }
func (tx StdTx) GetSignatures() []StdSignature { return tx.Signatures }

type TxDecoder func(txBytes []byte) (Tx, Error)

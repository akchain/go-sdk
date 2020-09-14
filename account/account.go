package account

import (
	"bytes"
	"encoding/gob"
	"encoding/hex"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
)

type Account struct {
	Xprv           chainkd.XPrv
	Xpub           chainkd.XPub
	AccountID      string
	AccountAlias   string
	Address        string
	ControlProgram string
}

func New(xprv chainkd.XPrv, xpub chainkd.XPub, accountID string,
	accountAlias string, address string, control string) *Account {
	return &Account{
		Xprv:           xprv,
		Xpub:           xpub,
		AccountID:      accountID,
		AccountAlias:   accountAlias,
		Address:        address,
		ControlProgram: control,
	}
}

func (a *Account) ToString() (string, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)

	if err := enc.Encode(a); err != nil {
		return "", err
	}

	return hex.EncodeToString(buf.Bytes()), nil
}

func LoadAcount(accountStr string) (*Account, error) {
	accountBytes, _ := hex.DecodeString(accountStr)
	buf := bytes.NewBuffer(accountBytes)
	var account Account

	dec := gob.NewDecoder(buf)

	if err := dec.Decode(&account); err != nil {
		return nil, err
	}

	return &account, nil
}

func (a *Account) Sign(data []byte) []byte {
	return a.Xprv.Sign(data)
}

func (a *Account) Verify(msg []byte, sig []byte) bool {
	return a.Xpub.Verify(msg, sig)
}

func (a *Account) GetPrivateKey() string {
	return a.Xprv.String()
}

func (a *Account) GetPublicKey() string {
	return a.Xpub.String()
}

func (a *Account) GetAddress() string {
	return a.Address
}

func (a *Account) GetControlProgram() string {
	return a.ControlProgram
}

func (a *Account) GetAccountID() string {
	return a.AccountID
}

func (a *Account) GetAccountAlias() string {
	return a.AccountAlias
}

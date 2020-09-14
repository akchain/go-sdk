package go_sdk

import (
	"akc/go-sdk/client"
	"bytes"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	"github.com/tyler-smith/go-bip39"
)

type AkchainSdk struct {
	client.Mgr
}

func NewAkchainSdk() *AkchainSdk {
	akcSdk := &AkchainSdk{}
	return akcSdk
}

func (a *AkchainSdk) GenerateMnemonicCodesStr() (string, error) {
	entropy, err := bip39.NewEntropy(256)
	if err != nil {
		return "", err
	}
	return bip39.NewMnemonic(entropy)
}

func (a *AkchainSdk) GetKeysFromMnemonic(mnemonicCodesStr string) (string, string, error) {
	seed := bip39.NewSeed(mnemonicCodesStr, "")
	xprv, xpub, err := chainkd.NewXKeys(bytes.NewBuffer(seed))
	if err != nil {
		return "", "", err
	}
	return xprv.String(), xpub.String(), nil
}

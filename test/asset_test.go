package test

import (
	"encoding/hex"
	"fmt"
	"github.com/akchain/go-sdk/utils"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCreateGetAsset(t *testing.T) {
	xpub := "815f9253e89d90b6daf5b9e128c802f72f077695cb25a1eb841dcc580f78f241a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690"
	alias := utils.RandString(10)
	asset, err := akcSdk.CreateAsset(xpub, alias)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", asset))
	assetInfo, err := akcSdk.GetAsset(asset.ID)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", assetInfo))
}

func TestIssueAsset(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAsset(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)
}

func TestIssueAssetAsync(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAssetAsync(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)
}

func TestUpdateAssetAlias(t *testing.T) {
	xpub := "815f9253e89d90b6daf5b9e128c802f72f077695cb25a1eb841dcc580f78f241a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690"
	alias := utils.RandString(20)
	asset, err := akcSdk.CreateAsset(xpub, alias)
	assert.Nil(t, err, err)
	newAlias := utils.RandString(20)
	err = akcSdk.UpdateAssetAlias(asset.ID, newAlias)
	assert.Nil(t, err, err)
	assetInfo, err := akcSdk.GetAsset(asset.ID)
	assert.Nil(t, err, err)
	assert.Equal(t, newAlias, assetInfo.Alias)
}

func TestRetireAsset(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAsset(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)

	txHash2, err := akcSdk.RetireAsset(2000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash2)
}

func TestRetireAssetAsync(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAsset(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)

	txHash2, err := akcSdk.RetireAssetAsync(2000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash2)
}

func TestTransferAsset(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAsset(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)

	receiver, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(receiver.GetAccountID())

	transferHash, err := akcSdk.TransferAsset(1000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", receiver.GetAddress(), account)
	assert.Nil(t, err, err)
	fmt.Println(transferHash)
}

func TestTransferAssetAsync(t *testing.T) {
	x, _ := hex.DecodeString("d0b006d3e95f2b57225da2cb5728e914bd5dbc73f3e22b4351bf66cdf06c9b50a38e85491b7b49dcd7452039a2cfaba40ede15652876ceadbf3de81a3715a690")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(account.GetAccountID())

	//asset 必须为此xpub发行的才能issue
	txHash, err := akcSdk.IssueAsset(10000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", "", account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)

	receiver, _ := akcSdk.NewAccountFromPrivateKey(xprv)
	fmt.Println(receiver.GetAccountID())

	transferHash, err := akcSdk.TransferAssetAsync(1000000, "e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b", receiver.GetAddress(), account)
	assert.Nil(t, err, err)
	fmt.Println(transferHash)
}

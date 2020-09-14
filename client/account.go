package client

import (
	"akc/go-sdk/account"
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"akc/go-sdk/utils"
	"crypto/rand"
	"encoding/json"
	"fmt"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
)

const (
	AccountAliasLength = 20
)

func (c *Mgr) ValidateAddress(address string) (*types.ValidateAddressInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.validateAddress(c.getNextQid(), address)
	if err != nil {
		return nil, err
	}
	info := &types.ValidateAddressInfo{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return info, nil
}

func (c *Mgr) NewAccountFromPrivateKey(xprv chainkd.XPrv) (*account.Account, error) {
	xpub := xprv.XPub()
	alias := utils.RandString(AccountAliasLength)
	accountID, err := c.createAccountID(xpub.String(), alias)
	if err != nil {
		return nil, err
	}

	address, control, err := c.createReceiver(accountID)
	if err != nil {
		return nil, err
	}

	return account.New(xprv, xpub, accountID, alias, address, control), nil
}

func (c *Mgr) NewAccount() (*account.Account, error) {
	xprv, xpub, err := chainkd.NewXKeys(rand.Reader)

	alias := utils.RandString(AccountAliasLength)
	accountID, err := c.createAccountID(xpub.String(), alias)
	if err != nil {
		return nil, err
	}

	address, control, err := c.createReceiver(accountID)
	if err != nil {
		return nil, err
	}
	return account.New(xprv, xpub, accountID, alias, address, control), nil
}

func (c *Mgr) ListAccountBalances(accountID string) (*[]types.AccountBalanceInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.listAccountBalances(c.getNextQid(), accountID)
	if err != nil {
		return nil, err
	}
	balances := &[]types.AccountBalanceInfo{}
	err = json.Unmarshal(data, &balances)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return balances, nil
}

func (c *Mgr) ListAccounts() (*[]types.CreateAccountInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.listAccounts(c.getNextQid())
	if err != nil {
		return nil, err
	}
	accounts := &[]types.CreateAccountInfo{}
	err = json.Unmarshal(data, &accounts)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return accounts, nil
}

func (c *Mgr) RecoveryAccount(xpubs []string) error {
	client := c.getClient()
	if client == nil {
		return fmt.Errorf(common.ErrClientNotFound)
	}

	_, err := client.recoveryAccount(c.getNextQid(), xpubs)
	return err
}

func (c *Mgr) ListPubKeys(accountID string) (*types.ListPublicKeysInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.listPubKeys(c.getNextQid(), accountID)
	if err != nil {
		return nil, err
	}
	pubKeys := &types.ListPublicKeysInfo{}
	err = json.Unmarshal(data, &pubKeys)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return pubKeys, nil
}

func (c *Mgr) GetAccountInfo(accountID string) (*[]types.AccountInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.getAccountInfo(c.getNextQid(), accountID)
	if err != nil {
		return nil, err
	}
	accountInfo := &[]types.AccountInfo{}
	err = json.Unmarshal(data, &accountInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return accountInfo, nil
}

func (c *Mgr) createAccountID(xpub string, alias string) (string, error) {
	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.createAccount(c.getNextQid(), xpub, alias)
	if err != nil {
		return "", err
	}
	a := &types.CreateAccountInfo{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return a.ID, nil
}

func (c *Mgr) createReceiver(accountID string) (string, string, error) {
	client := c.getClient()
	if client == nil {
		return "", "", fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.createReceiver(c.getNextQid(), accountID)
	if err != nil {
		return "", "", err
	}
	a := &types.AccountReceiverInfo{}
	err = json.Unmarshal(data, &a)
	if err != nil {
		return "", "", fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return a.Address, a.ControlProgram, nil
}

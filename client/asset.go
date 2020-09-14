package client

import (
	"akc/go-sdk/account"
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"encoding/json"
	"fmt"
)

func (c *Mgr) CreateAsset(xpub string, alias string) (*types.CreateAssetInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.createAsset(c.getNextQid(), xpub, alias)
	if err != nil {
		return nil, err
	}
	info := &types.CreateAssetInfo{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return info, nil
}

func (c *Mgr) IssueAsset(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, destAddress, singer, true, types.IssueAsset)
}

func (c *Mgr) IssueAssetAsync(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, destAddress, singer, false, types.IssueAsset)
}

func (c *Mgr) TransferAsset(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, destAddress, singer, true, types.TransferAsset)
}

func (c *Mgr) TransferAssetAsync(amount uint64, assetID string, destAddress string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, destAddress, singer, false, types.TransferAsset)
}

func (c *Mgr) RetireAsset(amount uint64, assetID string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, "", singer, true, types.RetireAsset)
}

func (c *Mgr) RetireAssetAsync(amount uint64, assetID string, singer *account.Account) (string, error) {
	return c.assetTx(amount, assetID, "", singer, false, types.RetireAsset)
}

func (c *Mgr) assetTx(amount uint64, assetID string, destAddress string, singer *account.Account, isSync bool, txType types.AssetTxType) (string, error) {
	buildReq := types.Transaction{}
	addr := destAddress

	//构建inpug action
	issueAction := make(map[string]interface{})
	issueAction["account_id"] = singer.GetAccountID()
	issueAction["amount"] = amount
	issueAction["asset_id"] = assetID

	if txType == types.IssueAsset {
		issueAction["type"] = types.IssueActionType
	} else {
		issueAction["type"] = types.TransferRetireActionType
	}

	buildReq.Actions = append(buildReq.Actions, issueAction)
	//构建输出
	controlAddressAction := make(map[string]interface{})
	controlAddressAction["amount"] = amount
	controlAddressAction["asset_id"] = assetID

	if txType != types.RetireAsset {
		controlAddressAction["type"] = types.ControlAddressActionType
		if addr == "" {
			addr = singer.GetAddress()
		}
		controlAddressAction["address"] = addr
	} else {
		controlAddressAction["type"] = types.RetireType
	}

	buildReq.Actions = append(buildReq.Actions, controlAddressAction)

	return c.submitAssetTx(buildReq, singer, isSync)
}

func (c *Mgr) GetAsset(id string) (*types.AssetInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getAsset(c.getNextQid(), id)
	if err != nil {
		return nil, err
	}
	info := &types.AssetInfo{}
	err = json.Unmarshal(data, &info)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return info, nil
}

func (c *Mgr) UpdateAssetAlias(id string, alias string) error {
	client := c.getClient()
	if client == nil {
		return fmt.Errorf(common.ErrClientNotFound)
	}
	return client.updateAssetAlias(c.getNextQid(), id, alias)
}

package client

import (
	"akc/go-sdk/account"
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"akc/go-sdk/utils"
	"encoding/json"
	"fmt"
)

func (c *Mgr) UploadContract(filePath string) (string, error) {
	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.uploadContract(c.getNextQid(), filePath)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func (c *Mgr) DeployContract(contractPath string, contractName string, args string, signer *account.Account) (string, error) {
	builtContract, err := c.buildContractTx(contractName, args, contractPath, signer.GetAccountID())
	if err != nil {
		return "", err
	}
	if err = utils.SignContract(signer.GetPrivateKey(), builtContract); err != nil {
		return "", err
	}
	return c.submitContract(builtContract)
}

func (c *Mgr) UpgradeContract(contractPath string, contractName string, signer *account.Account) (string, error) {
	newContract, err := c.upgradeContractTx(contractName, contractPath, signer.GetAccountID())
	if err != nil {
		return "", err
	}
	if err = utils.SignContract(signer.GetPrivateKey(), newContract); err != nil {
		return "", err
	}
	return c.submitContract(newContract)
}

func (c *Mgr) ListContracts(accountID string) (*[]types.ContractInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.listContracts(c.getNextQid(), accountID)
	if err != nil {
		return nil, err
	}
	contractInfo := &[]types.ContractInfo{}
	err = json.Unmarshal(data, &contractInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return contractInfo, nil
}

func (c *Mgr) InvokeMethod(contractName string, method string, args string, signer *account.Account) (string, error) {
	builtContract, err := c.invokeContractTx(signer.GetAccountID(), contractName, method, args)
	if err != nil {
		return "", err
	}
	if err = utils.SignContract(signer.GetPrivateKey(), builtContract); err != nil {
		return "", err
	}
	return c.submitContract(builtContract)
}

func (c *Mgr) QueryContract(contractName string, method string, args string) ([]*types.ContractResponse, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.queryContract(c.getNextQid(), contractName, method, args)
	if err != nil {
		return nil, err
	}
	queryResult := &types.QueryResult{}
	err = json.Unmarshal(data, &queryResult)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return queryResult.Responses, nil
}

package client

import (
	"akc/go-sdk/account"
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"akc/go-sdk/utils"
	"encoding/json"
	"fmt"
)

func (c *Mgr) GetNumUnconfirmedTx() (int, error) {
	client := c.getClient()
	if client == nil {
		return 0, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getNumUnconfirmedTx(c.getNextQid())
	if err != nil {
		return 0, err
	}
	tx := struct {
		Number int             `json:"n_txs"`
		Txs    json.RawMessage `json:"txs"`
	}{}
	err = json.Unmarshal(data, &tx)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return tx.Number, nil
}

func (c *Mgr) ListUnconfirmedTx() (*types.UnconfirmedTx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.listUnconfirmedTx(c.getNextQid())
	if err != nil {
		return nil, err
	}
	txs := &types.UnconfirmedTx{}
	err = json.Unmarshal(data, &txs)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return txs, nil
}

func (c *Mgr) submitAssetTx(tx types.Transaction, signer *account.Account, isSync bool) (string, error) {
	builtTx, err := c.buildTx(tx)
	if err != nil {
		return "", err
	}

	decodedTx, err := c.decodeRawTx(builtTx.Transaction)
	if err != nil {
		return "", err
	}

	err = utils.Sign(signer.GetPrivateKey(), builtTx, decodedTx)
	if err != nil {
		return "", err
	}

	subTx, err := c.signTx(builtTx)
	if err != nil {
		return "", err
	}

	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}

	hash, err := client.submitTx(c.getNextQid(), subTx, isSync)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (c *Mgr) buildTx(tx types.Transaction) (*types.Template, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.buildTx(c.getNextQid(), &tx)
	if err != nil {
		return nil, err
	}
	templ := &types.Template{}
	err = json.Unmarshal(data, &templ)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return templ, nil
}

func (c *Mgr) decodeRawTx(rawTransaction string) (*types.RawTx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.decodeRawTx(c.getNextQid(), rawTransaction)
	if err != nil {
		return nil, err
	}
	rawTx := &types.RawTx{}
	err = json.Unmarshal(data, &rawTx)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return rawTx, nil
}

func (c *Mgr) signTx(template *types.Template) (*types.SignedTemplate, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.signTx(c.getNextQid(), template)
	if err != nil {
		return nil, err
	}
	signed := &types.SignedTemplate{}
	err = json.Unmarshal(data, &signed)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return signed, nil
}

func (c *Mgr) SearchTx(tx string) (*types.Tx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.searchTx(c.getNextQid(), tx)
	if err != nil {
		return nil, err
	}
	txs := &types.Tx{}
	err = json.Unmarshal(data, &txs)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return txs, nil
}

func (c *Mgr) ListUTXOs() (*[]types.UTXO, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.listUTXOs(c.getNextQid())
	if err != nil {
		return nil, err
	}
	utxo := &[]types.UTXO{}
	err = json.Unmarshal(data, &utxo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return utxo, nil
}

func (c *Mgr) buildContractTx(contractName string, args string, contractPath string, accountID string) (*types.WasmTx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.buildContract(c.getNextQid(), contractName, args, contractPath, accountID)
	if err != nil {
		return nil, err
	}
	wasm := &types.WasmTx{}
	err = json.Unmarshal(data, &wasm)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return wasm, nil
}

func (c *Mgr) upgradeContractTx(contractName string, contractPath string, accountID string) (*types.WasmTx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.upgradeContract(c.getNextQid(), contractName, contractPath, accountID)
	if err != nil {
		return nil, err
	}
	wasm := &types.WasmTx{}
	err = json.Unmarshal(data, &wasm)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return wasm, nil
}

func (c *Mgr) invokeContractTx(accountID string, contractName string, methodName string, args string) (*types.WasmTx, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.invokeContract(c.getNextQid(), contractName, methodName, args, accountID)
	if err != nil {
		return nil, err
	}
	wasm := &types.WasmTx{}
	err = json.Unmarshal(data, &wasm)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}

	return wasm, nil
}

func (c *Mgr) submitContract(contractTx *types.WasmTx) (string, error) {
	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}

	data, err := client.submitContractTx(c.getNextQid(), contractTx)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

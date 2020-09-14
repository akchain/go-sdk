package client

import (
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"encoding/json"
	"fmt"
)

func (c *Mgr) GetVersion() (*types.VersionInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getVersion(c.getNextQid())
	if err != nil {
		return nil, err
	}
	version := &types.VersionInfo{}
	err = json.Unmarshal(data, &version)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return version, nil
}

func (c *Mgr) GetStatus() (*types.StatusInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getStatus(c.getNextQid())
	if err != nil {
		return nil, err
	}
	status := &types.StatusInfo{}
	err = json.Unmarshal(data, &status)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return status, nil
}

func (c *Mgr) GetGenesis() (*types.ResultGenesis, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getGenesis(c.getNextQid())
	if err != nil {
		return nil, err
	}
	genesis := &types.ResultGenesis{}
	err = json.Unmarshal(data, &genesis)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return genesis, nil
}

func (c *Mgr) GetConfig() (*types.BaseConfig, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getConfig(c.getNextQid())
	if err != nil {
		return nil, err
	}
	config := &types.BaseConfig{}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return config, nil
}

func (c *Mgr) ListPeers() (*types.ListPeersInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.listPeers(c.getNextQid())
	if err != nil {
		return nil, err
	}
	peers := &types.ListPeersInfo{}
	err = json.Unmarshal(data, &peers)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return peers, nil
}

func (c *Mgr) GetNetInfo() (*types.NetInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getNetInfo(c.getNextQid())
	if err != nil {
		return nil, err
	}
	netInfo := &types.NetInfo{}
	err = json.Unmarshal(data, &netInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return netInfo, nil
}

func (c *Mgr) GetNodeInfo() (*types.DefaultNodeInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getNodeInfo(c.getNextQid())
	if err != nil {
		return nil, err
	}
	nodeInfo := &types.DefaultNodeInfo{}
	err = json.Unmarshal(data, &nodeInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return nodeInfo, nil
}

func (c *Mgr) CheckHealth() error {
	client := c.getClient()
	if client == nil {
		return fmt.Errorf(common.ErrClientNotFound)
	}
	err := client.checkHealth(c.getNextQid())
	if err != nil {
		return err
	}
	return nil
}

func (c *Mgr) GetBlockHeight() (uint64, error) {
	client := c.getClient()
	if client == nil {
		return 0, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getBlockHeight(c.getNextQid())
	if err != nil {
		return 0, err
	}
	height := struct {
		BlockHeight uint64 `json:"block_height"`
	}{}
	err = json.Unmarshal(data, &height)
	if err != nil {
		return 0, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return height.BlockHeight, nil
}

func (c *Mgr) GetBlockHash() (string, error) {
	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getBlockHash(c.getNextQid())
	if err != nil {
		return "", err
	}
	hash := struct {
		BlockHash string `json:"block_hash"`
	}{}

	err = json.Unmarshal(data, &hash)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return hash.BlockHash, nil
}

func (c *Mgr) GetBlockHashByHeight(height uint64) (string, error) {
	client := c.getClient()
	if client == nil {
		return "", fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getBlockHashByHeight(c.getNextQid(), height)
	if err != nil {
		return "", err
	}
	hash := struct {
		BlockHash string `json:"block_hash"`
	}{}

	err = json.Unmarshal(data, &hash)
	if err != nil {
		return "", fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return hash.BlockHash, nil
}

func (c *Mgr) GetHardwareUsage() (*types.HardwareUsage, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getHardwareUsage(c.getNextQid())
	if err != nil {
		return nil, err
	}
	hardware := &types.HardwareUsage{}
	err = json.Unmarshal(data, &hardware)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return hardware, nil
}

func (c *Mgr) GetValidatorsByHeight(height uint64) (*types.HeightValidators, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getValidatorsByHeight(c.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	validators := &types.HeightValidators{}
	err = json.Unmarshal(data, &validators)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return validators, nil
}

func (c *Mgr) GetNodeTypeByHeight(height uint64) (*types.NodeType, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getNodeTypeByHeight(c.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	validator := &types.NodeType{}
	err = json.Unmarshal(data, &validator)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return validator, nil
}

func (c *Mgr) GetBlockHeader(minHeight uint64, maxHeight uint64) (*types.BlockChainInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getBlockHeader(c.getNextQid(), minHeight, maxHeight)
	if err != nil {
		return nil, err
	}
	headers := &types.BlockChainInfo{}
	err = json.Unmarshal(data, &headers)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return headers, nil
}

func (c *Mgr) GetBlockDetail(height uint64, hash string) (*types.BlockResp, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getBlockDetail(c.getNextQid(), height, hash)
	if err != nil {
		return nil, err
	}
	block := &types.BlockResp{}
	err = json.Unmarshal(data, &block)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return block, nil
}

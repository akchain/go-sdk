package client

import (
	"akc/go-sdk/common"
	"akc/go-sdk/types"
	"encoding/json"
	"fmt"
)

func (c *Mgr) GetConsensusState() (*types.ConsensusStateInfo, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getConsensusStateInfo(c.getNextQid())
	if err != nil {
		return nil, err
	}
	consensusStateInfo := &types.ConsensusStateInfo{}
	err = json.Unmarshal(data, &consensusStateInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return consensusStateInfo, nil
}

func (c *Mgr) DumpConsensusState() (*types.DumpConsensusState, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.dumpConsensusState(c.getNextQid())
	if err != nil {
		return nil, err
	}
	consensusStateInfo := &types.DumpConsensusState{}
	err = json.Unmarshal(data, &consensusStateInfo)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return consensusStateInfo, nil
}

func (c *Mgr) GetConsensusParamsByHeight(height uint64) (*types.ConsensusParams, error) {
	client := c.getClient()
	if client == nil {
		return nil, fmt.Errorf(common.ErrClientNotFound)
	}
	data, err := client.getConsensusParamsByHeight(c.getNextQid(), height)
	if err != nil {
		return nil, err
	}
	consensusParams := &types.ConsensusParams{}
	err = json.Unmarshal(data, &consensusParams)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal:%s error:%s", data, err)
	}
	return consensusParams, nil
}

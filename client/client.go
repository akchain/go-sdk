package client

import (
	"fmt"
	"sync/atomic"
)

type Mgr struct {
	rpc       *RpcClient  //Rpc client used the rpc api of akchain
	rest      *RestClient //Rest client used the rest api of akchain
	defClient AkchainClient
	qid       uint64
}

func (c *Mgr) NewRpcClient() *RpcClient {
	c.rpc = NewRpcClient()
	return c.rpc
}

func (c *Mgr) GetRpcClient() *RpcClient {
	return c.rpc
}

func (c *Mgr) NewRestClient() *RestClient {
	c.rest = NewRestClient()
	return c.rest
}

func (c *Mgr) GetRestClient() *RestClient {
	return c.rest
}

func (c *Mgr) SetDefaultClient(client AkchainClient) {
	c.defClient = client
}

func (c *Mgr) getClient() AkchainClient {
	if c.defClient != nil {
		return c.defClient
	}
	if c.rest != nil {
		return c.rest
	}
	if c.rpc != nil {
		return c.rpc
	}
	return nil
}

func (c *Mgr) getNextQid() string {
	return fmt.Sprintf("%d", atomic.AddUint64(&c.qid, 1))
}

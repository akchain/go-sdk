package client

import (
	"akc/go-sdk/types"
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

//JsonRpc version
const JSON_RPC_VERSION = "2.0"

//RpcClient for akchain rpc api
type RpcClient struct {
	addr       string
	httpClient *http.Client
}

func (r *RpcClient) listPubKeys(qid string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) recoveryAccount(qid string, xpubs []string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) queryContract(qid string, name string, method string, args string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) invokeContract(qid string, name string, method string, args string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listContracts(qid string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) upgradeContract(qid string, name string, path string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) submitContractTx(qid string, tx *types.WasmTx) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) buildContract(qid string, contractName string, args string, contractPath string, accountID string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) uploadContract(qid string, path string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getAccountInfo(qid string, accountID string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listUTXOs(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listAccounts(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getBlockDetail(qid string, height uint64, hash string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getBlockHeader(qid string, minHeight uint64, maxHeight uint64) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getConsensusParamsByHeight(qid string, height uint64) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getNodeTypeByHeight(qid string, height uint64) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getValidatorsByHeight(qid string, height uint64) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getBlockHashByHeight(qid string, height uint64) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listUnconfirmedTx(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) searchTx(qid string, tx string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) submitTx(qid string, tx *types.SignedTemplate, isSync bool) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) signTx(qid string, template *types.Template) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) decodeRawTx(qid string, transaction string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) buildTx(qid string, tx *types.Transaction) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listAccountBalances(qid string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) createReceiver(qid string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) createAccount(qid string, xpub string, alias string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) updateAssetAlias(qid string, id string, alias string) error {
	panic("implement me")
}

func (r *RpcClient) getAsset(qid string, id string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) createAsset(qid string, xpub string, alias string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) validateAddress(qid string, address string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getNumUnconfirmedTx(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getHardwareUsage(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getBlockHash(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getBlockHeight(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) dumpConsensusState(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) checkHealth(qid string) error {
	panic("implement me")
}

func (r *RpcClient) getConsensusStateInfo(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getNodeInfo(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getNetInfo(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) listPeers(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getConfig(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getGenesis(qid string) ([]byte, error) {
	panic("implement me")
}

func (r *RpcClient) getStatus(qid string) ([]byte, error) {
	panic("implement me")
}

//NewRpcClient return RpcClient instance
func NewRpcClient() *RpcClient {
	return &RpcClient{
		httpClient: &http.Client{
			Transport: &http.Transport{
				MaxIdleConnsPerHost:   5,
				DisableKeepAlives:     false, //enable keepalive
				IdleConnTimeout:       time.Second * 300,
				ResponseHeaderTimeout: time.Second * 300,
			},
			Timeout: time.Second * 300, //timeout for http response
		},
	}
}

//SetAddress set rpc server address.
func (r *RpcClient) SetAddress(addr string) *RpcClient {
	r.addr = addr
	return r
}

//SetHttpClient set http client to RpcClient. In most cases SetHttpClient is not necessary
func (r *RpcClient) SetHttpClient(httpClient *http.Client) *RpcClient {
	r.httpClient = httpClient
	return r
}

//GetVersion return the version of akchain
func (r *RpcClient) getVersion(qid string) ([]byte, error) {
	return r.sendRpcRequest(qid, RpcGetVersion, []interface{}{})
}

func (r *RpcClient) sendRpcRequest(qid, method string, params []interface{}) ([]byte, error) {
	rpcReq := &JsonRpcRequest{
		Version: JSON_RPC_VERSION,
		Id:      qid,
		Method:  method,
		Params:  params,
	}
	data, err := json.Marshal(rpcReq)
	if err != nil {
		return nil, fmt.Errorf("JsonRpcRequest json.Marsha error:%s", err)
	}
	resp, err := r.httpClient.Post(r.addr, "application/json", bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("http post request:%s error:%s", data, err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read rpc response body error:%s", err)
	}
	rpcRsp := &JsonRpcResponse{}
	err = json.Unmarshal(body, rpcRsp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal JsonRpcResponse:%s error:%s", body, err)
	}
	if rpcRsp.Error != 0 {
		return nil, fmt.Errorf("JsonRpcResponse error code:%d desc:%s result:%s", rpcRsp.Error, rpcRsp.Desc, rpcRsp.Result)
	}
	return rpcRsp.Result, nil
}

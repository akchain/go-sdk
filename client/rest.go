package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/akchain/go-sdk/common"
	"github.com/akchain/go-sdk/types"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
	"time"
)

type RestClient struct {
	addr       string
	httpClient *http.Client
}

func (r *RestClient) listPubKeys(qid string, id string) ([]byte, error) {
	reqPath := ListPubkeys
	var req = &struct {
		AccountID string `json:"account_id"`
	}{
		AccountID: id,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) recoveryAccount(qid string, xpubs []string) ([]byte, error) {
	reqPath := RecoveryAccount
	var req = &struct {
		Xpubs []string `json:"xpubs"`
	}{
		Xpubs: xpubs,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) queryContract(qid string, name string, method string, args string) ([]byte, error) {
	reqPath := WasmQuery
	var req = &struct {
		ContractName string `json:"contract_name"`
		Args         string `json:"args"`
		MethodName   string `json:"method_name"`
	}{
		ContractName: name,
		Args:         args,
		MethodName:   method,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) invokeContract(qid string, name string, method string, args string, id string) ([]byte, error) {
	reqPath := BuildWasmInvoke
	var req = &struct {
		Account      string `json:"account"`
		ContractName string `json:"contract_name"`
		Args         string `json:"args"`
		MethodName   string `json:"method_name"`
	}{Account: id,
		ContractName: name,
		Args:         args,
		MethodName:   method,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) listContracts(qid string, id string) ([]byte, error) {
	reqPath := ListContracts
	var req = &struct {
		AccountID string `json:"account_name"`
	}{AccountID: id}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) upgradeContract(qid string, name string, path string, id string) ([]byte, error) {
	reqPath := UpgradeWasmContract
	var req = &struct {
		Account      string `json:"account"`
		ContractName string `json:"contract_name"`
		Code         string `json:"code"`
	}{Account: id,
		ContractName: name,
		Code:         path,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) submitContractTx(qid string, tx *types.WasmTx) ([]byte, error) {
	reqPath := SubmitContractTx
	req := &struct {
		Txs *types.WasmTx `json:"transaction"`
	}{Txs: tx}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) buildContract(qid string, contractName string, args string, contractPath string, accountID string) ([]byte, error) {
	reqPath := BuildWasmContract
	var req = &struct {
		Account      string `json:"account"`
		ContractName string `json:"contract_name"`
		Args         string `json:"args"`
		Runtime      string `json:"runtime"`
		Code         string `json:"code"`
	}{Account: accountID,
		ContractName: contractName,
		Args:         args,
		Runtime:      "c",
		Code:         contractPath,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) uploadContract(qid string, path string) ([]byte, error) {
	reqPath := UploadContract
	return r.uploadFileRequest(path, reqPath)
}

func (r *RestClient) getAccountInfo(qid string, accountID string) ([]byte, error) {
	reqPath := ListAddresses
	var req = &struct {
		AccountID string `json:"account_id"`
	}{AccountID: accountID}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) listUTXOs(qid string) ([]byte, error) {
	reqPath := ListUtxos
	var req = &struct{}{}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) listAccounts(qid string) ([]byte, error) {
	reqPath := ListAccounts
	var req = &struct{}{}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getBlockDetail(qid string, height uint64, hash string) ([]byte, error) {
	reqPath := GetBlockDetail
	var req = &struct {
		BlockHeight uint64 `json:"block_height"`
		BlockHash   string `json:"block_hash"`
	}{BlockHeight: height, BlockHash: hash}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getBlockHeader(qid string, minHeight uint64, maxHeight uint64) ([]byte, error) {
	reqPath := GetBlockHeaders
	var req = &struct {
		MinHeight uint64 `json:"minHeight"`
		MaxHeight uint64 `json:"maxHeight"`
	}{MinHeight: minHeight, MaxHeight: maxHeight}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getConsensusParamsByHeight(qid string, height uint64) ([]byte, error) {
	reqPath := GetConsensusParamsByHeight
	req := &struct {
		Height uint64 `json:"height"`
	}{Height: height}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getNodeTypeByHeight(qid string, height uint64) ([]byte, error) {
	reqPath := GetNodeTypeByHeight
	req := &struct {
		Limit uint64 `json:"limit"`
	}{Limit: height}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getValidatorsByHeight(qid string, height uint64) ([]byte, error) {
	reqPath := GetValidatorsByHeight
	req := &struct {
		Limit uint64 `json:"limit"`
	}{Limit: height}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getBlockHashByHeight(qid string, height uint64) ([]byte, error) {
	reqPath := GetBlockHashByHeight
	req := &struct {
		Height uint64 `json:"height"`
	}{Height: height}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) listUnconfirmedTx(qid string) ([]byte, error) {
	reqPath := ListUnconfirmedTxs
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) searchTx(qid string, tx string) ([]byte, error) {
	reqPath := SearchTx
	req := &struct {
		TxHash string `json:"tx_hash"`
	}{TxHash: tx}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) submitTx(qid string, tx *types.SignedTemplate, isSync bool) ([]byte, error) {
	var reqPath = ""
	if isSync {
		reqPath = SubmitAssetTx
	} else {
		reqPath = SubmitAssetTxAsync
	}
	req := &struct {
		Tx string `json:"raw_transaction"`
	}{Tx: tx.Template.Transaction}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) signTx(qid string, template *types.Template) ([]byte, error) {
	reqPath := SignAssetTx
	req := &struct {
		Txs *types.Template `json:"transaction"`
	}{Txs: template}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) decodeRawTx(qid string, transaction string) ([]byte, error) {
	reqPath := DecodeRawAssetTx
	req := &struct {
		Tx string `json:"raw_transaction"`
	}{Tx: transaction}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) buildTx(qid string, tx *types.Transaction) ([]byte, error) {
	reqPath := BuildAssetTx
	return r.sendRestPostRequest(tx, reqPath)
}

func (r *RestClient) listAccountBalances(qid string, id string) ([]byte, error) {
	reqPath := ListBalances
	req := &struct {
		AccountID string `json:"account_id"`
	}{AccountID: id}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) createReceiver(qid string, id string) ([]byte, error) {
	reqPath := CreateAccountReceiver
	req := &struct {
		AccountID string `json:"account_id"`
	}{AccountID: id}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) createAccount(qid string, xpub string, alias string) ([]byte, error) {
	reqPath := CreateAccount
	req := &struct {
		RootXPubs []string `json:"root_xpubs"`
		Quorum    int      `json:"quorum"`
		Alias     string   `json:"alias"`
	}{RootXPubs: []string{xpub}, Quorum: 1, Alias: alias}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) updateAssetAlias(qid string, id string, alias string) error {
	reqPath := UpdateAssetAlias
	var req = &struct {
		ID       string `json:"id"`
		NewAlias string `json:"alias"`
	}{ID: id, NewAlias: alias}
	_, err := r.sendRestPostRequest(req, reqPath)
	return err
}

func (r *RestClient) getAsset(qid string, id string) ([]byte, error) {
	reqPath := GetAsset
	var req = &struct {
		ID string `json:"id"`
	}{ID: id,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) createAsset(qid string, xpub string, alias string) ([]byte, error) {
	reqPath := CreateAsset
	var req = &struct {
		RootXPubs []string `json:"root_xpubs"`
		Quorum    int      `json:"quorum"`
		Alias     string   `json:"alias"`
	}{RootXPubs: []string{xpub},
		Quorum: 1,
		Alias:  alias,
	}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) validateAddress(qid string, address string) ([]byte, error) {
	reqPath := ValidateAddress
	var req = &struct {
		Address string `json:"address"`
	}{Address: address}
	return r.sendRestPostRequest(req, reqPath)
}

func (r *RestClient) getNumUnconfirmedTx(qid string) ([]byte, error) {
	reqPath := GetNumberUnconfirmedTx
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getHardwareUsage(qid string) ([]byte, error) {
	reqPath := GetHardwareUsage
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getBlockHash(qid string) ([]byte, error) {
	reqPath := GetBlockHash
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getBlockHeight(qid string) ([]byte, error) {
	reqPath := GetBlockHeight
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) dumpConsensusState(qid string) ([]byte, error) {
	reqPath := DumpConsensusState
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) checkHealth(qid string) error {
	reqPath := CheckHealth
	_, err := r.sendRestGetRequest(reqPath)
	return err
}

func (r *RestClient) getConsensusStateInfo(qid string) ([]byte, error) {
	reqPath := GetConsensusStateInfo
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getNodeInfo(qid string) ([]byte, error) {
	reqPath := GetNodeInfo
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getNetInfo(qid string) ([]byte, error) {
	reqPath := GetNetInfo
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) listPeers(qid string) ([]byte, error) {
	reqPath := ListPeers
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getConfig(qid string) ([]byte, error) {
	reqPath := GetConfig
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getGenesis(qid string) ([]byte, error) {
	reqPath := GetGenesis
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) getStatus(qid string) ([]byte, error) {
	reqPath := GetStatus
	return r.sendRestGetRequest(reqPath)
}

//NewRpcClient return RpcClient instance
func NewRestClient() *RestClient {
	return &RestClient{
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

//SetAddress set rest server address. Simple http://localhost:20334
func (r *RestClient) SetAddress(addr string) *RestClient {
	r.addr = addr
	return r
}

//SetHttpClient set rest client to RestClient. In most cases SetHttpClient is not necessary
func (r *RestClient) SetHttpClient(httpClient *http.Client) *RestClient {
	r.httpClient = httpClient
	return r
}

func (r *RestClient) getVersion(qid string) ([]byte, error) {
	reqPath := GetVersion
	return r.sendRestGetRequest(reqPath)
}

func (r *RestClient) sendRestGetRequest(reqPath string, values ...*url.Values) ([]byte, error) {
	reqUrl, err := r.getRequestUrl(reqPath, values...)
	if err != nil {
		return nil, err
	}
	resp, err := r.httpClient.Get(reqUrl)
	if err != nil {
		return nil, fmt.Errorf("send http get request error:%s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("send http get request error:%s", resp.Status)
	}

	defer resp.Body.Close()
	return r.dealRestResponse(resp.Body)
}

func (r *RestClient) getAddress() (string, error) {
	if r.addr == "" {
		return "", fmt.Errorf("cannot get address, please add adrress first")
	}
	return r.addr, nil
}

func (r *RestClient) getRequestUrl(reqPath string, values ...*url.Values) (string, error) {
	addr, err := r.getAddress()
	if err != nil {
		return "", err
	}
	if !strings.HasPrefix(addr, "http") {
		addr = "http://" + addr
	}
	reqUrl, err := new(url.URL).Parse(addr)
	if err != nil {
		return "", fmt.Errorf("Parse address:%s error:%s", addr, err)
	}
	reqUrl.Path = reqPath
	if len(values) > 0 && values[0] != nil {
		reqUrl.RawQuery = values[0].Encode()
	}
	return reqUrl.String(), nil
}

func (r *RestClient) sendRestPostRequest(data interface{}, reqPath string, values ...*url.Values) ([]byte, error) {
	reqUrl, err := r.getRequestUrl(reqPath, values...)
	if err != nil {
		return nil, err
	}

	reqData, err := json.Marshal(data)
	if err != nil {
		return nil, fmt.Errorf("json.Marshal error:%s", err)
	}
	resp, err := r.httpClient.Post(reqUrl, "application/json", bytes.NewReader(reqData))
	if err != nil {
		return nil, fmt.Errorf("send http post request error:%s", err)
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("send http post request error:%s", resp.Status)
	}

	defer resp.Body.Close()
	return r.dealRestResponse(resp.Body)
}

func (r *RestClient) dealRestResponse(body io.Reader) ([]byte, error) {
	data, err := ioutil.ReadAll(body)
	if err != nil {
		return nil, fmt.Errorf("read http body error:%s", err)
	}
	restRsp := &RestfulResp{}
	err = json.Unmarshal(data, restRsp)
	if err != nil {
		return nil, fmt.Errorf("json.Unmarshal RestfulResp:%s error:%s", body, err)
	}

	if restRsp.Status != SUCCESS {
		return nil, fmt.Errorf("send http post request error:%s", restRsp.ErrorDetail)
	}

	return restRsp.Data, nil
}

func (r *RestClient) uploadFileRequest(filename string, reqPath string, values ...*url.Values) ([]byte, error) {
	if path.Ext(filename) != WasmFile {
		return nil, fmt.Errorf(common.ErrContractFileType)
	}

	reqUrl, err := r.getRequestUrl(reqPath, values...)
	if err != nil {
		return nil, err
	}

	bodyBuf := &bytes.Buffer{}
	bodyWriter := multipart.NewWriter(bodyBuf)

	fileWriter, err := bodyWriter.CreateFormFile("file", filepath.Base(filename))
	if err != nil {
		return nil, err
	}

	fh, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer fh.Close()

	_, err = io.Copy(fileWriter, fh)
	if err != nil {
		return nil, err
	}
	contentType := bodyWriter.FormDataContentType()
	bodyWriter.Close()
	resp, err := r.httpClient.Post(reqUrl, contentType, bodyBuf)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return nil, fmt.Errorf("send http post request error:%s", resp.Status)
	}

	defer resp.Body.Close()
	return r.dealRestResponse(resp.Body)
}

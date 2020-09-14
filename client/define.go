package client

import (
	"encoding/json"
	"github.com/akchain/go-sdk/types"
)

type AkchainClient interface {
	getVersion(qid string) ([]byte, error)
	getStatus(qid string) ([]byte, error)
	getGenesis(qid string) ([]byte, error)
	getConfig(qid string) ([]byte, error)
	listPeers(qid string) ([]byte, error)
	getNetInfo(qid string) ([]byte, error)
	getNodeInfo(qid string) ([]byte, error)
	getConsensusStateInfo(qid string) ([]byte, error)
	checkHealth(qid string) error
	dumpConsensusState(qid string) ([]byte, error)
	getBlockHeight(qid string) ([]byte, error)
	getBlockHash(qid string) ([]byte, error)
	getHardwareUsage(qid string) ([]byte, error)
	getNumUnconfirmedTx(qid string) ([]byte, error)
	validateAddress(qid string, address string) ([]byte, error)
	createAsset(qid string, xpub string, alias string) ([]byte, error)
	getAsset(qid string, id string) ([]byte, error)
	updateAssetAlias(qid string, id string, alias string) error
	createAccount(qid string, xpub string, alias string) ([]byte, error)
	createReceiver(qid string, id string) ([]byte, error)
	listAccountBalances(qid string, id string) ([]byte, error)
	buildTx(qid string, tx *types.Transaction) ([]byte, error)
	decodeRawTx(qid string, transaction string) ([]byte, error)
	signTx(qid string, template *types.Template) ([]byte, error)
	submitTx(qid string, tx *types.SignedTemplate, isSync bool) ([]byte, error)
	searchTx(qid string, tx string) ([]byte, error)
	listUnconfirmedTx(qid string) ([]byte, error)
	getBlockHashByHeight(qid string, height uint64) ([]byte, error)
	getValidatorsByHeight(qid string, height uint64) ([]byte, error)
	getNodeTypeByHeight(qid string, height uint64) ([]byte, error)
	getConsensusParamsByHeight(qid string, height uint64) ([]byte, error)
	getBlockHeader(qid string, minHeight uint64, maxHeight uint64) ([]byte, error)
	getBlockDetail(qid string, height uint64, hash string) ([]byte, error)
	listAccounts(qid string) ([]byte, error)
	listUTXOs(qid string) ([]byte, error)
	getAccountInfo(qid string, accountID string) ([]byte, error)
	uploadContract(qid string, path string) ([]byte, error)
	buildContract(qid string, contractName string, args string, contractPath string, accountID string) ([]byte, error)
	submitContractTx(qid string, tx *types.WasmTx) ([]byte, error)
	upgradeContract(qid string, name string, path string, id string) ([]byte, error)
	listContracts(qid string, id string) ([]byte, error)
	invokeContract(qid string, name string, method string, args string, id string) ([]byte, error)
	queryContract(qid string, name string, method string, args string) ([]byte, error)
	recoveryAccount(qid string, xpubs []string) ([]byte, error)
	listPubKeys(qid string, id string) ([]byte, error)
}

const (
	WasmFile = ".wasm"
)

const (
	RpcGetVersion = "version"
)

//JsonRpcRequest object in rpc
type JsonRpcRequest struct {
	Version string        `json:"jsonrpc"`
	Id      string        `json:"id"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

//JsonRpcResponse object response for JsonRpcRequest
type JsonRpcResponse struct {
	Error  int64           `json:"error"`
	Desc   string          `json:"desc"`
	Result json.RawMessage `json:"result"`
}

type WSRequest struct {
	Id     string
	Params map[string]interface{}
	ResCh  chan *WSResponse
}

type WSResponse struct {
	Id     string
	Action string
	Result json.RawMessage
	Error  int
	Desc   string
}

type WSAction struct {
	Action string
	Result interface{}
}

const (
	GetVersion                 = "/version"
	GetStatus                  = "/status"
	GetGenesis                 = "/genesis"
	GetConfig                  = "/get_config"
	ListPeers                  = "/list_peers"
	GetNetInfo                 = "/net_info"
	GetNodeInfo                = "/get_nodeinfo"
	GetConsensusStateInfo      = "/consensus_state"
	GetBlockHeight             = "/get_block_height"
	GetBlockHash               = "/get_block_hash"
	GetBlockHashByHeight       = "/get_block_hash_by_height"
	GetValidatorsByHeight      = "/get_validators_of_height"
	GetConsensusParamsByHeight = "/consensus_params"
	GetNodeTypeByHeight        = "/get_node_type"
	GetBlockHeaders            = "/blockchain"
	GetBlockDetail             = "/block_search"
	GetHardwareUsage           = "/get_hardware_target"
	GetNumberUnconfirmedTx     = "/num_unconfirmed_txs"
	DumpConsensusState         = "/dump_consensus_state"
	CheckHealth                = "/health"
	ValidateAddress            = "/validate_address"
	CreateAsset                = "/create_asset"
	GetAsset                   = "/get_asset"
	UpdateAssetAlias           = "/update_asset_alias"
	CreateAccountReceiver      = "/create_account_receiver"
	CreateAccount              = "/create_account"
	ListBalances               = "/list_balances"
	ListAccounts               = "/list_accounts"
	ListAddresses              = "/list_addresses"
	ListContracts              = "/get_contracts"
	ListUtxos                  = "/list_unspent_outputs"
	BuildAssetTx               = "/transfer_build"
	SignAssetTx                = "/transfer_sign"
	SubmitContractTx           = "/tx_submit"
	SubmitAssetTxAsync         = "/transfer_submit_async"
	SubmitAssetTx              = "/transfer_submit"
	SearchTx                   = "/tx_search"
	DecodeRawAssetTx           = "/decode_raw_transfer"
	ListUnconfirmedTxs         = "list_unconfirmed_transactions"
	UploadContract             = "/upload_contract"
	BuildWasmContract          = "/wasm_deploy"
	BuildWasmInvoke            = "/wasm_invoke"
	WasmQuery                  = "/wasm_query"
	UpgradeWasmContract        = "/wasm_upgrade"
	RecoveryAccount            = "/recovery_wallet"
	ListPubkeys                = "/list_pubkeys"
)

const (
	// SUCCESS indicates the rpc calling is successful.
	SUCCESS = "success"
	// FAIL indicated the rpc calling is failed.
	FAIL = "fail"
)

// Response describes the response standard.
type RestfulResp struct {
	Status      string          `json:"status,omitempty"`
	Code        string          `json:"code,omitempty"`
	Msg         string          `json:"msg,omitempty"`
	ErrorDetail string          `json:"error_detail,omitempty"`
	Data        json.RawMessage `json:"data,omitempty"`
}

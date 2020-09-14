package types

import bjson "github.com/bytom/bytom/encoding/json"

type WasmTx struct {
	TxSigWitness TxSigWitness   `json:"tx_sig_witness"`
	Transaction  string         `json:"raw_transaction"`
	DigestHash   bjson.HexBytes `json:"digesthash"`
}

type TxSigWitness struct {
	Type   string           `json:"type"`
	Quorum int              `json:"quorum"`
	Keys   []KeyID          `json:"keys"`
	Sigs   []bjson.HexBytes `json:"signatures"`
}

type SignedContractTx struct {
	Tx           *WasmTx `json:"tx_info"`
	SignComplete bool    `json:"sign_complete"`
}

type ContractInfo struct {
	ContractName string `json:"contract_name"`
	TxHash       string `json:"tx_hash"`
}

type QueryResult struct {
	Inputs    []*TxInputExt       `json:"txInputsExt"`
	Outputs   []*TxOutputExt      `json:"txOutputsExt"`
	Response  []string            `json:"responses"`
	Requests  []*InvokeRequest    `json:"contractRequests"`
	Responses []*ContractResponse `json:"contractResponses"`
}

type ContractResponse struct {
	Status  int32  `json:"status"`
	Message string `json:"message"`
	Body    string `json:"body"`
}

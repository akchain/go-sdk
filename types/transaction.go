package types

import (
	"encoding/json"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	bjson "github.com/bytom/bytom/encoding/json"
)

type Transaction struct {
	Tx                   *string                  `json:"base_transaction"`
	Actions              []map[string]interface{} `json:"actions"`
	ReservedDuration     bjson.Duration           `json:"reserved_duration"`
	IncludeOfBlockHeight uint64                   `json:"include_of_block_height"`
	MetaData             string                   `json:"meta_data"`
}

type Template struct {
	Transaction         string                `json:"raw_transaction"` //签名之后此字段提交
	SigningInstructions []*SigningInstruction `json:"signing_instructions"`
	Fee                 uint64                `json:"fee"`
	AllowAdditional     bool                  `json:"allow_additional_actions"`
}

type SigningInstruction struct {
	Position          int                 `json:"position"`
	WitnessComponents []*WitnessComponent `json:"witness_components"`
}

type WitnessComponent struct {
	Type   string           `json:"type"`
	Value  string           `json:"value"`
	Quorum int              `json:"quorum"`
	Keys   []KeyID          `json:"keys"`
	Sigs   []bjson.HexBytes `json:"signatures"`
}

type KeyID struct {
	Xpub           bjson.HexBytes   `json:"xpub"`
	DerivationPath []bjson.HexBytes `json:"derivation_path"`
}

type RawTx struct {
	TransferId string             `json:"transfer_id"`
	Version    uint64             `json:"version"`
	Size       uint64             `json:"size"`
	Inputs     []*AnnotatedInput  `json:"inputs"`
	Outputs    []*AnnotatedOutput `json:"outputs"`
	MuxID      string             `json:"mux_id"`
	MetaData   string             `json:"meta_data"`
}

type SignedTemplate struct {
	Template     *Template `json:"transaction"`
	SignComplete bool      `json:"sign_complete"`
}

type AnnotatedInput struct {
	Type             string           `json:"type"`
	AssetID          string           `json:"asset_id"`
	AssetAlias       string           `json:"asset_alias,omitempty"`
	AssetDefinition  *json.RawMessage `json:"asset_definition,omitempty"`
	Amount           uint64           `json:"amount"`
	IssuanceProgram  string           `json:"issuance_program,omitempty"`
	ControlProgram   string           `json:"control_program,omitempty"`
	Address          string           `json:"address,omitempty"`
	SpentOutputID    string           `json:"spent_output_id,omitempty"`
	AccountID        string           `json:"account_id,omitempty"`
	AccountAlias     string           `json:"account_alias,omitempty"`
	Arbitrary        string           `json:"arbitrary,omitempty"`
	InputID          string           `json:"input_id"`
	WitnessArguments interface{}      `json:"witness_arguments"`
}

type AnnotatedOutput struct {
	Type            string           `json:"type"`
	Id              string           `json:"id"`
	TransactionID   string           `json:"transaction_id,omitempty"`
	Position        int              `json:"position"`
	AssetID         string           `json:"asset_id"`
	AssetAlias      string           `json:"asset_alias,omitempty"`
	AssetDefinition *json.RawMessage `json:"asset_definition,omitempty"`
	Amount          uint64           `json:"amount"`
	AccountID       string           `json:"account_id,omitempty"`
	AccountAlias    string           `json:"account_alias,omitempty"`
	ControlProgram  string           `json:"control_program"`
	Address         string           `json:"address,omitempty"`
}

type Tx struct {
	Time            string           `json:"time,omitempty"`
	Height          *int64           `json:"height,omitempty"`
	TxHash          string           `json:"tx_hash"`
	Version         uint64           `json:"version"`
	Size            uint64           `json:"size"`
	Status          string           `json:"status"`
	Type            string           `json:"type"`
	TransferTxInfo  *TransferTxInfo  `json:"transfer_tx_info,omitempty"`
	WasmTxInfo      *WasmTxInfo      `json:"wasm_tx_info,omitempty"`
	ValidatorUpdate *ValidatorUpdate `json:"validator_update,omitempty"`
}

type TransferTxInfo struct {
	ID       string             `json:"id"`
	Inputs   []*AnnotatedInput  `json:"inputs"`
	Outputs  []*AnnotatedOutput `json:"outputs"`
	MuxID    string             `json:"mux_id"`
	MetaData string             `json:"meta_data"`
}

type ValidatorUpdate struct {
	AccountId string `json:"account_id"`
	PubKey    string `json:"node_pubkey"`
	Address   string `json:"node_address"`
	Power     int64  `json:"power"`
	Moniker   string `json:"moniker"`
	NodeId    string `json:"node_id"`
}

type WasmTxInfo struct {
	Desc             string           `json:"desc"`
	TxInputsExt      []TxInputExt     `json:"txInputsExt"`
	TxOutputsExt     []TxOutputExt    `json:"txOutputsExt"`
	ContractRequests []*InvokeRequest `json:"contractRequests"`
	Initiator        string           `json:"initiator"`
	TransferTxInfo   *TransferTxInfo  `json:"transfer_tx_info"`
}

type TxInputExt struct {
	Bucket    string         `json:"bucket"`
	Key       string         `json:"key"`
	RefTxid   bjson.HexBytes `json:"ref_txid"`
	RefOffset int32          `json:"ref_off_set"`
}

type TxOutputExt struct {
	Bucket string `json:"bucket"`
	Key    string `json:"key"`
	Value  []byte `json:"value"`
}

type InvokeRequest struct {
	ModuleName    string            `json:"module_name"`
	ContractName  string            `json:"contract_name"`
	MethodName    string            `json:"method_name"`
	Args          map[string]string `json:"args"`
	ResouceLimits []ResourceLimit   `json:"resource_limits"`
}

type UnconfirmedTx struct {
	Total uint32   `json:"total"`
	TXIDs []string `json:"tx_ids"`
}

type ResourceLimit struct {
	Type  string `json:"type"`
	Limit int64  `json:"limit"`
}

func (sw *WitnessComponent) Sign(prv chainkd.XPrv, message []byte) error {
	if len(sw.Sigs) < len(sw.Keys) {
		// Each key in sw.Keys may produce a signature in sw.Sigs. Make
		// sure there are enough slots in sw.Sigs and that we preserve any
		// sigs already present.
		newSigs := make([]bjson.HexBytes, len(sw.Keys))
		copy(newSigs, sw.Sigs)
		sw.Sigs = newSigs
	}
	for i, keyID := range sw.Keys {
		if len(sw.Sigs[i]) > 0 {
			// Already have a signature for this key
			continue
		}
		if string(keyID.Xpub) != string(prv.XPub().Bytes()) {
			continue
		}
		path := make([][]byte, len(keyID.DerivationPath))
		for i, p := range keyID.DerivationPath {
			path[i] = p
		}
		sigBytes, err := XSign(prv, path, message)
		if err != nil {
			continue
		}

		// This break is ordered to avoid signing transaction successfully only once for a multiple-sign account
		// that consist of different keys by the same password. Exit immediately when the signature is success,
		// it means that only one signature will be successful in the loop for this multiple-sign account.
		sw.Sigs[i] = sigBytes
		break
	}
	return nil
}

func XSign(prv chainkd.XPrv, path [][]byte, msg []byte) ([]byte, error) {
	if len(path) > 0 {
		prv = prv.Derive(path)
	}
	return prv.Sign(msg), nil
}

package types

import "encoding/json"

type AssetTxType int32

const (
	IssueAsset    AssetTxType = 0
	TransferAsset AssetTxType = 1
	RetireAsset   AssetTxType = 2
)

const (
	IssueActionType          = "issue"
	TransferRetireActionType = "spend_account"
	ControlAddressActionType = "control_address"
	RetireType               = "retire"
)

type CreateAssetInfo struct {
	ID              string           `json:"id"`
	Alias           string           `json:"alias,omitempty"`
	IssuanceProgram string           `json:"issuance_program"`
	Keys            []*AssetKey      `json:"keys"`
	Quorum          int              `json:"quorum"`
	Definition      *json.RawMessage `json:"definition"`
}

type AssetKey struct {
	RootXPub            string   `json:"root_xpub"`
	AssetPubkey         string   `json:"asset_pubkey"`
	AssetDerivationPath []string `json:"asset_derivation_path"`
}

type AssetInfo struct {
	Signer
	AssetID           string                 `json:"id"`
	Alias             string                 `json:"alias"`
	VMVersion         uint64                 `json:"vm_version"`
	IssuanceProgram   string                 `json:"issue_program"`
	RawDefinitionByte string                 `json:"raw_definition_byte"`
	DefinitionMap     map[string]interface{} `json:"definition"`
}

type Signer struct {
	Type       string   `json:"type"`
	XPubs      []string `json:"xpubs"`
	Quorum     int      `json:"quorum"`
	KeyIndex   uint64   `json:"key_index"`
	DeriveRule uint8    `json:"derive_rule"`
}

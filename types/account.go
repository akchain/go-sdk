package types

type ValidateAddressInfo struct {
	Valid   bool `json:"valid"`
	IsLocal bool `json:"is_local"`
}

type AccountBalanceInfo struct {
	AccountID       string                 `json:"account_id"`
	Alias           string                 `json:"account_alias"`
	AssetAlias      string                 `json:"asset_alias"`
	AssetID         string                 `json:"asset_id"`
	Amount          uint64                 `json:"amount"`
	AssetDefinition map[string]interface{} `json:"asset_definition"`
}

type CreateAccountInfo struct {
	ID         string   `json:"id"`
	Alias      string   `json:"alias,omitempty"`
	XPubs      []string `json:"xpubs"`
	Quorum     int      `json:"quorum"`
	KeyIndex   uint64   `json:"key_index"`
	DeriveRule uint8    `json:"derive_rule"`
}

type AccountReceiverInfo struct {
	ControlProgram string `json:"control_program,omitempty"`
	Address        string `json:"address,omitempty"`
}

type AccountInfo struct {
	AccountAlias   string `json:"account_alias,omitempty"`
	AccountID      string `json:"account_id,omitempty"`
	Address        string `json:"address,omitempty"`
	ControlProgram string `json:"control_program,omitempty"`
	Change         bool   `json:"change,omitempty"`
	KeyIndex       uint64 `json:"key_index,omitempty"`
}

type ListPublicKeysInfo struct {
	RootXPub    string       `json:"root_xpub"`
	PubKeyInfos []PubKeyInfo `json:"pubkey_infos"`
}

type PubKeyInfo struct {
	Pubkey string   `json:"pubkey"`
	Path   []string `json:"derivation_path"`
}

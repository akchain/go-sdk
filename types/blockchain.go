package types

import (
	"encoding/json"
	"time"
)

type VersionInfo struct {
	Version    string `json:"version"`
	Update     uint16 `json:"update"` // 0: no update; 1: small update; 2: significant update
	NewVersion string `json:"new_version"`
}

type DefaultNodeInfoOther struct {
	TxIndex    string `json:"tx_index"`
	RPCAddress string `json:"rpc_address"`
}

type ProtocolVersion struct {
	P2P   uint64 `json:"p2p"`
	Block uint64 `json:"block"`
	App   uint64 `json:"app"`
}

type DefaultNodeInfo struct {
	ProtocolVersion ProtocolVersion `json:"protocol_version"`

	// Authenticate
	// TODO: replace with NetAddress
	ID_        string `json:"id"`          // authenticated identifier
	ListenAddr string `json:"listen_addr"` // accepting incoming

	// Check compatibility.
	// Channels are HexBytes so easier to read as JSON
	Network  string `json:"network"`  // network/chain ID
	Version  string `json:"version"`  // major.minor.revision
	Channels string `json:"channels"` // channels this node knows about

	// ASCIIText fields
	Moniker string               `json:"moniker"` // arbitrary moniker
	Other   DefaultNodeInfoOther `json:"other"`   // other application specific data
}

type SyncInfo struct {
	LatestBlockHash   []byte    `json:"latest_block_hash"`
	LatestAppHash     []byte    `json:"latest_app_hash"`
	LatestBlockHeight int64     `json:"latest_block_height"`
	LatestBlockTime   time.Time `json:"latest_block_time"`
	CatchingUp        bool      `json:"catching_up"`
}

type StatusInfo struct {
	NodeInfo      DefaultNodeInfo `json:"node_info"`
	SyncInfo      SyncInfo        `json:"sync_info"`
	ValidatorInfo ValidatorInfo   `json:"validator_info"`
}

type ValidatorInfo struct {
	Address     []byte `json:"address"`
	PubKey      string `json:"pub_key"`
	VotingPower int64  `json:"voting_power"`
}

// BlockSizeParams define limits on the block size.
type BlockSizeParams struct {
	MaxBytes int64 `json:"max_bytes"`
	MaxGas   int64 `json:"max_gas"`
}

// EvidenceParams determine how we handle evidence of malfeasance
type EvidenceParams struct {
	MaxAge int64 `json:"max_age"` // only accept new evidence more recent than this
}

// ValidatorParams restrict the public key types validators can use.
// NOTE: uses ABCI pubkey naming, not Amino routes.
type ValidatorParams struct {
	PubKeyTypes []string `json:"pub_key_types"`
}

type ConsensusParamsInfo struct {
	BlockSize BlockSizeParams `json:"block_size"`
	Evidence  EvidenceParams  `json:"evidence"`
	Validator ValidatorParams `json:"validator"`
}

type GenesisValidator struct {
	Address string `json:"address"`
	PubKey  string `json:"pub_key"`
	Power   int64  `json:"power"`
	Name    string `json:"name"`
}

type GenesisDoc struct {
	GenesisTime     time.Time            `json:"genesis_time"`
	ChainID         string               `json:"chain_id"`
	ConsensusParams *ConsensusParamsInfo `json:"consensus_params,omitempty"`
	Validators      []GenesisValidator   `json:"validators,omitempty"`
	AppHash         string               `json:"app_hash"`
	AppState        json.RawMessage      `json:"app_state,omitempty"`
	NativeAddress   string               `json:"native_address"` //native asset address
}

type ResultGenesis struct {
	Genesis GenesisDoc `json:"genesis"`
}

type BaseConfig struct {
	// The root directory for all data.
	// This should be set in viper so it can unmarshal into this struct
	RootDir    string `mapstructure:"home"`
	ApiAddress string `mapstructure:"api_addr"`
	// Database backend: leveldb | memdb
	DBBackend string `mapstructure:"db_backend"`
	// Database directory
	DBPath string `mapstructure:"db_dir"`
	// Keystore directory
	KeysPath string `mapstructure:"keys_dir"`
}

type ListPeersInfo struct {
	Peers []Peer
}

type Peer struct {
	NodeInfo         DefaultNodeInfo  `json:"node_info"`
	IsOutbound       bool             `json:"is_outbound"`
	ConnectionStatus ConnectionStatus `json:"connection_status"`
}

type ConnectionStatus struct {
	Duration    time.Duration
	SendMonitor MonitorStatus
	RecvMonitor MonitorStatus
	Channels    []ChannelStatus
}

type ChannelStatus struct {
	ID                byte
	SendQueueCapacity int
	SendQueueSize     int
	Priority          int
	RecentlySent      int64
}

type MonitorStatus struct {
	Active   bool          // Flag indicating an active transfer
	Start    time.Time     // Transfer start time
	Duration time.Duration // Time period covered by the statistics
	Idle     time.Duration // Time since the last transfer of at least 1 byte
	Bytes    int64         // Total number of bytes transferred
	Samples  int64         // Total number of samples taken
	InstRate int64         // Instantaneous transfer rate
	CurRate  int64         // Current transfer rate (EMA of InstRate)
	AvgRate  int64         // Average transfer rate (Bytes / Duration)
	PeakRate int64         // Maximum instantaneous transfer rate
	BytesRem int64         // Number of bytes remaining in the transfer
	TimeRem  time.Duration // Estimated time to completion
	Progress uint32        // Overall transfer progress
}

type NetInfo struct {
	Listening    bool         `json:"listening"`
	Syncing      bool         `json:"syncing"`
	PeerCount    int          `json:"peer_count"`
	CurrentBlock uint64       `json:"current_block"`
	HighestBlock uint64       `json:"highest_block"`
	NetWorkID    string       `json:"network_id"`
	Version      *VersionInfo `json:"version_info"`
	Listeners    []string     `json:"listeners"`
	Peers        []Peer       `json:"peers"`
}

type HardwareUsage struct {
	CpuUsageRate string `json:"cpu_usage_rate"`
	MemUsageRate string `json:"mem_usage_rate"`
}

type BlockChainInfo struct {
	LastHeight int64        `json:"last_height"`
	BlockMetas []*BlockMeta `json:"block_metas"`
}

type BlockMeta struct {
	BlockID BlockID `json:"block_id"` // the block hash and partsethash
	Header  Header  `json:"header"`   // The block's Header
}

type BlockID struct {
	Hash        string        `json:"hash"`
	PartsHeader PartSetHeader `json:"parts"`
}

type PartSetHeader struct {
	Total int    `json:"total"`
	Hash  string `json:"hash"`
}

type Consensus struct {
	Block uint64 `json:"block"`
	App   uint64 `json:"app"`
}

type Header struct {
	// basic block info
	Version  Consensus `json:"version"`
	ChainID  string    `json:"chain_id"`
	Height   int64     `json:"height"`
	Time     time.Time `json:"time"`
	NumTxs   int64     `json:"num_txs"`
	TotalTxs int64     `json:"total_txs"`

	// prev block info
	LastBlockID BlockID `json:"last_block_id"`

	// hashes of block data
	LastCommitHash string `json:"last_commit_hash"` // commit from validators from the last block
	DataHash       string `json:"data_hash"`        // transactions

	// hashes from the app output from the prev block
	ValidatorsHash     string `json:"validators_hash"`      // validators for the current block
	NextValidatorsHash string `json:"next_validators_hash"` // validators for the next block
	ConsensusHash      string `json:"consensus_hash"`       // consensus params for current block
	AppHash            string `json:"app_hash"`             // state after txs from the previous block
	LastResultsHash    string `json:"last_results_hash"`    // root hash of all results from the txs from the previous block

	// consensus info
	EvidenceHash    string `json:"evidence_hash"`    // evidence included in the block
	ProposerAddress string `json:"proposer_address"` // original proposer of the block
}

type BlockResp struct {
	BlockHeaderResp *BlockHeaderResp `json:"block_header"`
	Transactions    []*BlockTx       `json:"block_txs"`
}

type BlockHeaderResp struct {
	Hash              string `json:"hash"`
	Version           uint64 `json:"version"` // The version of the block.
	ChainID           string `json:"chain_id"`
	Height            int64  `json:"height"` // The height of the block.
	Time              string `json:"time"`
	PreviousBlockHash string `json:"previous_block_hash"` // The hash of the previous block.
	AppHash           string `json:"app_hash"`            // state after txs from the previous block
	NumTxs            int64  `json:"num_txs"`
	TotalTxs          int64  `json:"total_txs"`
	DataHash          string `json:"transaction_merkle_root"` // transactions
	LastResultsHash   string `json:"transaction_status_hash"`
	ProposerAddress   string `json:"proposer_address"` // original proposer of the block
	Size              uint64 `json:"size"`
}

type BlockTx struct {
	TxHash          string           `json:"tx_hash"`
	Version         uint64           `json:"version"`
	Size            uint64           `json:"size"`
	Status          string           `json:"status"`
	Type            string           `json:"type"`
	TransferTxInfo  *TransferTxInfo  `json:"transfer_tx_info,omitempty"`
	WasmTxInfo      *WasmTxInfo      `json:"wasm_tx_info,omitempty"`
	ValidatorUpdate *ValidatorUpdate `json:"validator_update,omitempty"`
}

type UTXO struct {
	Alias               string `json:"account_alias"`
	OutputID            string `json:"id"`
	AssetID             string `json:"asset_id"`
	AssetAlias          string `json:"asset_alias"`
	Amount              uint64 `json:"amount"`
	AccountID           string `json:"account_id"`
	Address             string `json:"address"`
	ControlProgramIndex uint64 `json:"control_program_index"`
	Program             string `json:"program"`
	SourceID            string `json:"source_id"`
	SourcePos           uint64 `json:"source_pos"`
	ValidHeight         uint64 `json:"valid_height"`
	Change              bool   `json:"change"`
	DeriveRule          uint8  `json:"derive_rule"`
}

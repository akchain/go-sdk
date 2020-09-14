package types

type ConsensusStateInfo struct {
	RoundState RoundState `json:"round_state"`
}

type HeightValidators struct {
	BlockHeight     int64 `json:"block_height"`
	Validators      []Validator
	ProposerAddress string `json:"proposer_address"`
}

type ConsensusParams struct {
	BlockHeight     int64               `json:"block_height"`
	ConsensusParams ConsensusParamsInfo `json:"consensus_params"`
}

type Validator struct {
	Address     string `json:"address"`
	PubKey      string `json:"pub_key"`
	VotingPower int    `json:"voting_power"`
	Accum       int    `json:"accum"`
}

type NodeType struct {
	Validator bool   `json:"validator"`
	Address   string `json:"validator_address"`
}

type RoundState struct {
	HeightRoundStep   string       `json:"height/round/step"`
	StartTime         string       `json:"start_time"`
	ProposalBlockHash string       `json:"proposal_block_hash"`
	LockedBlockHash   string       `json:"locked_block_hash"`
	ValidBlockHash    string       `json:"valid_block_hash"`
	HeightVoteSet     []HeightVote `json:"height_vote_set"`
}

type HeightVote struct {
	Round              string   `json:"round"`
	Prevotes           []string `json:"prevotes"`
	PrevotesBitArray   string   `json:"prevotes_bit_array"`
	Precommits         []string `json:"precommits"`
	PrecommitsBitArray string   `json:"precommits_bit_array"`
}

type DumpConsensusState struct {
	Peers []struct {
		NodeAddress string `json:"node_address"`
		PeerState   struct {
			RoundState struct {
				CatchupCommit            string      `json:"catchup_commit"`
				CatchupCommitRound       string      `json:"catchup_commit_round"`
				Height                   string      `json:"height"`
				LastCommit               interface{} `json:"last_commit"`
				LastCommitRound          string      `json:"last_commit_round"`
				Precommits               string      `json:"precommits"`
				Prevotes                 string      `json:"prevotes"`
				Proposal                 bool        `json:"proposal"`
				ProposalBlockParts       interface{} `json:"proposal_block_parts"`
				ProposalBlockPartsHeader struct {
					Hash  string `json:"hash"`
					Total string `json:"total"`
				} `json:"proposal_block_parts_header"`
				ProposalPol      string `json:"proposal_pol"`
				ProposalPolRound string `json:"proposal_pol_round"`
				Round            string `json:"round"`
				StartTime        string `json:"start_time"`
				Step             int64  `json:"step"`
			} `json:"round_state"`
			Stats struct {
				BlockParts string `json:"block_parts"`
				Votes      string `json:"votes"`
			} `json:"stats"`
		} `json:"peer_state"`
	} `json:"peers"`
	RoundState struct {
		CommitRound    string      `json:"commit_round"`
		CommitTime     string      `json:"commit_time"`
		Height         string      `json:"height"`
		LastCommit     interface{} `json:"last_commit"`
		LastValidators struct {
			Proposer   interface{}   `json:"proposer"`
			Validators []interface{} `json:"validators"`
		} `json:"last_validators"`
		LockedBlock        interface{} `json:"locked_block"`
		LockedBlockParts   interface{} `json:"locked_block_parts"`
		LockedRound        string      `json:"locked_round"`
		Proposal           interface{} `json:"proposal"`
		ProposalBlock      interface{} `json:"proposal_block"`
		ProposalBlockParts interface{} `json:"proposal_block_parts"`
		Round              string      `json:"round"`
		StartTime          string      `json:"start_time"`
		Step               int64       `json:"step"`
		ValidBlock         interface{} `json:"valid_block"`
		ValidBlockParts    interface{} `json:"valid_block_parts"`
		ValidRound         string      `json:"valid_round"`
		Validators         struct {
			Proposer struct {
				Accum   string `json:"accum"`
				Address string `json:"address"`
				PubKey  struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"pub_key"`
				VotingPower string `json:"voting_power"`
			} `json:"proposer"`
			Validators []struct {
				Accum   string `json:"accum"`
				Address string `json:"address"`
				PubKey  struct {
					Type  string `json:"type"`
					Value string `json:"value"`
				} `json:"pub_key"`
				VotingPower string `json:"voting_power"`
			} `json:"validators"`
		} `json:"validators"`
		Votes []struct {
			Precommits         []string `json:"precommits"`
			PrecommitsBitArray string   `json:"precommits_bit_array"`
			Prevotes           []string `json:"prevotes"`
			PrevotesBitArray   string   `json:"prevotes_bit_array"`
			Round              string   `json:"round"`
		} `json:"votes"`
	} `json:"round_state"`
}

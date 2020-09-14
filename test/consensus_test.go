package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//{RoundState:{Height_Round_Step:8146/0/2
//StartTime:2020-07-24T08:59:17.920958914Z
//ProposalBlockHash:
//LockedBlockHash:
//ValidBlockHash:
//HeightVoteSet:[{Round:0 Prevotes:[nil-Vote nil-Vote nil-Vote nil-Vote] PrevotesBitArray:BA{4:____} 0/40 = 0.00 Precommits:[nil-Vote nil-Vote nil-Vote nil-Vote] PrecommitsBitArray:BA{4:____} 0/40 = 0.00} {Round:1 Prevotes:[nil-Vote nil-Vote nil-Vote nil-Vote] PrevotesBitArray:BA{4:____} 0/40 = 0.00 Precommits:[nil-Vote nil-Vote nil-Vote nil-Vote] PrecommitsBitArray:BA{4:____} 0/40 = 0.00}]}}
func TestConsensus(t *testing.T) {
	consensusRes, err := akcSdk.GetConsensusState()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", consensusRes))
}

//{Peers:[{NodeAddress:0.0.0.0:26656
//PeerState:{RoundState:{CatchupCommit:____
//CatchupCommitRound:-1
//Height:8148
//LastCommit:xxxx LastCommitRound:0
//Precommits:____ Prevotes:____ Proposal:false ProposalBlockParts:<nil> ProposalBlockPartsHeader:{Hash: Total:0} ProposalPol:____ ProposalPolRound:-1 Round:0 StartTime:2020-07-24T08:59:44.938270712Z Step:1} Stats:{BlockParts:12615 Votes:16244}}} {NodeAddress:0.0.0.0:26656 PeerState:{RoundState:{CatchupCommit:____ CatchupCommitRound:-1 Height:8148 LastCommit:xxxx LastCommitRound:0 Precommits:____ Prevotes:____ Proposal:false ProposalBlockParts:<nil> ProposalBlockPartsHeader:{Hash: Total:0} ProposalPol:____ ProposalPolRound:-1 Round:0 StartTime:2020-07-24T08:59:44.875928604Z Step:1} Stats:{BlockParts:9401 Votes:17416}}} {NodeAddress:0.0.0.0:26656 PeerState:{RoundState:{CatchupCommit:____ CatchupCommitRound:-1 Height:8148 LastCommit:xxxx LastCommitRound:0 Precommits:____ Prevotes:____ Proposal:false ProposalBlockParts:<nil> ProposalBlockPartsHeader:{Hash: Total:0} ProposalPol:____ ProposalPolRound:-1 Round:0 StartTime:2020-07-24T08:59:44.876495099Z Step:1} Stats:{BlockParts:9061 Votes:17188}}}] RoundState:{CommitRound:-1 CommitTime:2020-07-24T08:59:44.776438339Z Height:8148 LastCommit:map[peer_maj_23s:map[] votes:[Vote{0:13B3FCAB7437 8147/00/2(Precommit) AB2C75BF3B69 E8649C2E0880 @ 2020-07-24T08:59:44.677570569Z} Vote{1:6AFF0D270B91 8147/00/2(Precommit) AB2C75BF3B69 958F13F92B09 @ 2020-07-24T08:59:44.690358232Z} Vote{2:8E50C097641C 8147/00/2(Precommit) AB2C75BF3B69 BD15DBEC75CF @ 2020-07-24T08:59:44.679157905Z} Vote{3:D2AC87065C3B 8147/00/2(Precommit) AB2C75BF3B69 D91C11E7CB36 @ 2020-07-24T08:59:44.683677486Z}] votes_bit_array:BA{4:xxxx} 40/40 = 1.00] LastValidators:{Proposer:map[accum:-10 address:8E50C097641C1943F8825FE8257BD08668DA3D81 pub_key:map[type:tendermint/PubKeyEd25519 value:nU2CcsC8R1/Kgavav53UO3TlmMCUwmu2XQN5Ylu8S8o=] voting_power:10] Validators:[map[accum:-10 address:13B3FCAB74376D4F22B8C9513AEB715E7E81FB42 pub_key:map[type:tendermint/PubKeyEd25519 value:GFPBEhi9hifP+jv1lmo3JP0WUGd3dTmlJK/ptlfPL+g=] voting_power:10] map[accum:-10 address:6AFF0D270B91B05881F28766C0047414BCE11C71 pub_key:map[type:tendermint/PubKeyEd25519 value:YtKS2TLQ18dokoEhYsl5iGY1chjva6wFwFDI3lyO+oU=] voting_power:10] map[accum:-10 address:8E50C097641C1943F8825FE8257BD08668DA3D81 pub_key:map[type:tendermint/PubKeyEd25519 value:nU2CcsC8R1/Kgavav53UO3TlmMCUwmu2XQN5Ylu8S8o=] voting_power:10] map[accum:30 address:D2AC87065C3B306EF5204C895EE16F7B11FD834C pub_key:map[type:tendermint/PubKeyEd25519 value:OIpNEqTaAHRd336qZZEmeYFori7Jp4f/HQiM4M8RhM0=] voting_power:10]]} LockedBlock:<nil> LockedBlockParts:<nil> LockedRound:-1 Proposal:<nil> ProposalBlock:<nil> ProposalBlockParts:<nil> Round:0 StartTime:2020-07-24T08:59:45.776438339Z Step:2 ValidBlock:<nil> ValidBlockParts:<nil> ValidRound:-1 Validators:{Proposer:{Accum:0 Address:D2AC87065C3B306EF5204C895EE16F7B11FD834C PubKey:{Type:tendermint/PubKeyEd25519 Value:OIpNEqTaAHRd336qZZEmeYFori7Jp4f/HQiM4M8RhM0=} VotingPower:10} Validators:[{Accum:0 Address:13B3FCAB74376D4F22B8C9513AEB715E7E81FB42 PubKey:{Type:tendermint/PubKeyEd25519 Value:GFPBEhi9hifP+jv1lmo3JP0WUGd3dTmlJK/ptlfPL+g=} VotingPower:10} {Accum:0 Address:6AFF0D270B91B05881F28766C0047414BCE11C71 PubKey:{Type:tendermint/PubKeyEd25519 Value:YtKS2TLQ18dokoEhYsl5iGY1chjva6wFwFDI3lyO+oU=} VotingPower:10} {Accum:0 Address:8E50C097641C1943F8825FE8257BD08668DA3D81 PubKey:{Type:tendermint/PubKeyEd25519 Value:nU2CcsC8R1/Kgavav53UO3TlmMCUwmu2XQN5Ylu8S8o=} VotingPower:10} {Accum:0 Address:D2AC87065C3B306EF5204C895EE16F7B11FD834C PubKey:{Type:tendermint/PubKeyEd25519 Value:OIpNEqTaAHRd336qZZEmeYFori7Jp4f/HQiM4M8RhM0=} VotingPower:10}]} Votes:[{Precommits:[nil-Vote nil-Vote nil-Vote nil-Vote] PrecommitsBitArray:BA{4:____} 0/40 = 0.00 Prevotes:[nil-Vote nil-Vote nil-Vote nil-Vote] PrevotesBitArray:BA{4:____} 0/40 = 0.00 Round:0} {Precommits:[nil-Vote nil-Vote nil-Vote nil-Vote] PrecommitsBitArray:BA{4:____} 0/40 = 0.00 Prevotes:[nil-Vote nil-Vote nil-Vote nil-Vote] PrevotesBitArray:BA{4:____} 0/40 = 0.00 Round:1}]}}
func TestDumpConsensus(t *testing.T) {
	dumpConsensusRes, err := akcSdk.DumpConsensusState()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", dumpConsensusRes))
}

//{BlockHeight:1000
//ConsensusParams:{BlockSize:{MaxBytes:22020096 MaxGas:-1} Evidence:{MaxAge:100000} Validator:{PubKeyTypes:[ed25519]}}}
func TestConsensusParams(t *testing.T) {
	consensusParams, err := akcSdk.GetConsensusParamsByHeight(1000)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", consensusParams))
}

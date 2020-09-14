package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//{Version:1.0.0
//Update:0
//NewVersion:1.0.0}
func TestVersion(t *testing.T) {
	versionRes, err := akcSdk.GetVersion()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", versionRes))
}

//{NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:2db683eb2f02599fbd48765618a64a7b87e2adc5
//ListenAddr:tcp://0.0.0.0:26656
//Network:m0test
//Version:0.26.0
//Channels:4020212223303800
//Moniker:m0test
//Other:{TxIndex:on
//RPCAddress:tcp://0.0.0.0:26657}}
//SyncInfo:{LatestBlockHash:[23 80 67 219 141 120 212 13 60 216 64 132 8 61 189 247 175 118 212 16 182 247 176 120 227 158 54 232 1 53 235 159 120 244 16 121 3 143 61 19 109 250 0 93 58 7 158 3] LatestAppHash:[243 144 57 247 174 59 0 45 188 231 78 132 240 80 117 23 222 118 3 173 192 243 128 247 224 80 3 211 159 65 228 80 4 232 63 59 215 141 186 236 61 66 243 80 132 3 128 130] LatestBlockHeight:8101 LatestBlockTime:2020-07-24 08:49:46.988825306 +0000 UTC CatchingUp:false} ValidatorInfo:{Address:[240 78 116 11 79 123 235 141 66 215 222 55 23 207 54 228 81 60 219 158 193 15 79 58 235 192 192 220 63 53]
//PubKey:1624de64209d4d8272c0bc475fca81abdabf9dd43b74e598c094c26bb65d0379625bbc4bca VotingPower:10}}
func TestStatus(t *testing.T) {
	statusRes, err := akcSdk.GetStatus()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", statusRes))
}

//{Genesis:{GenesisTime:2020-01-02 15:04:05 +0000 UTC
//ChainID:m0test
//ConsensusParams:0xc000069d10
//Validators:[{Address:8E50C097641C1943F8825FE8257BD08668DA3D81 PubKey:1624de64209d4d8272c0bc475fca81abdabf9dd43b74e598c094c26bb65d0379625bbc4bca Power:10 Name:}
//{Address:13B3FCAB74376D4F22B8C9513AEB715E7E81FB42 PubKey:1624de64201853c11218bd8627cffa3bf5966a3724fd165067777539a524afe9b657cf2fe8 Power:10 Name:}
//{Address:D2AC87065C3B306EF5204C895EE16F7B11FD834C PubKey:1624de6420388a4d12a4da00745ddf7eaa659126798168ae2ec9a787ff1d088ce0cf1184cd Power:10 Name:}
//{Address:6AFF0D270B91B05881F28766C0047414BCE11C71 PubKey:1624de642062d292d932d0d7c76892812162c9798866357218ef6bac05c050c8de5c8efa85 Power:10 Name:}]
//AppHash: AppState:[] NativeAddress:m01qhcj9eze7rs3ka7rpzhp8hg4mc0ay3c2nct7qvr}}
func TestGenesis(t *testing.T) {
	genesisRes, err := akcSdk.GetGenesis()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", genesisRes))
}

//{RootDir:/data/cluster/.m0
//ApiAddress:0.0.0.0:30000
//DBBackend:leveldb
//DBPath:data
//KeysPath:keystore}
func TestConfig(t *testing.T) {
	configRes, err := akcSdk.GetConfig()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", configRes))
}

//{Peers:[{NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:a76c577ff25af1c4c70ae0bcae2224b596253991 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:false ConnectionStatus:{Duration:2562047h47m16.854775807s SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}} {NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:d7a40dc0d03f2983fa4702be89d47c7a116d2ee7 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:false ConnectionStatus:{Duration:2562047h47m16.854775807s SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}} {NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:369303c7b8560ff55dfe7a56433069ee7a6cae07 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:true ConnectionStatus:{Duration:2562047h47m16.854775807s SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}}]}
func TestListPeers(t *testing.T) {
	peersRes, err := akcSdk.ListPeers()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", peersRes))
}

//{Listening:true
//Syncing:true
//PeerCount:3
//CurrentBlock:8116
//HighestBlock:8116
//NetWorkID:solonet
//Version:0xc000288a50
//Listeners:[Listener(@)]
//Peers:[{NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:a76c577ff25af1c4c70ae0bcae2224b596253991 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:false
//ConnectionStatus:{Duration:2562047h47m16.854775807s
//SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}} {NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:d7a40dc0d03f2983fa4702be89d47c7a116d2ee7 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:false ConnectionStatus:{Duration:2562047h47m16.854775807s SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}} {NodeInfo:{ProtocolVersion:{P2P:4 Block:7 App:0} ID_:369303c7b8560ff55dfe7a56433069ee7a6cae07 ListenAddr:tcp://0.0.0.0:26656 Network:m0test Version:0.26.0 Channels:4020212223303800 Moniker:m0test Other:{TxIndex:on RPCAddress:tcp://0.0.0.0:26657}} IsOutbound:true ConnectionStatus:{Duration:2562047h47m16.854775807s SendMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} RecvMonitor:{Active:false Start:0001-01-01 00:00:00 +0000 UTC Duration:0s Idle:0s Bytes:0 Samples:0 InstRate:0 CurRate:0 AvgRate:0 PeakRate:0 BytesRem:0 TimeRem:0s Progress:0} Channels:[{ID:48 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:64 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:32 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:33 SendQueueCapacity:0 SendQueueSize:0 Priority:10 RecentlySent:0} {ID:34 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:35 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0} {ID:56 SendQueueCapacity:0 SendQueueSize:0 Priority:5 RecentlySent:0} {ID:0 SendQueueCapacity:0 SendQueueSize:0 Priority:1 RecentlySent:0}]}}]}
func TestNetInfo(t *testing.T) {
	netInfoRes, err := akcSdk.GetNetInfo()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", netInfoRes))
}

//{ProtocolVersion:{P2P:4 Block:7 App:0}
//ID_:2db683eb2f02599fbd48765618a64a7b87e2adc5
//ListenAddr:tcp://0.0.0.0:26656
//Network:m0test Version:0.26.0
//Channels:4020212223303800
//Moniker:m0test
//Other:{TxIndex:on
//RPCAddress:tcp://0.0.0.0:26657}}
func TestNodeInfo(t *testing.T) {
	nodeInfoRes, err := akcSdk.GetNodeInfo()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", nodeInfoRes))
}

func TestHealth(t *testing.T) {
	assert.Nil(t, akcSdk.CheckHealth())
}

//8125
func TestBlockHeight(t *testing.T) {
	height, err := akcSdk.GetBlockHeight()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", height))
}

//E0D5197A5CC2149107183FEA7DDF6E656F40962837CE81F73D6EC04A2D1E5CF5
func TestBlockHashByHeight(t *testing.T) {
	hash, err := akcSdk.GetBlockHashByHeight(1000)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", hash))
}

//{BlockHeight:1000
//Validators:
//[{Address:13B3FCAB74376D4F22B8C9513AEB715E7E81FB42
//PubKey:1624de64201853c11218bd8627cffa3bf5966a3724fd165067777539a524afe9b657cf2fe8
//VotingPower:10
//Accum:-30} {Address:6AFF0D270B91B05881F28766C0047414BCE11C71 PubKey:1624de642062d292d932d0d7c76892812162c9798866357218ef6bac05c050c8de5c8efa85 VotingPower:10 Accum:10} {Address:8E50C097641C1943F8825FE8257BD08668DA3D81 PubKey:1624de64209d4d8272c0bc475fca81abdabf9dd43b74e598c094c26bb65d0379625bbc4bca VotingPower:10 Accum:10} {Address:D2AC87065C3B306EF5204C895EE16F7B11FD834C PubKey:1624de6420388a4d12a4da00745ddf7eaa659126798168ae2ec9a787ff1d088ce0cf1184cd VotingPower:10 Accum:10}]
//ProposerAddress:D2AC87065C3B306EF5204C895EE16F7B11FD834C}
func TestValidatorsByHeight(t *testing.T) {
	validators, err := akcSdk.GetValidatorsByHeight(1000)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", validators))
}

//{Validator:true
//Address:8E50C097641C1943F8825FE8257BD08668DA3D81}
func TestNodeTypeByHeight(t *testing.T) {
	validator, err := akcSdk.GetNodeTypeByHeight(1000)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", validator))
}

//88B07A74214E6CC6E18223C82578C75C778619414688F24BE80268ADA423E9A5
func TestBlockHash(t *testing.T) {
	hash, err := akcSdk.GetBlockHash()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", hash))
}

//{CpuUsageRate:1.81%
//MemUsageRate:0.95%}
func TestHardwareUsage(t *testing.T) {
	hardware, err := akcSdk.GetHardwareUsage()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", hardware))
}

//{LastHeight:8135
//BlockMetas:[0xc0000f0000 0xc0000f0140 0xc0000f0280 0xc0000f03c0 0xc0000f0500 0xc0000f0640 0xc0000f0780 0xc0000f08c0 0xc0000f0a00 0xc0000f0b40 0xc0000f0c80 0xc0000f0dc0 0xc0000f0f00 0xc0000f1040 0xc0000f1180 0xc0000f12c0 0xc0000f1400 0xc0000f1540 0xc0000f1680 0xc0000f17c0]}
func TestBlockchainInfo(t *testing.T) {
	info, err := akcSdk.GetBlockHeader(1000, 1500)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", info))
}

//{Hash:E0D5197A5CC2149107183FEA7DDF6E656F40962837CE81F73D6EC04A2D1E5CF5
//Version:1
//ChainID:m0test
//Height:1000
//Time:2020-07-23T09:54:49.498992536Z
//PreviousBlockHash:3B89D96A4A570A2658D768CF6C8E6CCE8229756052469E2A3AE655B54A5C0717
//AppHash:488F0CAE866D6EEA63A3EC2D083B077AC55A5D7494E362C308E728DD4DFF8774
//NumTxs:0
//TotalTxs:2
//DataHash:
//LastResultsHash:
//ProposerAddress:D2AC87065C3B306EF5204C895EE16F7B11FD834C
//Size:1120}
func TestBlockDetail(t *testing.T) {
	details, err := akcSdk.GetBlockDetail(1000, "")
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", details.BlockHeaderResp))
	fmt.Println(fmt.Sprintf("%+v", details.Transactions))
}

//[{Alias:dgkcyxgbjqyffnmrlvpw
//OutputID:e871e513aced23709ea279acab991677f7a585256486e62250e2db64dd879b94
//AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b
//AssetAlias:NLWNPZXSJF
//Amount:10000000
//AccountID:0xC990EFBd72A0c3B16a1b5Cc8d41953677D96fDE2
//Address:m01q3f2d2u63stj3fgccru9xwdeg3wjup5pgr3p3j0
//ControlProgramIndex:1
//Program:00148a54d5735182e514a3181f0a6737288ba5c0d028
//SourceID:d615323ae2303d513c0596259e78f0fde77174e5e906003353e837e5351bd428
//SourcePos:1
//ValidHeight:0
//Change:false
//DeriveRule:0} {Alias:oybiznxpssdreidxmbnd OutputID:d3c26f90f75056343d94e591e5775207b7a0aed5a0c44e3a848b9098f6aa3573 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:9000000 AccountID:0x3E8cA543276Dd1b463856bB7199584cbBefBb8a0 Address:m01ql44d84q5fle9ve8qe7g9e8e9s82gazxwqguzea ControlProgramIndex:1 Program:0014fd6ad3d4144ff25664e0cf905c9f2581d48e88ce SourceID:079d24b6906417569749fcb428ad325199a3965f53622fb9078cfadff104bab8 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:xucewewqcfuyyfvxzbur OutputID:ba446968a15d308cf4d97cba7e0f68d19efc370da44c8ba0790db6a616f381df AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x3F53163C9D90deEB2785d1fC49788B336814Ad36 Address:m01qtd90gs7ctxtxgxmlfaljgket9r4jm9tzqm0j5a ControlProgramIndex:1 Program:00145b4af443d85996641b7f4f7f245b2b28eb2d9562 SourceID:e259c6afe764b135b7c5565cb89c966757e6345b4c21b757ce27211fa6424bd2 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:nhctcnutajiuwslrripg OutputID:b9285d26a457cd77a43a2e1b4ab8b605fff76c1644102a9e385b2b2433a0cfd8 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x374ba9f08bb4b1ed45E115356c077Ff082a6B1Dd Address:m01qh452n2urcrvwkruuhkmwn8xu942sckgya3wpfj ControlProgramIndex:1 Program:0014bd68a9ab83c0d8eb0f9cbdb6e99cdc2d550c5904 SourceID:29e00820616bbfc2d3236406a12023408f20574cffdd954f160e34f34954fc44 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:a02cb70317a9962bcabfbd4b40fc1f7c79e26c554043d27e5e7ab58662db6499 AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:fdc1eb942e260e9e1f58374bc8a7d1d16576c13874c2051f376f3b7def8ad625 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:oxfuetazfjqtefvlaiar OutputID:a02051d5b218086de6c917d488ac8e5f7548f38f5e8f4e74fa97a60adefd85fb AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x1E3B91Ec5904746f07A8045725672B65b1a0a3D8 Address:m01q2jw5sp4mmx4cyqqdm5cvurkhs4pnmycvqj3379 ControlProgramIndex:1 Program:0014549d4806bbd9ab82000ddd30ce0ed785433d930c SourceID:bae911b7800000631cde9d9b0d6aa0236db9522a50bf36959f33966886dcd72d SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:wfcjhrgftpfkdpqebbpi OutputID:9e7a9d095d0fc502c451e7d865c25c45e7531bece9a76950610f0f4400406cf6 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0xAcDE80a7B54111538014F2125c16F20E47bF6dd7 Address:m01q9533y0mjqg0w8pm3v86c54cx0cufjldk5c8uan ControlProgramIndex:1 Program:00142d23123f72021ee3877161f58a57067e38997db6 SourceID:424b32717af5a69e5d380c84348fd091bc62377af82c352e79604606a85306b0 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:hwjhnrkqjiicjkwuxwej OutputID:93a17333281cb550285de8317a0c200a5c2eee801c3b44ddb2b5d57205b676e7 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0xd7177fBC0e701Bd9b5D8221040F1BCc950E77E28 Address:m01q27elwy72a546sd540rr78y9f8rhea4j6kag7rj ControlProgramIndex:1 Program:001457b3f713caed2ba8369578c7e390a938ef9ed65a SourceID:f09ce6594f122098df2d1e1d0b4ab43f36f96440833753c48cdb893b6d0201f3 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:zqjkdjvvihyypmpgkait OutputID:8e889549de083a996d37548eeb383e5425e28a0c058f60b3905bbfdb74b17fdb AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:1000000 AccountID:0x49624Cb054033Ec4EF7052D1Ee5B03AacF4ab735 Address:m01qmr5e232s6emzw7yck8hek2ynhus27kuzvyykax ControlProgramIndex:1 Program:0014d8e9954550d676277898b1ef9b2893bf20af5b82 SourceID:acdc093f0681f9b225f053899754ccaebadd5afad65129b08b6d95e3bd49671f SourcePos:2 ValidHeight:0 Change:false DeriveRule:0} {Alias:najpsidivfluygzfoqyt OutputID:7406414bd0f1ea78e7479a0b7244eb9d5befb5ef6c7691b0fb94d11336ba63d1 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:8000000 AccountID:0x13AC067C41C595E4600D382Eb0CEd8685BC88631 Address:m01q838t0k32fqlz7ysq3m34tu0nypk77wq8dwaepm ControlProgramIndex:1 Program:00143c4eb7da2a483e2f12008ee355f1f3206def3807 SourceID:c72967a9dc2788808a3744128e4321ec408ef1c0779a6d32651b39341090b66c SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:6b9cadf42b550332fbcc3fb3464936796f850cd4b3597501f021ffaaf408c264 AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:dee7dce5d419095ffb7557d5b73c3a58a32c5fc95b71e145a1b914b6831e1c7b SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:aytntetigumarondjjhc OutputID:6a5f2b4eacd687ef32a2dc941da4babe2bac57917db3749e5d54b7cab413f034 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:1000000 AccountID:0x6E9612A6219860cD837fF1F2DEeC9eA3B8038C1c Address:m01qvjreexswf5qeunfkk666dglm73p2f3lh678d8r ControlProgramIndex:1 Program:001464879c9a0e4d019e4d36b6b5a6a3fbf442a4c7f7 SourceID:079d24b6906417569749fcb428ad325199a3965f53622fb9078cfadff104bab8 SourcePos:2 ValidHeight:0 Change:false DeriveRule:0} {Alias:kixusslguyxvkmczbimq OutputID:59980e8a8e16fbb3f083a4830c60ca181fbca9ad2cba6d26cdd75939dfa3feac AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:9000000 AccountID:0xA3eaefB21f022a6E528f110a68bB15B3d85742d4 Address:m01qgw9gckexf4zzx2g3hzl03fk0t6dy7hysh2g7de ControlProgramIndex:1 Program:0014438a8c5b264d44232911b8bef8a6cf5e9a4f5c90 SourceID:acdc093f0681f9b225f053899754ccaebadd5afad65129b08b6d95e3bd49671f SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:kdheedybjzlgzbytbkcd OutputID:46c1612b38cde8df684b773284bfbaa40608eee7ce3e7e2dc21bf3bef3e3af76 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0xBF997A463908eb35e0654c98B25FFE23c9F17f58 Address:m01qfuy7ddgqhvyz5d5s8jy3mjctrfs3gu5wn0skpq ControlProgramIndex:1 Program:00144f09e6b500bb082a36903c891dcb0b1a6114728e SourceID:5500cd69419805ec005a9c140b07b59f958b13f9032809b01f1a5c4b6d9e56d9 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:zwatrltyqbapcnlqgzxh OutputID:3cf333b501b34eeb4bb645a746165872a65f74a48790f9226da63e1f85f52434 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x6CdCC9dbC544B0C645Ee8Ec76da3b55B22E18275 Address:m01q9k8c9hq7fhw98ed3s0qsvrmwh563u27kf3327m ControlProgramIndex:1 Program:00142d8f82dc1e4ddc53e5b183c1060f6ebd351e2bd6 SourceID:98b27bbf959096e04ffa83503d1b64cc43cca6bd37f2afac2a8e093336822644 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:397fbfeef1ec084639b3416d84aabb55a57fef8e76a27a6d05698ecabeffa82e AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:54b5e8c08345c684b3058c6db43de3e438a0f639a5e4179a1714c4a97abba332 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:31d9eb3e7372129213a75481dc289588544e1207444508ed734c39139d71bb02 AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:a4d0f33db08d9a86d853b17df3e59dd3abedc1f47e95b5b9f2d1df0c4cbe947b SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:3051f97e086b6339d1952f7f7d0de649f88b7f9cd7dc4b4e673e472272605655 AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:4355879c2317c268d6e8fb0a3a236a648b842bb27e7fab4fcd21e8097a793f6b SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ahtqkkkrluqddhbxrwyo OutputID:2ea3cf957af54f48fc11a765daa85fedc77e097ff49f04d68f296a706fdc7bdd AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:8000000 AccountID:0xF035b6fFF12506CFf2227bd479a9b4CDF7f13cb3 Address:m01q4n55jaz4pyaq27jr32mnd29jmekxzu07m5he55 ControlProgramIndex:1 Program:0014ace9497455093a057a438ab736a8b2de6c6171fe SourceID:1ddc50eb352dd28b192b08dc93ce3a475e9d7b7adb49d351f8196af8ac1ea1a9 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:qmxjypyigdextkuvbkxl OutputID:28682d4fbe404027eaf23b01f4c02c75bc6729c0e4ada60d9fad14ede58a62de AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x494503833d9d82bc319786B98cfA3C1BDC609b63 Address:m01qdzkp87gwdpsdn3mmuqyyjcxhqxc6fqyy5d7d3h ControlProgramIndex:1 Program:001468ac13f90e6860d9c77be0084960d701b1a48084 SourceID:fa98b9c91956fa9cce24679af8922b497cf9be3ee58ce0dbe1824a9aadd78323 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:sgjvgzmmualyrqpnwhcz OutputID:27a716e90352da9a797c1f369b7fd5bc8b908fcf113acda498a85b4038ab8725 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0xaBebF8dd77fda368dF69daBB2c92FD893a3A3c3C Address:m01qcfqtyvegva8apt33jqhchltvyrrkwtzpl7r5aa ControlProgramIndex:1 Program:0014c240b23328674fd0ae31902f8bfd6c20c7672c41 SourceID:8b075b3f5fb30c0881b682a2c832b4acbef676670172752319296dd99a3de1b9 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ufjhcqpfwljkhnaaglum OutputID:185f49077d8c1f3b4afda8bcdbab5d328b0cde0b353fb36cc9f71361a6414427 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:8000000 AccountID:0x80E9Cc5C417E46AFfA7e2f2aada2A21CA57bC3E3 Address:m01qdxy0khh0ckeu62lmz7zyhfhjehrwkhhcakd8ch ControlProgramIndex:1 Program:00146988fb5eefc5b3cd2bfb17844ba6f2cdc6eb5ef8 SourceID:895feebdab89e5c0aca050031acc3afa3d7b56a8f9cb30951af238b07e87bb0e SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:ekisbxeuzfazsjttvoek OutputID:06429003fe0f817203f0f9e056a9a39fef0a9887751936fdcae11e985ac6890a AssetID:07bd5dbe9393c51a438a6417da00eba686a8a419b42bb081f3b86f85eaf1aab8 AssetAlias:CQVKYRXKPL Amount:10000000 AccountID:0xC3FAc690969B04FE241dFef3a5dEA241214C46e4 Address:m01qw0ylvs4xc8thjt2kahtx08mc8ttff5djlgpl0e ControlProgramIndex:1 Program:001473c9f642a6c1d7792d56edd6679f783ad694d1b2 SourceID:8535488c97222b3b552ec3e4e5ed13155a3abc0cf9997054ccf79b060c774500 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:gredrzhoblmtcinkqdnp OutputID:0427a656f895ab192fa39a94cae80304608defbb3800ec0309106c3d32083f26 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x3cF3EbaED64384f7ee3F4A3E07E589eACaB87Db6 Address:m01qccaqtxfwy3gnf2ykxqq7z3hpl2yh4y9psy64gq ControlProgramIndex:1 Program:0014c63a05992e245134a8963001e146e1fa897a90a1 SourceID:5c5629a75015d196772a66db0624f9054fae8210d2b0eb06c3529c6ed67de064 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0} {Alias:gredrzhoblmtcinkqdnp OutputID:03f3983ec8ca180c8484199c70a53ae1b2a47be527b8bae23d0cd463eda01ad2 AssetID:e58fd15d45782fd12ffb3a2ef8e0e4a71f42843fcce91bf07d5ec4bf31cfa22b AssetAlias:NLWNPZXSJF Amount:10000000 AccountID:0x3cF3EbaED64384f7ee3F4A3E07E589eACaB87Db6 Address:m01qccaqtxfwy3gnf2ykxqq7z3hpl2yh4y9psy64gq ControlProgramIndex:1 Program:0014c63a05992e245134a8963001e146e1fa897a90a1 SourceID:bed25dc771b60fb97d124093211c3aa4b8d6558f7232a3aa465b1487a2cdd3d6 SourcePos:1 ValidHeight:0 Change:false DeriveRule:0}]
func TestListUTXOs(t *testing.T) {
	utxo, err := akcSdk.ListUTXOs()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", utxo))
}

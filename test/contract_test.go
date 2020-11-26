package test

import (
	"encoding/hex"
	"fmt"
	account2 "github.com/akchain/go-sdk/account"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUploadContract(t *testing.T) {
	filePath := "C:\\Users\\zhaojing01\\Desktop\\test2.wasm"
	file, err := akcSdk.UploadContract(filePath)
	assert.Nil(t, err, err)
	fmt.Println(file)
}

func TestDeployContract(t *testing.T) {
	account, err := akcSdk.NewAccount()
	assert.Nil(t, err, err)
	//fmt.Println(account.ToString())
	fmt.Println(account.Xprv.String())
	//67ff81030101074163636f756e7401ff8200010601045870727601ff840001045870756201ff860001094163636f756e744944010c00010c4163636f756e74416c696173010c00010741646472657373010c00010e436f6e74726f6c50726f6772616d010c00000015ff83010101045850727601ff8400010601ff80000015ff85010101045850756201ff8600010601ff800000fe0163ff820140fff034725535ff92ffddff83ffe9ffe41602ffebffe9ff89ff962340ff8674ffb4011b20ffaaff925affdefff3ffdd1c4a4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973014053ff9dfffcffec2408ffe73fffe10149ffac66ffc0ff9319115affdbffccffae3affa1ff9661ff973a51ff9effbb5dffcc4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973012a3078446136664165373630336235666163353362653741453032613833643232446433333436363144460114574a49564f594756414c444b5349564b4a4c4f46012a6d303171346a72776b63306b797270353034666d72366a6461746b397636786d7677753830716838777a012c303031346163383665623631663632306333343764353362316561346465616563353636386462363362383700
	contractPath := "/home/smpl/go/src/github.com/akchain/akfs/contract/evidence.wasm"
	contractName := "evidence"
	args := "{\"creator\": \"someone\"}"
	txHash, err := akcSdk.DeployContract(contractPath, contractName, args, account)
	assert.Nil(t, err, err)
	fmt.Println(txHash)
}

func TestInvokeContract(t *testing.T) {
	//account, err := account2.LoadAcount("a04bc1eba844afadd62b7f4191b85faab81f0c7375884cc9ff9f8a7e2e28864311dece302fc4bcb76a28b7dcbdd066ea35e6aa4623a564a235e0cc3a305bef32")
	//assert.Nil(t, err, err)
	x, _ := hex.DecodeString("a04bc1eba844afadd62b7f4191b85faab81f0c7375884cc9ff9f8a7e2e28864311dece302fc4bcb76a28b7dcbdd066ea35e6aa4623a564a235e0cc3a305bef32")
	var xprv chainkd.XPrv
	copy(xprv[:], x[:64])
	account, err := akcSdk.NewAccountFromPrivateKey(xprv)
	assert.Nil(t, err, err)
	contractName := "evidence"
	method := "storeFileInfo"
	args := "{\"user_id\":\"test_user_id_2\",\"uuid\":\"test_uuid_2\",\"evidence_data\":\"test_evidence_data\"}"
	hash, err := akcSdk.InvokeMethod(contractName, method, args, account)
	assert.Nil(t, err, err)
	fmt.Println(hash)
}

func TestUpgradeContract(t *testing.T) {
	account, err := account2.LoadAcount("67ff81030101074163636f756e7401ff8200010601045870727601ff840001045870756201ff860001094163636f756e744944010c00010c4163636f756e74416c696173010c00010741646472657373010c00010e436f6e74726f6c50726f6772616d010c00000015ff83010101045850727601ff8400010601ff80000015ff85010101045850756201ff8600010601ff800000fe0163ff820140fff034725535ff92ffddff83ffe9ffe41602ffebffe9ff89ff962340ff8674ffb4011b20ffaaff925affdefff3ffdd1c4a4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973014053ff9dfffcffec2408ffe73fffe10149ffac66ffc0ff9319115affdbffccffae3affa1ff9661ff973a51ff9effbb5dffcc4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973012a3078446136664165373630336235666163353362653741453032613833643232446433333436363144460114574a49564f594756414c444b5349564b4a4c4f46012a6d303171346a72776b63306b797270353034666d72366a6461746b397636786d7677753830716838777a012c303031346163383665623631663632306333343764353362316561346465616563353636386462363362383700")
	assert.Nil(t, err, err)
	contractName := "test30"
	newPath := "/data/cluster/wasm/contract/counter.wasm"
	hash, err := akcSdk.UpgradeContract(newPath, contractName, account)
	assert.Nil(t, err, err)
	fmt.Println(hash)
}

func TestQueryContract(t *testing.T) {
	contractName := "evidence"
	args := "{\"user_id\":\"1234567\"}"
	method := "queryFileInfoByUser"
	queryResult, err := akcSdk.QueryContract(contractName, method, args)
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", queryResult[0]))
}

func TestListContracts(t *testing.T) {
	account, err := account2.LoadAcount("67ff81030101074163636f756e7401ff8200010601045870727601ff840001045870756201ff860001094163636f756e744944010c00010c4163636f756e74416c696173010c00010741646472657373010c00010e436f6e74726f6c50726f6772616d010c00000015ff83010101045850727601ff8400010601ff80000015ff85010101045850756201ff8600010601ff800000fe0163ff820140fff034725535ff92ffddff83ffe9ffe41602ffebffe9ff89ff962340ff8674ffb4011b20ffaaff925affdefff3ffdd1c4a4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973014053ff9dfffcffec2408ffe73fffe10149ffac66ffc0ff9319115affdbffccffae3affa1ff9661ff973a51ff9effbb5dffcc4377ffb8721dff8328607dfff7ffc225ff86ffc9112fffdbff93ffdcffc7fffbffcb186e674affb4ffb4ff9e603973012a3078446136664165373630336235666163353362653741453032613833643232446433333436363144460114574a49564f594756414c444b5349564b4a4c4f46012a6d303171346a72776b63306b797270353034666d72366a6461746b397636786d7677753830716838777a012c303031346163383665623631663632306333343764353362316561346465616563353636386462363362383700")
	assert.Nil(t, err, err)
	contracts, err := akcSdk.ListContracts(account.GetAccountID())
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", contracts))
}

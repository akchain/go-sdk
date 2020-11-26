package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNumberUnconfirmedTx(t *testing.T) {
	number, err := akcSdk.GetNumUnconfirmedTx()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", number))
}

//&{Total:0 TXIDs:[]}
func TestListUnconfirmedTx(t *testing.T) {
	txs, err := akcSdk.ListUnconfirmedTx()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", txs))
}

//{Time:2020-07-24T03:45:25.523482807Z
//Height:0xc00008b350
//TxHash:21C2B2A0AB9266EF8D5A54FB33953F450F8E27AA979C02077173FCB33A80E2E7
//Version:1
//Size:328
//Status:success
//Type:bank
//TransferTxInfo:0xc0000861e0
//WasmTxInfo:<nil>
//ValidatorUpdate:<nil>}
//
//{ID:a87d16e5a78a158508cd48d7e99986eb8ff4c53386ec802a7e3de3d74436e99b
//Inputs:[0xc0000c2d00]
//Outputs:[0xc0000b6210 0xc0000b62c0]
//MuxID:5c5629a75015d196772a66db0624f9054fae8210d2b0eb06c3529c6ed67de064
//MetaData:}
func TestSearchTx(t *testing.T) {
	tx, err := akcSdk.SearchTx("C43E32BA7E93807F20D180F9C6E63AC9AA19F6C0425C883364E1E736BA174BDD")
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", tx))
	fmt.Println(fmt.Sprintf("%+v", tx.TransferTxInfo))
}

package test

import . "github.com/akchain/go-sdk"

var (
	akcSdk *AkchainSdk
)

func init() {
	akcSdk = NewAkchainSdk()
	akcSdk.NewRestClient().SetAddress("http://106.14.210.116:30000")
}

package test

import . "github.com/akchain/go-sdk"

var (
	akcSdk *AkchainSdk
)

func init() {
	akcSdk = NewAkchainSdk()
	akcSdk.NewRestClient().SetAddress("http://localhost:30000")
}

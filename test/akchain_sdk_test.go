package test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//spot rebel three size duty patient hood flame crowd find oblige armor woman salon large eight figure prosper flock patrol gas amount jungle tourist
func TestGenerateMnemonicCodesStr(t *testing.T) {
	mnemonic, err := akcSdk.GenerateMnemonicCodesStr()
	assert.Nil(t, err, err)
	fmt.Println(fmt.Sprintf("%+v", mnemonic))
}

func TestGetKeysFromMnemonic(t *testing.T) {
	mnemonic := "barrel put drive fine series income buzz doctor today soup summer put"
	xprv, xpub, err := akcSdk.GetKeysFromMnemonic(mnemonic)
	assert.Nil(t, err, err)
	assert.Equal(t, "382b69c48e87880b8f7f1c4146a6fbac383a80d5fff701475f64934c8297be5c6ff8fb5c43be3c0c"+
		"86f1fdbd6db8d9c0e943c43091e0b257e9a7c57bfccf2575", xprv)
	assert.Equal(t,
		"991b5023ff07a5634e66ee3c1e29b5414977b9c224c0938fd63c777df61f44df6ff8fb5c4"+
			"3be3c0c86f1fdbd6db8d9c0e943c43091e0b257e9a7c57bfccf2575", xpub)
}

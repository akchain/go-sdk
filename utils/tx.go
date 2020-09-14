package utils

import (
	"akc/go-sdk/types"
	"encoding/hex"
	"github.com/bytom/bytom/crypto/ed25519/chainkd"
	"github.com/bytom/bytom/crypto/sha3pool"
	"github.com/bytom/bytom/errors"
	"io"
)

func SignContract(privateKey string, tx *types.WasmTx) error {
	xprvBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return errors.WithDetailf(err, "bad key string")
	}
	var xprv chainkd.XPrv
	copy(xprv[:], xprvBytes[:])
	wc := tx.TxSigWitness
	switch wc.Type {
	case "raw_tx_signature":
		dst, err := FindDst(privateKey, tx.TxSigWitness.Keys)
		if err != nil {
			return err
		}
		if len(tx.TxSigWitness.Keys[dst].DerivationPath) > 0 {
			pathBytes := [][]byte{}
			path := tx.TxSigWitness.Keys[dst].DerivationPath
			for _, p := range path {
				pathBytes = append(pathBytes, p)
			}
			xprv = xprv.Derive(pathBytes)
		}
		signBytes := xprv.Sign(tx.DigestHash)
		wc.Sigs[dst] = signBytes
	}
	return nil
}

func FindDst(privateKey string, keyids []types.KeyID) (int, error) {
	dst := -1
	for k := 0; k < len(keyids); k++ {
		xprvBytes, err := hex.DecodeString(privateKey)
		if err != nil {
			return dst, errors.WithDetailf(err, "bad key string")
		}
		xprv := new(chainkd.XPrv)
		copy(xprv[:], xprvBytes[:])
		if string(keyids[k].Xpub[:]) == string(xprv.XPub().Bytes()) {
			dst = k
			break
		}
	}
	if dst == -1 {
		return dst, errors.New("Not a proper pub key to sign transaction.")
	}

	return dst, nil
}

func Sign(privateKey string, tpl *types.Template, decodeTx *types.RawTx) error {
	xprvBytes, err := hex.DecodeString(privateKey)
	if err != nil {
		return errors.WithDetailf(err, "bad key string")
	}
	xprv := new(chainkd.XPrv)
	copy(xprv[:], xprvBytes[:])
	for i, sigInst := range tpl.SigningInstructions {
		for j, wc := range sigInst.WitnessComponents {
			switch wc.Type {
			case "raw_tx_signature":
				input := decodeTx.Inputs[sigInst.Position].InputID
				decodeInput, err := hex.DecodeString(input)
				if err != nil {
					return errors.WithDetailf(err, "adding signature(s) to raw-signature witness component %d of input %d", j, i)
				}
				txId := decodeTx.TransferId

				decodeTxId, err := hex.DecodeString(txId)
				if err != nil {
					return errors.WithDetailf(err, "adding signature(s) to raw-signature witness component %d of input %d", j, i)
				}
				message := SigHash(decodeInput, decodeTxId)
				err = wc.Sign(*xprv, message)
				if err != nil {
					return errors.WithDetailf(err, "adding signature(s) to raw-signature witness component %d of input %d", j, i)
				}
			}
		}
	}
	return nil
}

func SigHash(input, txId []byte) []byte {
	hasher := sha3pool.Get256()
	defer sha3pool.Put256(hasher)
	hasher.Write(input)
	hasher.Write(txId)
	var b32 [32]byte
	_, err := io.ReadFull(hasher, b32[:])
	if err != nil {
		return nil
	}
	return b32[:]
}

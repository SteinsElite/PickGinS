package util

import (
	"encoding/json"
	"testing"

	"github.com/ethereum/go-ethereum/crypto"
)

func TestVerifySignature(t *testing.T) {
	privateKey, err := crypto.HexToECDSA("ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80")
	if err != nil {
		t.Fatal(err)
	}
	data := []byte("PICK")

	dataHash := crypto.Keccak256Hash(data)

	t.Log(dataHash.Bytes())
	k,_:= json.Marshal(dataHash.Bytes())
	t.Log(k)
	t.Log(string(k))
	sig, err := crypto.Sign(dataHash.Bytes(), privateKey)
	if err != nil {
		t.Fatal(err)
	}

	res := VerifySignature("0xf39fd6e51aad88f6f4ce6ab8827279cfffb92266", dataHash.Bytes(), sig)
	if !res {
		t.Fatal("fail to verify")
	}
}

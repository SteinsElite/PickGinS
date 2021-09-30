package util

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
)

// PriKeyToPub convert the private key to the public key
func PriKeyToPub(hexKey string) []byte{
	privateKey, err := crypto.HexToECDSA(hexKey)
	if err != nil{
		log.Fatal(err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	return publicKeyBytes
}

// VerifySig Verify the signature of dataHash sign by the publicKeyBytes
func VerifySig(sig []byte, dataHash common.Hash, publicKeyBytes []byte ) bool {
	sigNoRecoverID := sig[:len(sig)-1] // remove recovery id
	return crypto.VerifySignature(publicKeyBytes, dataHash.Bytes(), sigNoRecoverID)
}
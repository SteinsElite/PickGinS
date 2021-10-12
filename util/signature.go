package util

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func VerifySignature(addr string, data []byte, sig []byte) bool {
	sigAddr := SigToAddress(data, sig)
	addrByte := common.HexToAddress(addr)
	if sigAddr == addrByte {
		return true
	}
	return false
}

func SigToAddress(data []byte, sig []byte) common.Address {
	sigPublicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		log.Fatal(err)
	}
	sigAddr := crypto.PubkeyToAddress(*sigPublicKey)
	return sigAddr
}

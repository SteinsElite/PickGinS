package auth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

const (
	admin = "0x397a9e1719113Cd68ba79d59e1e1988C669cA7F3"
)

func IsPublisher(data string, sig string) bool {
	if VerifySignature(admin, []byte(data), []byte(sig)) {
		return true
	}
	return false
}

func VerifySignature(addr string, data []byte, sig []byte) bool {
	sigPublicKey, err := crypto.SigToPub(data, sig)
	if err != nil {
		panic(err)
	}
	sigAddress := crypto.PubkeyToAddress(*sigPublicKey)
	if sigAddress == common.HexToAddress(addr) {
		panic(err)
	}
	return false
}

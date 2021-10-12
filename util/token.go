package util

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// token.go is some constant and tool to deal with token on heco

// address of the Heco-peg assets(should never change)
var (
	BTCAddr  = common.HexToAddress("0x66a79d23e58475d2738179ca52cd0b41d73f0bea")
	ETHAddr  = common.HexToAddress("0x64ff637fb478863b7468bc97d30a5bf3a428a1fd")
	USDTAddr = common.HexToAddress("0xa71edc38d189767582c38a3145b5873052c3e47a")
	HTAddr   = common.HexToAddress("0x5545153CCFcA01fbd7Dd11C0b23ba694D9509A6F")
	MDXAddr  = common.HexToAddress("0x25d2e80cb6b86881fd7e07dd263fb79f4abe033c")
)

// symbol of token
const (

	// coin symbol

	BTC  = "BTC"
	ETH  = "ETH"
	USDT = "USDT"
	HT   = "HT"
	MDX  = "MDX"

	// coin Ids to query from coinGecko api

	BTCIds  = "bitcoin"
	ETHIds  = "ethereum"
	USDTIds = "tether"
	HTIds   = "huobi-token"
	MDXIds  = "mdex"

	VsCurrency = "usd"
)

// IdentifyToken convert token address ==> token symbol, return error when the token is not
// one of the supported address
func IdentifyToken(token common.Address) (string, error) {
	var tokenType string
	switch token {
	case MDXAddr:
		tokenType = MDX
	case USDTAddr:
		tokenType = USDT
	case HTAddr:
		tokenType = HT
	case ETHAddr:
		tokenType = ETH
	case BTCAddr:
		tokenType = BTC
	default:
		return tokenType, fmt.Errorf("unknown token address")
	}
	return tokenType, nil
}

// TokenIds get the token ids of the specific token, return error when token is not supported
func TokenIds(coinSymbol string) (ids string, err error) {
	switch coinSymbol {
	case BTC:
		ids = BTCIds
	case ETH:
		ids = ETHIds
	case USDT:
		ids = USDTIds
	case HT:
		ids = HTIds
	case MDX:
		ids = MDXIds
	default:
		err = fmt.Errorf("unknown token symbol")
	}
	return
}

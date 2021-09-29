package token

import (
	"errors"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
)

// address of the Heco-peg assets
var (
	MDXAddr  = common.HexToAddress("0x25d2e80cb6b86881fd7e07dd263fb79f4abe033c")
	USDTAddr = common.HexToAddress("0xa71edc38d189767582c38a3145b5873052c3e47a")
	HTAddr   = common.HexToAddress("0x5545153CCFcA01fbd7Dd11C0b23ba694D9509A6F")
	ETHAddr  = common.HexToAddress("0x64ff637fb478863b7468bc97d30a5bf3a428a1fd")
	BTCAddr  = common.HexToAddress("0x66a79d23e58475d2738179ca52cd0b41d73f0bea")
)

// ids to query info from 3rd-party price-info api
const (
	BTCIds      = "bitcoin"
	ETHIds      = "ethereum"
	USDTIds     = "tether"
	HTIds       = "huobi-token"
	MDXIds      = "mdex"
	Vs_currency = "usd"
)

// token symbol of the assets
const (
	BTC  = "BTC"
	ETH  = "ETH"
	USDT = "USDT"
	HT   = "HT"
	MDX  = "MDX"
)

// Identify the token by address on the heco
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
		return tokenType, errors.New("unsupported token address")
	}
	return tokenType, nil
}

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
		err = fmt.Errorf("unrecognized Token symbol")
	}
	return
}

// func TokenAddr(coin string) (addr common.Address, err error) {
// 	switch coin {
// 	case BTC:
// 		addr = BTCAddr
// 	case ETH:
// 		addr = ETHAddr
// 	case USDT:
// 		addr = USDTAddr
// 	case HT:
// 		addr = HTAddr
// 	case MDX:
// 		addr = MDXAddr
// 	default:
// 		err = fmt.Errorf("unrecognized Token")
// 	}
// 	return
// }

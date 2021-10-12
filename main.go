package main

import (
	"github.com/SteinsElite/pickGinS/router"
	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
)

func main() {
	// start some service
	transaction.StartTxWatcher()
	coin.StartCoinInfoWatcher()
	vault.StartVaultWatcher()

	router.InitAuth()
	// start the api server
	r := router.SetupGinServer()
	r.Run(":8080")
}

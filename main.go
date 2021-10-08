package main

import (
	"github.com/SteinsElite/pickGinS/router"
	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
)

func main() {
	go transaction.PollTxInterval()
	go coin.RunCoinInfoWatcher()
	// start the vault watcher
	vault.StartVaultWatcher()

	r := router.SetupGinServer()
	r.Run(":8080")
}
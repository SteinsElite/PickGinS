package main

import (
	"github.com/SteinsElite/pickGinS/internal/coin"
	"github.com/SteinsElite/pickGinS/internal/transaction"
	"github.com/SteinsElite/pickGinS/internal/vault"
	"github.com/SteinsElite/pickGinS/router"
)

func main() {
	go transaction.PollTxInterval()
	go coin.RunCoinInfoWatcher()
	go vault.RunVaultWatcher()
	r := router.SetupGinServer()
	r.Run(":8090")
}

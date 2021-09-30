package main

import (
	"github.com/SteinsElite/pickGinS/internal/coin"
	"github.com/SteinsElite/pickGinS/internal/vault"
	"github.com/SteinsElite/pickGinS/router"
	"github.com/SteinsElite/pickGinS/service/transaction"
)

func main() {
	go transaction.PollTxInterval()
	go coin.RunCoinInfoWatcher()
	go vault.RunVaultWatcher()
	r := router.SetupGinServer()
	r.Run(":8080")
}
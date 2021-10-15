package main

import (
	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/logging"
	"github.com/SteinsElite/pickGinS/router"
	"github.com/SteinsElite/pickGinS/service/coin"
	"github.com/SteinsElite/pickGinS/service/transaction"
	"github.com/SteinsElite/pickGinS/service/vault"
)

const (
	dbname = "pick"
	logDir = "./logdata"
)

func main() {

	// prepare the database and the log
	logging.InitLogger(logDir)
	defer logging.SyncLogger()

	storage.InitDB(dbname)

	// start each service
	transaction.StartTxWatcher()
	coin.StartCoinInfoWatcher()
	vault.StartVaultCache()

	// start the api server
	r := router.SetupGinServer()
	r.Run(":8080")
}

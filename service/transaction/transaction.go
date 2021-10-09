package transaction

import (
	"context"
	"fmt"
	"log"
	"math/big"
	"sync"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/SteinsElite/pickGinS/internal/gateway"
	"github.com/SteinsElite/pickGinS/internal/storage"
	"github.com/SteinsElite/pickGinS/util"
)

const (
	// the interval that heco mine new block
	hecoBatch = 3
	interval  = hecoBatch * 3
	// indicate that the block that the pickRouter contract deploy transaction is included
	geniusBlock = 8121824
	hecoLimit   = 5000
	txColl      = "transaction"
)

var (
	depositSigHash     = crypto.Keccak256Hash([]byte("Deposit(address,address,uint256)"))
	withdrawSigHash    = crypto.Keccak256Hash([]byte("Withdraw(address,address,uint256)"))
	claimProfitSigHash = crypto.Keccak256Hash([]byte("ClaimProfit(address,address,uint256)"))
)

const (
	Deposit     string = "deposit"
	Withdraw    string = "withdraw"
	ClaimProfit string = "claimProfit"
)

type BlockSpan struct {
	FromBlock int64
	ToBlock   int64
}

type TxRecord struct {
	TxHash      string  `json:"tx_hash" bson:"tx_hash"`
	BlockNumber uint64  `json:"block_number" bson:"block_number"`
	Timestamp   uint64  `json:"timestamp" bson:"timestamp"`
	TxType      string  `json:"tx_type" bson:"tx_type"`
	User        string  `json:"user" bson:"user"`
	Token       string  `json:"token" bson:"token"`
	Amount      float64 `json:"amount" bson:"amount"`
}

type TxWatcher struct {
	gateway.RpcClient
	CurrentBlock int64 // indicate the last block that has been processed
}

func newTxWatcher(currentBlock int64) *TxWatcher {
	return &TxWatcher{
		RpcClient:    *gateway.GetRpcClient(),
		CurrentBlock: currentBlock,
	}
}

// InitTxWatcher the tx watcher is context related, so when we restart the application,
// we should recovery the currentBlock
func InitTxWatcher() *TxWatcher {
	coll := storage.AccessCollections(txColl)
	// get the tx record with the biggest block number in the database
	opt := options.Find()
	opt.SetSort(bson.D{{"block_number", -1}})
	opt.SetLimit(1)

	cur, err := coll.Find(
		context.TODO(),
		bson.D{},
		opt,
	)
	if err != nil {
		log.Println(err)
	}
	defer cur.Close(context.TODO())

	var record TxRecord
	if cur.Next(context.TODO()) {
		// if this is not the first time to start the tx watcher, set the current block, or we
		// will use the genesis block as default
		err := cur.Decode(&record)
		if err != nil {
			log.Println(err)
		}
		currentBlock := int64(record.BlockNumber) + 1
		return newTxWatcher(currentBlock)
	}
	return GenesisTxWatcher()
}

// GenesisTxWatcher When the first time to start the app, we should query start form the genesis
// contract created block number
func GenesisTxWatcher() *TxWatcher {
	return newTxWatcher(geniusBlock)
}

func (tw *TxWatcher) ObtainTxUntil(toblk int64) (txs []TxRecord) {
	var blockSpans []BlockSpan
	for startblk := tw.CurrentBlock; startblk <= toblk; startblk += hecoLimit {
		endBlk := startblk + hecoLimit - 1
		if endBlk > toblk {
			endBlk = toblk
		}
		blockSpans = append(blockSpans, BlockSpan{startblk, endBlk})
	}

	var wg sync.WaitGroup
	// lock just used for append the slice, so the performance loss doesn't matter
	var m sync.Mutex
	for i := 0; i < len(blockSpans); i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			records, _ := tw.obtainTxRange(blockSpans[i])
			m.Lock()
			txs = append(txs, records...)
			m.Unlock()
		}(i)
	}
	wg.Wait()

	tw.CurrentBlock = toblk + 1
	log.Println("finish get the transaction with amount: ", len(txs))
	return
}

// get all the required logs in range [fromblk,toblk],make sure that toblk-fromblk is less that
// 5000 due to heco node limit
func (tw *TxWatcher) obtainTxRange(span BlockSpan) ([]TxRecord, error) {
	var txs []TxRecord
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(span.FromBlock),
		ToBlock:   big.NewInt(span.ToBlock),
		Addresses: []common.Address{tw.ContractAddr},
	}
	elogs, err := tw.RpcClient.Client.FilterLogs(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	for _, elog := range elogs {
		tx, err := tw.populateTxFromLog(elog)
		if err != nil {
			log.Println("[populateTxFromLog]: ", err)
		} else {
			txs = append(txs, tx)
		}
	}
	return txs, nil
}

func (tw *TxWatcher) populateTxFromLog(vlog types.Log) (TxRecord, error) {
	var tx TxRecord
	switch vlog.Topics[0].Hex() {
	case depositSigHash.Hex():
		tx.TxType = Deposit
		deposit, err := tw.Instance.ParseDeposit(vlog)
		if err != nil {
			return tx, err
		}
		tx.User = deposit.User.String()
		tx.Token, _ = util.IdentifyToken(deposit.Token)
		tx.Amount = util.Amount2Float(deposit.Value)

	case withdrawSigHash.Hex():
		tx.TxType = Withdraw
		withdraw, err := tw.Instance.ParseWithdraw(vlog)
		if err != nil {
			return tx, err
		}
		tx.User = withdraw.User.String()
		tx.Token, _ = util.IdentifyToken(withdraw.Token)
		tx.Amount = util.Amount2Float(withdraw.Value)

	case claimProfitSigHash.Hex():
		tx.TxType = ClaimProfit
		claimProfit, err := tw.Instance.ParseClaimProfit(vlog)
		if err != nil {
			return tx, err
		}
		tx.User = claimProfit.User.String()
		tx.Token, _ = util.IdentifyToken(claimProfit.Token)
		tx.Amount = util.Amount2Float(claimProfit.Value)
	default:
		return tx, fmt.Errorf("unmatched event")
	}

	tx.BlockNumber = vlog.BlockNumber
	blockHeader, err := tw.RpcClient.Client.HeaderByNumber(
		context.TODO(),
		big.NewInt(int64(tx.BlockNumber)),
	)
	if err != nil {
		return tx, err
	}
	tx.Timestamp = blockHeader.Time
	tx.TxHash = vlog.TxHash.String()
	return tx, nil
}

// persist the transaction record in the database,
// TODO(ERIJ): handle the error when fail to insert the transaction
func persistRecord(txs []TxRecord) {
	coll := storage.AccessCollections(txColl)
	for _, tx := range txs {
		_, err := coll.InsertOne(context.TODO(), tx)
		if err != nil {
			log.Println("Fail to write to database due to: ", err)
		}
	}
}

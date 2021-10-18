package gateway

import (
	"context"
	"fmt"
	"math/big"

	"github.com/SteinsElite/pickGinS/logging"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"go.uber.org/zap"

	"github.com/SteinsElite/pickGinS/internal/gateway/pickrouter"
)

// This file is a wrapper to interact  with the contract on the heco

const (
	// Visit from China Mainland
	hecoUrl = "https://http-mainnet-node.huobichain.com"
	// Visit to the community node
	hecoCommunity = "https://http-mainnet-node.defibox.com"
	// international visit to the heco
	hecoInternational = "https://http-mainnet.hecochain.com"
)

var (
	contractAddr = common.HexToAddress("0x94eD19Fff97f1864372E60734A569e23941EeCB8")
	hecoEndpoint = []string{
		hecoUrl,
		hecoCommunity,
		hecoInternational,
	}
)

type RpcClient struct {
	Client       *ethclient.Client
	ContractAddr common.Address
	Instance     *pickrouter.Pickrouter
}

// get a new RpcClient connected with rpc endpoint
func newRpcClient(rpcUrl string, contractAddr common.Address) (*RpcClient, error) {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		return nil, fmt.Errorf("fail to dial ethclient-[%w]", err)
	}
	instance, err := pickrouter.NewPickrouter(contractAddr, client)
	if err != nil {
		return nil, fmt.Errorf("fail to bind pick router-[%w]", err)
	}

	return &RpcClient{
		Client:       client,
		ContractAddr: contractAddr,
		Instance:     instance,
	}, nil
}

func (rc *RpcClient) IsClientConnected() bool {
	if _, err := rc.Client.ChainID(context.TODO()); err != nil {
		// if we can't get the chainId now, it seems that something going wrong with the connection
		// to the heco node
		return false
	}
	return true
}

// SelectEndpoint change the rpc endpoint, if no endpoint is available, return false, else change to
// the available endpoint and return true
func (rc *RpcClient) SelectEndpoint() (bool, string) {
	for i := 0; i < 3; i++ {
		var err error
		rc.Client, err = ethclient.Dial(hecoEndpoint[i])
		if err != nil {
			return false, ""
		}
		if rc.IsClientConnected() {
			return true, hecoEndpoint[i]
		}
	}
	return false, ""
}

// some function include change the endpoint and retry

func (rc *RpcClient) TimestampByNumber(blockNumber int64) (uint64, error) {
	blockHeader, err := rc.Client.HeaderByNumber(context.TODO(), big.NewInt(blockNumber))
	if err != nil {
		err = rc.CheckInteractionWithContract(err)
		if err != nil {
			return 0, err
		}
		blockHeader, _ = rc.Client.HeaderByNumber(context.TODO(), big.NewInt(blockNumber))
	}
	return blockHeader.Time, nil
}

func (rc *RpcClient) FilterLogs(query ethereum.FilterQuery) ([]types.Log, error) {
	elog, err := rc.Client.FilterLogs(context.TODO(), query)
	err = rc.CheckInteractionWithContract(err)
	if err != nil {
		return nil, err
	}
	elog, err = rc.Client.FilterLogs(context.TODO(), query)
	if err != nil {
		return nil, err
	}
	return elog, nil
}

func (rc *RpcClient) CheckInteractionWithContract(err error) error {
	if err != nil {
		if rc.IsClientConnected() {
			return err
		}
		if ok, endpoint := rc.SelectEndpoint(); ok {
			// if we could use a available endpoint
			logging.Z().Info(
				"[ENDPOINT CHANGE]",
				zap.String("new_endpoint", endpoint),
			)
			return nil
		}
		return fmt.Errorf("can't find availbe connection to blockchain")
	}
	return nil
}

// GetRpcClient get a rpc client to interact with the pickrouter contract
// if fail to get the client then panic
func GetRpcClient() *RpcClient {
	// due to that the GetRpcClient often occur at begin of start the program, and shouldn't fail,so
	// if fail to new a client, just panic
	client, err := newRpcClient(hecoUrl, contractAddr)
	if err != nil {
		panic(err)
	}
	if !client.IsClientConnected() {
		if ok, _ := client.SelectEndpoint(); !ok {
			panic("can't connect to the heco node")
		}
	}
	return client
}

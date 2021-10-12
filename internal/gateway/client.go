package gateway

import (
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"

	"github.com/SteinsElite/pickGinS/internal/gateway/pickrouter"
)

// This file is a wrapper to interact  with the contract on the heco

const (
	hecoUrl = "https://http-mainnet-node.huobichain.com"
)

var (
	contractAddr = common.HexToAddress("0x94eD19Fff97f1864372E60734A569e23941EeCB8")
)

type RpcClient struct {
	Client       *ethclient.Client
	ContractAddr common.Address
	Instance     *pickrouter.Pickrouter
}

// get a new RpcClient connected with rpc endpoint
func newRpcClient(rpcUrl string, contractAddr common.Address) *RpcClient {
	client, err := ethclient.Dial(rpcUrl)
	if err != nil {
		log.Println(err)
	}
	instance, err := pickrouter.NewPickrouter(contractAddr, client)
	if err != nil {
		log.Println(err)
	}

	return &RpcClient{
		Client:       client,
		ContractAddr: contractAddr,
		Instance:     instance,
	}
}

// GetRpcClient get a rpc client to interact with the pickrouter contract
func GetRpcClient() *RpcClient {
	return newRpcClient(hecoUrl, contractAddr)
}

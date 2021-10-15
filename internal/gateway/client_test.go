package gateway

import "testing"

func TestSelectEndpoint(t *testing.T) {
	rc, err := newRpcClient("https://http-mainnet-node.huobichain.com", contractAddr)
	if err != nil {
		t.Fatal(err)
	}
	if !rc.IsClientConnected() {
		t.Fatal("fail")
	}

	if ok, endpoint := rc.SelectEndpoint(); ok {
		t.Log("select the endpoint: ", endpoint)
	}
}

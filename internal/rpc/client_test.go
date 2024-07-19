package rpc

import (
	"net/http"
	"net/http/httptest"
	"simple-blockchain-client/pkg/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSendRPCRequest(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"jsonrpc": "2.0", "result": "0x1", "id": 1}`))
	}))
	defer server.Close()

	rpcClient := NewRPCClient(server.URL)

	req := models.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      1,
	}

	resp, err := rpcClient.SendRPCRequest(req)
	assert.NoError(t, err)
	assert.Equal(t, "0x1", resp.Result)
}

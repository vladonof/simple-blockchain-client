package handlers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"simple-blockchain-client/internal/rpc"
	"simple-blockchain-client/pkg/models"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestGetBlockNumber(t *testing.T) {
	mockClient := new(rpc.MockRPCClient)
	mockClient.On("SendRPCRequest", mock.Anything).Return(models.RPCResponse{
		Result: "0x1",
		ID:     2,
	}, nil)

	h := NewHandlers(mockClient)

	req := httptest.NewRequest("POST", "/blockNumber", nil)
	w := httptest.NewRecorder()

	h.GetBlockNumber(w, req)

	expectedRPCRequest := models.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      2,
	}
	mockClient.AssertCalled(t, "SendRPCRequest", expectedRPCRequest)

	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var responseBody map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "0x1", responseBody["result"])

	mockClient.AssertExpectations(t)
}

func TestGetBlockByNumber(t *testing.T) {
	mockClient := new(rpc.MockRPCClient)
	mockClient.On("SendRPCRequest", mock.Anything).Return(models.RPCResponse{
		Result: "0x1",
		ID:     2,
	}, nil)

	h := NewHandlers(mockClient)

	body := `{"blockNumber": "0x1"}`
	req := httptest.NewRequest("POST", "/blockByNumber", strings.NewReader(body))
	w := httptest.NewRecorder()

	h.GetBlockByNumber(w, req)

	expectedRPCRequest := models.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{"0x1", true},
		ID:      2,
	}
	mockClient.AssertCalled(t, "SendRPCRequest", expectedRPCRequest)

	res := w.Result()
	assert.Equal(t, http.StatusOK, res.StatusCode)

	var responseBody map[string]interface{}
	err := json.NewDecoder(res.Body).Decode(&responseBody)
	assert.NoError(t, err)
	assert.Equal(t, "0x1", responseBody["result"])

	mockClient.AssertExpectations(t)
}

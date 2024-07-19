package server

import (
	"net/http"
	"net/http/httptest"
	"simple-blockchain-client/internal/handlers"
	"simple-blockchain-client/internal/rpc"
	"simple-blockchain-client/pkg/models"
	"strings"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func setupRouter(mockClient rpc.ClientInterface) *mux.Router {
	handlers := handlers.NewHandlers(mockClient)

	router := mux.NewRouter()
	router.HandleFunc("/blockNumber", handlers.GetBlockNumber).Methods("POST")
	router.HandleFunc("/blockByNumber", handlers.GetBlockByNumber).Methods("POST")

	return router
}

func TestStart(t *testing.T) {
	mockClient := new(rpc.MockRPCClient)
	mockClient.On("SendRPCRequest", mock.Anything).Return(models.RPCResponse{
		Result: "0x1",
		ID:     1,
	}, nil)

	router := setupRouter(mockClient)

	server := httptest.NewServer(router)
	defer server.Close()

	res, err := http.Post(server.URL+"/blockNumber", "application/json", nil)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	res, err = http.Post(server.URL+"/blockByNumber", "application/json", strings.NewReader(`{"blockNumber": "0x1"}`))
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, res.StatusCode)

	mockClient.AssertExpectations(t)
}

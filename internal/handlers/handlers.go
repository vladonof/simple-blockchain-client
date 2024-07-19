package handlers

import (
	"encoding/json"
	"net/http"
	"trustwallet-simple-blockchain-client/internal/rpc"
	"trustwallet-simple-blockchain-client/pkg/models"
)

type Handlers struct {
	Client rpc.ClientInterface
}

func NewHandlers(client rpc.ClientInterface) *Handlers {
	return &Handlers{Client: client}
}

func (h *Handlers) GetBlockNumber(w http.ResponseWriter, r *http.Request) {
	rpcRequest := models.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      3,
	}

	handleRPCRequest(h.Client, w, rpcRequest)
}

func (h *Handlers) GetBlockByNumber(w http.ResponseWriter, r *http.Request) {
	var params struct {
		BlockNumber string `json:"blockNumber"`
	}
	if err := json.NewDecoder(r.Body).Decode(&params); err != nil {
		respondWithJSON(w, http.StatusBadRequest, map[string]string{"error": err.Error()})
		return
	}

	rpcRequest := models.RPCRequest{
		JSONRPC: "2.0",
		Method:  "eth_getBlockByNumber",
		Params:  []interface{}{params.BlockNumber, true},
		ID:      2,
	}

	handleRPCRequest(h.Client, w, rpcRequest)
}

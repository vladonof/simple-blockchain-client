package handlers

import (
	"encoding/json"
	"net/http"
	"trustwallet-simple-blockchain-client/internal/rpc"
	"trustwallet-simple-blockchain-client/pkg/models"
)

func respondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

func handleRPCRequest(client rpc.ClientInterface, w http.ResponseWriter, rpcRequest models.RPCRequest) {
	rpcResponse, err := client.SendRPCRequest(rpcRequest)
	if err != nil {
		respondWithJSON(w, http.StatusInternalServerError, map[string]string{"error": err.Error()})
		return
	}

	response := struct {
		Result  interface{} `json:"result"`
		Error   interface{} `json:"error"`
		JSONRPC string      `json:"jsonrpc"`
		ID      int         `json:"id"`
	}{
		JSONRPC: "2.0",
		Result:  rpcResponse.Result,
		Error:   rpcResponse.Error,
		ID:      rpcResponse.ID,
	}

	respondWithJSON(w, http.StatusOK, response)
}

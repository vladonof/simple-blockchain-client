package server

import (
	"net/http"
	"trustwallet-simple-blockchain-client/internal/handlers"
	"trustwallet-simple-blockchain-client/internal/rpc"

	"github.com/gorilla/mux"
)

func Start() {
	client := rpc.NewRPCClient("https://polygon-rpc.com/")
	handlers := handlers.NewHandlers(client)

	router := mux.NewRouter()
	router.HandleFunc("/blockNumber", handlers.GetBlockNumber).Methods("POST")
	router.HandleFunc("/blockByNumber", handlers.GetBlockByNumber).Methods("POST")

	http.ListenAndServe(":8080", router)
}

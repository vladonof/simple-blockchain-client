package rpc

import "trustwallet-simple-blockchain-client/pkg/models"

type ClientInterface interface {
	SendRPCRequest(req models.RPCRequest) (models.RPCResponse, error)
}

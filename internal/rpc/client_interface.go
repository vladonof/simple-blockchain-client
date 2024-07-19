package rpc

import "simple-blockchain-client/pkg/models"

type ClientInterface interface {
	SendRPCRequest(req models.RPCRequest) (models.RPCResponse, error)
}

package rpc

import (
	"simple-blockchain-client/pkg/models"

	"github.com/stretchr/testify/mock"
)

type MockRPCClient struct {
	mock.Mock
}

func (m *MockRPCClient) SendRPCRequest(req models.RPCRequest) (models.RPCResponse, error) {
	args := m.Called(req)
	return args.Get(0).(models.RPCResponse), args.Error(1)
}

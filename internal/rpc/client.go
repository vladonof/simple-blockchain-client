package rpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"trustwallet-simple-blockchain-client/pkg/models"
)

type RPCClient struct {
	Client *http.Client
	URL    string
}

func NewRPCClient(url string) *RPCClient {
	return &RPCClient{
		URL:    url,
		Client: &http.Client{},
	}
}

func (c *RPCClient) SendRPCRequest(req models.RPCRequest) (models.RPCResponse, error) {
	jsonData, err := json.Marshal(req)
	if err != nil {
		return models.RPCResponse{}, fmt.Errorf("error marshalling request: %v", err)
	}

	resp, err := c.Client.Post(c.URL, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return models.RPCResponse{}, fmt.Errorf("error sending request: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return models.RPCResponse{}, fmt.Errorf("error reading response: %v", err)
	}

	var rpcResponse models.RPCResponse
	if err := json.Unmarshal(body, &rpcResponse); err != nil {
		return models.RPCResponse{}, fmt.Errorf("error unmarshalling response: %v", err)
	}

	return rpcResponse, nil
}

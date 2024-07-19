package models

type RPCRequest struct {
	JSONRPC string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params,omitempty"`
	ID      int           `json:"id"`
}

type RPCResponse struct {
	Result interface{} `json:"result"`
	Error  interface{} `json:"error"`
	ID     int         `json:"id"`
}

package live

import "encoding/json"

// RPCRequest is the struct representing a JSON-RPC over socket.io request.
type RPCRequest struct {
	ID     string            `json:"id,omitempty"`
	Params []json.RawMessage `json:"params,omitempty"`
	Method string            `json:"method,omitempty"`
}

// RPCResponse is a struct representing a JSON-RPC over socket.io response.
type RPCResponse struct {
	ID     string          `json:"id,omitempty"`
	Result json.RawMessage `json:"result,omitempty"`
	Error  *RPCError       `json:"error,omitempty"`
}

// RPCError is a struct representing a JSON-RPC over socket.io error.
type RPCError struct {
	Code    int32       `json:"code,omitempty"`
	Message string      `json:"message,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

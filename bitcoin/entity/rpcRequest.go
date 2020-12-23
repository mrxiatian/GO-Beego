package entity

//norm:标椎、规范
type RPCNorm struct {
	Id      int64         `json:"id"`
	Jsonrpc string        `json:"jsonrpc"`
	Method  string        `json:"method"`
	Params  []interface{} `json:"params"`
}

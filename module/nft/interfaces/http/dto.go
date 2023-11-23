package http

type GreeterContractEventLogResponse struct {
	TxHash          string                 `json:"txHash"`
	ContractAddress string                 `json:"contractAddress"`
	Event           string                 `json:"event"`
	Metadata        map[string]interface{} `json:"metadata"`
	BlockTimestamp  uint64                 `json:"blockTimestamp"`
}

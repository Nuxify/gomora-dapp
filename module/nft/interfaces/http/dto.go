package http

import "gomora-dapp/module/nft/domain/entity"

type GreeterContractEventLogResponse struct {
	TxHash          string                 `json:"txHash"`
	LogIndex        uint                   `json:"logIndex"`
	ContractAddress string                 `json:"contractAddress"`
	Chain           entity.Chain           `json:"chain"`
	Event           entity.Event           `json:"event"`
	Metadata        map[string]interface{} `json:"metadata"`
	BlockTimestamp  uint64                 `json:"blockTimestamp"`
}

package types

import "time"

type CreateNFTContractEventLog struct {
	TxHash          string
	ContractAddress string
	Event           string
	Metadata        string
	BlockTimestamp  time.Time
}

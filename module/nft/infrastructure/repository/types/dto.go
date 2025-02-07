package types

import "time"

type CreateGreeterContractEventLog struct {
	TxHash          string
	LogIndex        uint
	ContractAddress string
	Event           string
	Metadata        string
	BlockTimestamp  time.Time
}

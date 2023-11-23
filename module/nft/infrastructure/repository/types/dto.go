package types

import "time"

type CreateGreeterContractEventLog struct {
	TxHash          string
	ContractAddress string
	Event           string
	Metadata        string
	BlockTimestamp  time.Time
}

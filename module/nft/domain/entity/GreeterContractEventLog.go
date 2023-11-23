package entity

import "time"

const (
	LogSetGreeting string = "LogSetGreeting"
)

// GreeterContractEventLog holds the  greeter contract event log entity fields
type GreeterContractEventLog struct {
	TxHash          string `db:"tx_hash"`
	ContractAddress string `db:"contract_address"`
	Event           string
	Metadata        string
	BlockTimestamp  time.Time `db:"block_timestamp"`
}

// GetModelName returns the model name of record entity that can be used for naming schemas
func (entity *GreeterContractEventLog) GetModelName() string {
	return "greeter_contract_event_logs"
}

package entity

import "time"

type Chain string
type Event string

const (
	LogSetGreetingEvent Event = "LogSetGreeting"

	EthereumChain Chain = "ETHEREUM"
)

// GreeterContractEventLog holds the  greeter contract event log entity fields
type GreeterContractEventLog struct {
	TxHash          string    `db:"tx_hash"`
	LogIndex        uint      `db:"log_index"`
	ContractAddress string    `db:"contract_address"`
	Chain           Chain     // enum
	Event           Event     `db:"event"` // enum
	Metadata        string    `db:"metadata"`
	BlockTimestamp  time.Time `db:"block_timestamp"`
}

// GetModelName returns the model name of record entity that can be used for naming schemas
func (entity *GreeterContractEventLog) GetModelName() string {
	return "greeter_contract_event_logs"
}

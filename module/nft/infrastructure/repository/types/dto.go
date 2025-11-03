package types

import (
	"time"

	"gomora-dapp/module/nft/domain/entity"
)

type CreateGreeterContractEventLog struct {
	TxHash          string
	LogIndex        uint
	ContractAddress string
	Chain           entity.Chain
	Event           entity.Event
	Metadata        string
	BlockTimestamp  time.Time
}

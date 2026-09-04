package repository

import "gomora-dapp/module/nft/domain/entity"

// NFTQueryRepositoryInterface holds the methods for the nft query repository
type NFTQueryRepositoryInterface interface {
	// SelectGreeterContractEventLogs selects the greeter contract event logs
	SelectGreeterContractEventLogs() ([]entity.GreeterContractEventLog, error)
}

package repository

import repositoryTypes "gomora-dapp/module/nft/infrastructure/repository/types"

// NFTCommandRepositoryInterface holds the implementable methods for the nft command repository
type NFTCommandRepositoryInterface interface {
	InsertGreeterContractEventLog(data repositoryTypes.CreateGreeterContractEventLog) error
}

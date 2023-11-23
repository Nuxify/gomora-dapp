package repository

import repositoryTypes "gomora-dapp/module/nft/infrastructure/repository/types"

// NFTCommandRepositoryInterface holds the implementable methods for the nft command repository
type NFTCommandRepositoryInterface interface {
	InsertNFTContractEventLog(data repositoryTypes.CreateNFTContractEventLog) error
}

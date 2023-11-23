package repository

import "gomora-dapp/module/nft/domain/entity"

// NFTQueryRepositoryInterface holds the methods for the nft query repository
type NFTQueryRepositoryInterface interface {
	SelectNFTContractEventLogs() ([]entity.NFTGreeterContractEventLogs, error)
}

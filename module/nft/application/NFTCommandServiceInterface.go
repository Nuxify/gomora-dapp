package application

import (
	"context"

	"gomora-dapp/module/nft/infrastructure/service/types"
)

// NFTCommandServiceInterface holds the implementable methods for the nft command service
type NFTCommandServiceInterface interface {
	// CreateNFTLogSetGreeting creates a new log set greeting event
	CreateNFTLogSetGreeting(ctx context.Context, txHash string, logIndex uint, contractAddress string, data types.CreateNFTLogSetGreeting) error
}

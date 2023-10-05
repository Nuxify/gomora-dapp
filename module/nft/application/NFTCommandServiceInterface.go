package application

import (
	"context"

	"gomora-dapp/module/nft/infrastructure/service/types"
)

// NFTCommandServiceInterface holds the implementable methods for the nft command service
type NFTCommandServiceInterface interface {
	CreateNFTLogSetGreeting(ctx context.Context, txHash, contractAddress string, data types.CreateNFTLogSetGreeting) error
}

package application

import (
	"context"
)

// NFTQueryServiceInterface holds the implementable methods for the nft query service
type NFTQueryServiceInterface interface {
	GetNFTByID(ctx context.Context, ID int64) map[string]interface{}
}

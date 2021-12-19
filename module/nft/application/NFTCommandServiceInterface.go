package application

import (
	"context"
)

// NFTCommandServiceInterface holds the implementable methods for the nft command service
type NFTCommandServiceInterface interface {
	UploadMint(ctx context.Context, data map[string]interface{}) error
}

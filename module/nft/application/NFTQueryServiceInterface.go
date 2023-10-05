package application

import (
	"context"
)

// NFTQueryServiceInterface holds the implementable methods for the nft query service
type NFTQueryServiceInterface interface {
	GetGreeting(ctx context.Context) (string, error)
}

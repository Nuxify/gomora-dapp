package application

import (
	"context"
	"gomora-dapp/module/nft/domain/entity"
)

// NFTQueryServiceInterface holds the implementable methods for the nft query service
type NFTQueryServiceInterface interface {
	GetGreeting(ctx context.Context) (string, error)
	GetGreeterContractEventLogs(ctx context.Context) ([]entity.GreeterContractEventLog, error)
}

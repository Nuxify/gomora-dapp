package application

import (
	"context"
	"gomora-dapp/module/nft/domain/entity"
)

// NFTQueryServiceInterface holds the implementable methods for the nft query service
type NFTQueryServiceInterface interface {
	// GetGreeting gets the greeting message
	GetGreeting(ctx context.Context) (string, error)
	// GetGreeterContractEventLogs gets the greeter contract event logs
	GetGreeterContractEventLogs(ctx context.Context) ([]entity.GreeterContractEventLog, error)
}

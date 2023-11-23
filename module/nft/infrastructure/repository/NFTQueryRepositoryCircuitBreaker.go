package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"gomora-dapp/module/nft/domain/entity"
	"gomora-dapp/module/nft/domain/repository"
)

// NFTQueryRepositoryCircuitBreaker is the circuit breaker for the nft query repository
type NFTQueryRepositoryCircuitBreaker struct {
	repository.NFTQueryRepositoryInterface
}

// SelectNFTContractEventLogs is a decorator for the select nft greeter contract event logs repository
func (repository *NFTQueryRepositoryCircuitBreaker) SelectNFTContractEventLogs() ([]entity.NFTGreeterContractEventLogs, error) {
	output := make(chan []entity.NFTGreeterContractEventLogs, 1)
	hystrix.ConfigureCommand("select_nft_greeter_contract_event_logs", config.Settings())
	errors := hystrix.Go("select_nft_greeter_contract_event_logs", func() error {
		user, err := repository.NFTQueryRepositoryInterface.SelectNFTContractEventLogs()
		if err != nil {
			return err
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.NFTGreeterContractEventLogs{}, err
	}
}

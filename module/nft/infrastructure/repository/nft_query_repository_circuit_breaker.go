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

// SelectGreeterContractEventLogs is a decorator for the select greeter contract event logs repository
func (repository *NFTQueryRepositoryCircuitBreaker) SelectGreeterContractEventLogs() ([]entity.GreeterContractEventLog, error) {
	output := make(chan []entity.GreeterContractEventLog, 1)
	errChan := make(chan error, 1)

	hystrix.ConfigureCommand("select_greeter_contract_event_logs", config.Settings())
	errors := hystrix.Go("select_greeter_contract_event_logs", func() error {
		logs, err := repository.NFTQueryRepositoryInterface.SelectGreeterContractEventLogs()
		if err != nil {
			errChan <- err
			return nil
		}

		output <- logs
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errChan:
		return []entity.GreeterContractEventLog{}, err
	case err := <-errors:
		return []entity.GreeterContractEventLog{}, err
	}
}

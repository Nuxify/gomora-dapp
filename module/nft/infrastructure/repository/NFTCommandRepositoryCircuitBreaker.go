package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "gomora-dapp/configs/hystrix"
	"gomora-dapp/module/nft/domain/repository"
	repositoryTypes "gomora-dapp/module/nft/infrastructure/repository/types"
)

// UserCommandRepositoryCircuitBreaker circuit breaker for user command repository
type NFTCommandRepositoryCircuitBreaker struct {
	repository.NFTCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// InsertNFTContractEventLog is the decorator for the nft repository insert nft greeter contract event log
func (repository *NFTCommandRepositoryCircuitBreaker) InsertNFTContractEventLog(data repositoryTypes.CreateNFTContractEventLog) error {
	output := make(chan bool, 1)
	hystrix.ConfigureCommand("insert_nft_greeter_contract_event_log", config.Settings())
	errors := hystrix.Go("insert_nft_greeter_contract_event_log", func() error {
		err := repository.NFTCommandRepositoryInterface.InsertNFTContractEventLog(data)
		if err != nil {
			return err
		}

		output <- true
		return nil
	}, nil)

	select {
	case <-output:
		return nil
	case err := <-errors:
		return err
	}
}

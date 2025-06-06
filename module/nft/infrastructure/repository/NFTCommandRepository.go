package repository

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"gomora-dapp/infrastructures/database/mysql/types"
	apiError "gomora-dapp/internal/errors"
	"gomora-dapp/module/nft/domain/entity"
	repositoryTypes "gomora-dapp/module/nft/infrastructure/repository/types"
)

// NFTCommandRepository handles the nft command repository logic
type NFTCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// InsertGreeterContractEventLog insert a greeter contract event log record
func (repository *NFTCommandRepository) InsertGreeterContractEventLog(data repositoryTypes.CreateGreeterContractEventLog) error {
	eventLog := entity.GreeterContractEventLog{
		TxHash:          data.TxHash,
		LogIndex:        data.LogIndex,
		ContractAddress: data.ContractAddress,
		Event:           data.Event,
		Metadata:        data.Metadata,
		BlockTimestamp:  data.BlockTimestamp,
	}

	stmt := fmt.Sprintf("INSERT INTO %s (tx_hash,log_index,contract_address,event,metadata,block_timestamp) "+
		"VALUES (:tx_hash,:log_index,:contract_address,:event,:metadata,:block_timestamp)", eventLog.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, eventLog)
	if err != nil {
		log.Println(err)
		if strings.Contains(err.Error(), "Duplicate entry") {
			return errors.New(apiError.DuplicateRecord)
		}
		return errors.New(apiError.DatabaseError)
	}

	return nil
}

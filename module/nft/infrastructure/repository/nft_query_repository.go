package repository

import (
	"errors"
	"fmt"
	"log"

	"gomora-dapp/infrastructures/database/mysql/types"
	apiError "gomora-dapp/internal/errors"
	"gomora-dapp/module/nft/domain/entity"
)

// NFTQueryRepository handles the nft query repository logic
type NFTQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectGreeterContractEventLogs select greeter contract event logs
func (repository *NFTQueryRepository) SelectGreeterContractEventLogs() ([]entity.GreeterContractEventLog, error) {
	var eventLog entity.GreeterContractEventLog
	var eventLogs []entity.GreeterContractEventLog

	stmt := fmt.Sprintf("SELECT * FROM %s ORDER BY block_timestamp DESC", eventLog.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &eventLogs)
	if err != nil {
		log.Println(err)
		return []entity.GreeterContractEventLog{}, errors.New(apiError.DatabaseError)
	} else if len(eventLogs) == 0 {
		return []entity.GreeterContractEventLog{}, errors.New(apiError.MissingRecord)
	}

	return eventLogs, nil
}

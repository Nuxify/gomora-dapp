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

// SelectNFTContractEventLogs select marketplace active listings
func (repository *NFTQueryRepository) SelectNFTContractEventLogs() ([]entity.NFTGreeterContractEventLogs, error) {
	var listing entity.NFTGreeterContractEventLogs
	var listings []entity.NFTGreeterContractEventLogs

	stmt := fmt.Sprintf("SELECT * FROM %s ORDER BY block_timestamp DESC", listing.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &listings)
	if err != nil {
		log.Println(err)
		return []entity.NFTGreeterContractEventLogs{}, errors.New(apiError.DatabaseError)
	} else if len(listings) == 0 {
		return []entity.NFTGreeterContractEventLogs{}, errors.New(apiError.MissingRecord)
	}

	return listings, nil
}

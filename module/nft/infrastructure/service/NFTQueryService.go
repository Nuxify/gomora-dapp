package service

import (
	"context"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	greeter "gomora-dapp/infrastructures/smartcontracts/greeter"
	"gomora-dapp/module/nft/domain/entity"
	"gomora-dapp/module/nft/domain/repository"
)

// NFTQueryService handles the nft query service logic
type NFTQueryService struct {
	repository.NFTQueryRepositoryInterface
	GreeterContractInstance *greeter.Greeter
}

// GetGreeting get latest greeting
func (service *NFTQueryService) GetGreeting(ctx context.Context) (string, error) {
	greeting, err := service.GreeterContractInstance.Greet(&bind.CallOpts{})
	if err != nil {
		return "", err
	}

	return greeting, nil
}

// GetGreeterContractEventLogs get greeter contract event logs
func (service *NFTQueryService) GetGreeterContractEventLogs(ctx context.Context) ([]entity.GreeterContractEventLog, error) {
	res, err := service.NFTQueryRepositoryInterface.SelectGreeterContractEventLogs()
	if err != nil {
		return []entity.GreeterContractEventLog{}, err
	}

	return res, nil
}

// GetNFTByID retrieves the nft provided by its id
// TODO: Example common code to get metadata by token id
// func (service *NFTQueryService) GetNFTByID(ctx context.Context, ID int64) map[string]interface{} {
// 	// placeHolderMetadata := map[string]interface{}{
// 	// 	"name":         "Unknown NFT",
// 	// 	"description":  "Unknown NFT",
// 	// 	"image":        "ipfs://<your-cid-here>",
// 	// 	"external_url": "ipfs://<your-cid-here>",
// 	// 	"attributes":   []string{},
// 	// }

// 	// check if token id is minted
// 	// doesExist, err := service.GreeterContractInstance.Exists(&bind.CallOpts{}, big.NewInt(ID))
// 	// if !doesExist || err != nil {
// 	// 	return placeHolderMetadata
// 	// }

// 	var metadata map[string]interface{}

// 	// rootPath, _ := os.Getwd()
// 	// plan, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.json", rootPath, constants.NFTMetadataPath, ID))
// 	// err = json.Unmarshal(plan, &metadata)
// 	// if err != nil {
// 	// 	log.Println("[Error] cannot unmarshal metadata json", err)
// 	// 	return placeHolderMetadata
// 	// }

// 	return metadata
// }

package service

import (
	"context"

	greeter "gomora-dapp/infrastructures/smartcontracts/greeter"
)

// NFTQueryService handles the nft query service logic
type NFTQueryService struct {
	GreeterContractInstance *greeter.Greeter
}

// GetNFTByID retrieves the nft provided by its id
func (service *NFTQueryService) GetNFTByID(ctx context.Context, ID int64) map[string]interface{} {
	// placeHolderMetadata := map[string]interface{}{
	// 	"name":         "Unknown NFT",
	// 	"description":  "Unknown NFT",
	// 	"image":        "ipfs://<your-cid-here>",
	// 	"external_url": "ipfs://<your-cid-here>",
	// 	"attributes":   []string{},
	// }

	// check if token id is minted
	// doesExist, err := service.GreeterContractInstance.Exists(&bind.CallOpts{}, big.NewInt(ID))
	// if !doesExist || err != nil {
	// 	return placeHolderMetadata
	// }

	var metadata map[string]interface{}

	// rootPath, _ := os.Getwd()
	// plan, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.json", rootPath, constants.NFTMetadataPath, ID))
	// err = json.Unmarshal(plan, &metadata)
	// if err != nil {
	// 	log.Println("[Error] cannot unmarshal metadata json", err)
	// 	return placeHolderMetadata
	// }

	return metadata
}

package service

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"os"

	"gomora-dapp/infrastructures/smartcontracts"
	"gomora-dapp/internal/constants"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

// NFTQueryService handles the nft query service logic
type NFTQueryService struct {
	SampleContractContractInstance *smartcontracts.Smartcontracts
}

// GetNFTByID retrieves the nft provided by its id
func (service *NFTQueryService) GetNFTByID(ctx context.Context, ID int64) map[string]interface{} {
	placeHolderMetadata := map[string]interface{}{
		"name":         "Unknown NFT",
		"description":  "Unknown NFT",
		"image":        "ipfs://<your-cid-here>",
		"external_url": "ipfs://<your-cid-here>",
		"attributes":   []string{},
	}

	// check if token id is minted
	doesExist, err := service.SampleContractContractInstance.Exists(&bind.CallOpts{}, big.NewInt(ID))
	if !doesExist || err != nil {
		return placeHolderMetadata
	}

	var metadata map[string]interface{}

	rootPath, _ := os.Getwd()
	plan, _ := ioutil.ReadFile(fmt.Sprintf("%s/%s/%d.json", rootPath, constants.NFTMetadataPath, ID))
	err = json.Unmarshal(plan, &metadata)
	if err != nil {
		log.Println("[Error] cannot unmarshal metadata json", err)
		return placeHolderMetadata
	}

	return metadata
}

package interfaces

import (
	"context"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	nftService "gomora-dapp/module/nft/infrastructure/service"
)

var (
	SampleContractContractAddress common.Address
	SampleContractContractABI     abi.ABI
)

func SampleContractEventWatcher() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{SampleContractContractAddress},
	}

	logs := make(chan types.Log)
	sub, err := EthWsClient.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		panic(err)
	}

	// for nft command service
	commandService := &nftService.NFTCommandService{}

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			// get topics, topic 0 is signature of event, topic 1 is first indexed
			var topics [4]string
			for i := range vLog.Topics {
				topics[i] = vLog.Topics[i].Hex()
			}

			eventSignature := topics[0]

			/// on firstEvent
			mintEventName := "Mint"
			mintPurchaseMap := map[string]interface{}{}
			mintTopic := crypto.Keccak256Hash([]byte("Mint(address,uint256,string)"))

			err := SampleContractContractABI.UnpackIntoMap(mintPurchaseMap, mintEventName, vLog.Data)
			if err == nil && eventSignature == mintTopic.Hex() {
				event := map[string]interface{}{
					"token_id": mintPurchaseMap["tokenID"].(*big.Int).Int64(),
					"tier":     mintPurchaseMap["tier"].(string),
					"wallet":   topics[1],
				}

				err := commandService.UploadMint(context.TODO(), event)
				if err != nil {
					log.Println("[error] Mint cannot upload mint", err)
				}
			}

			/// on secondEvent
			BatchMintEventName := "BatchMint"
			BatchMintPurchaseMap := map[string]interface{}{}
			BatchMintTopic := crypto.Keccak256Hash([]byte("BatchMint(address,uint256,string)"))

			err = SampleContractContractABI.UnpackIntoMap(BatchMintPurchaseMap, BatchMintEventName, vLog.Data)
			if err == nil && eventSignature == BatchMintTopic.Hex() {
				event := map[string]interface{}{
					"token_id": mintPurchaseMap["tokenID"].(*big.Int).Int64(),
					"tier":     mintPurchaseMap["tier"].(string),
					"wallet":   topics[1],
				}

				err := commandService.UploadMint(context.TODO(), event)
				if err != nil {
					log.Println("[error] BatchMint cannot upload mint", err)
				}
			}
		}
	}
}

package interfaces

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"

	serviceTypes "gomora-dapp/module/nft/infrastructure/service/types"
)

var (
	GreeterContractAddress common.Address
	GreeterContractABI     abi.ABI
)

// GreeterEventListener greeter contract indexer
func GreeterEventListener() {
	query := ethereum.FilterQuery{
		Addresses: []common.Address{GreeterContractAddress},
	}

	logs := make(chan types.Log)
	sub := event.Resubscribe(2*time.Second, func(ctx context.Context) (event.Subscription, error) {
		return EthWsClient.SubscribeFilterLogs(ctx, query, logs)
	})

	// for nft command service
	commandService := NFTCommandServiceDI()

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

			txHash := vLog.TxHash.Hex()
			block, _ := EthHttpClient.BlockByNumber(context.TODO(), big.NewInt(int64(vLog.BlockNumber)))
			blockTimestamp := int64(block.Time())
			eventSignature := topics[0]

			/// LogMint event
			eventName := "Mint"
			eventData := map[string]interface{}{}
			mintTopic := crypto.Keccak256Hash([]byte("Mint(address,uint256,string)"))

			err := GreeterContractABI.UnpackIntoMap(eventData, eventName, vLog.Data)
			if err == nil && eventSignature == mintTopic.Hex() {
				event := serviceTypes.Upload{
					TxHash:         txHash,
					BlockTimestamp: blockTimestamp,
					TokenID:        eventData["tokenID"].(*big.Int).Int64(),
					Tier:           eventData["tier"].(string),
					Wallet:         common.HexToAddress(topics[1]).String(),
				}

				err := commandService.UploadMint(context.TODO(), event)
				if err != nil {
					log.Println("[error] Mint cannot upload mint", err)
				}
			}

			/// LogBatchMint event
			eventData = map[string]interface{}{}
			mintTopic = crypto.Keccak256Hash([]byte("BatchMint(address,uint256,string)"))

			err = GreeterContractABI.UnpackIntoMap(eventData, eventName, vLog.Data)
			if err == nil && eventSignature == mintTopic.Hex() {
				event := serviceTypes.Upload{
					TxHash:         txHash,
					BlockTimestamp: blockTimestamp,
					TokenID:        eventData["tokenID"].(*big.Int).Int64(),
					Tier:           eventData["tier"].(string),
					Wallet:         common.HexToAddress(topics[1]).String(),
				}

				err := commandService.UploadMint(context.TODO(), event)
				if err != nil {
					log.Println("[error] Mint cannot upload mint", err)
				}
			}
		}
	}
}

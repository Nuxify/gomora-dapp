package interfaces

import (
	"context"
	"log"
	"math/big"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/event"

	ethRPC "gomora-dapp/internal/ethereum"
	"gomora-dapp/module/nft/infrastructure/service"
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
	nftCommandService := NFTCommandServiceDI()

	for {
		select {
		case err := <-sub.Err():
			panic(err)
		case vLog := <-logs:
			greeterEventsHandler(nftCommandService, vLog, time.Now().Unix())
		}
	}
}

// GreeterEventListenerReplayer replay greeter events from certain blocks
func GreeterEventListenerReplayer(fromBlock, toBlock int64) error {
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(fromBlock),
		ToBlock:   big.NewInt(toBlock),
		Addresses: []common.Address{GreeterContractAddress},
	}

	logs, err := EthHttpClient.FilterLogs(context.Background(), query)
	if err != nil {
		log.Println(err)
		return err
	}

	log.Println("REPLAY:", len(logs), fromBlock, toBlock)

	// for nft command service
	nftCommandService := NFTCommandServiceDI()

	var lastKnownTimestamp int64

	for _, vLog := range logs {
		// get block timestamp
		block, err := EthHttpClient.BlockByNumber(context.TODO(), big.NewInt(int64(vLog.BlockNumber)))
		if err != nil {
			log.Println(err)
		} else {
			lastKnownTimestamp = int64(block.Time())
		}

		greeterEventsHandler(nftCommandService, vLog, lastKnownTimestamp)
	}

	return nil
}

// GreeterPollFilter poll filter missed events backup
func GreeterPollFilter(rpcURL string) {
	// create a new filter
	query := ethereum.FilterQuery{
		Addresses: []common.Address{GreeterContractAddress},
	}

	filterID, err := ethRPC.NewFilter(rpcURL, query)
	if err != nil {
		panic(err)
	}

	log.Println("POLL FILTER ID:", filterID)

	// for nft command service
	nftCommandService := NFTCommandServiceDI()

	var lastKnownTimestamp int64

	// get filter changes
	for {
		logs, err := ethRPC.GetFilterChanges(rpcURL, filterID)
		if err != nil {
			log.Println(err)

			// check if filter not found
			if strings.Contains(err.Error(), "not found") {
				filterID, _ = ethRPC.NewFilter(rpcURL, query)
				if err != nil {
					log.Println(err)
				}

				log.Println("POLL FILTER ID:", filterID)
			}
			continue
		}

		log.Println("POLL FILTER:", len(logs))

		for _, vLog := range logs {
			// get block timestamp
			block, err := EthHttpClient.BlockByNumber(context.TODO(), big.NewInt(int64(vLog.BlockNumber)))
			if err != nil {
				log.Println(err)
			} else {
				lastKnownTimestamp = int64(block.Time())
			}

			greeterEventsHandler(nftCommandService, vLog, lastKnownTimestamp)
		}

		time.Sleep(30 * time.Second)
	}
}

// Handle greeter contract events
func greeterEventsHandler(nftCommandService *service.NFTCommandService, vLog types.Log, blockTimestampInSeconds int64) {
	// get topics, topic 0 is signature of event, topic 1 is first indexed
	var topics [4]string
	for i := range vLog.Topics {
		topics[i] = vLog.Topics[i].Hex()
	}

	txHash := vLog.TxHash.Hex()
	eventSignature := topics[0]

	/// LogMint event
	eventName := "LogMint"
	eventData := map[string]interface{}{}
	mintTopic := crypto.Keccak256Hash([]byte("LogMint(address,uint256,string)"))

	err := GreeterContractABI.UnpackIntoMap(eventData, eventName, vLog.Data)
	if err == nil && eventSignature == mintTopic.Hex() {
		event := serviceTypes.Upload{
			TxHash:         txHash,
			BlockTimestamp: blockTimestampInSeconds,
			TokenID:        eventData["tokenID"].(*big.Int).Int64(),
			Tier:           eventData["tier"].(string),
			Wallet:         common.HexToAddress(topics[1]).String(),
		}

		err := nftCommandService.UploadMint(context.TODO(), event)
		if err != nil {
			log.Println("[error] LogMint cannot upload mint", err)
		}
	}

	/// LogBatchMint event
	eventData = map[string]interface{}{}
	mintTopic = crypto.Keccak256Hash([]byte("LogBatchMint(address,uint256,string)"))

	err = GreeterContractABI.UnpackIntoMap(eventData, eventName, vLog.Data)
	if err == nil && eventSignature == mintTopic.Hex() {
		event := serviceTypes.Upload{
			TxHash:         txHash,
			BlockTimestamp: blockTimestampInSeconds,
			TokenID:        eventData["tokenID"].(*big.Int).Int64(),
			Tier:           eventData["tier"].(string),
			Wallet:         common.HexToAddress(topics[1]).String(),
		}

		err := nftCommandService.UploadMint(context.TODO(), event)
		if err != nil {
			log.Println("[error] LogBatchMint cannot upload mint", err)
		}
	}
}

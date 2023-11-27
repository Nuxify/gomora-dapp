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

	"gomora-dapp/infrastructures/smartcontracts/greeter"
	ethRPC "gomora-dapp/internal/ethereum/rpc"
	"gomora-dapp/module/nft/infrastructure/service"
	nftServiceTypes "gomora-dapp/module/nft/infrastructure/service/types"
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
			greeterEventsHandler(nftCommandService, vLog, true)
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

	for _, vLog := range logs {
		greeterEventsHandler(nftCommandService, vLog, false)
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

	// get filter changes
	for {
		logs, err := ethRPC.GetFilterChanges(rpcURL, filterID)
		if err != nil {
			log.Println(err)

			// check if filter not found
			if strings.Contains(err.Error(), "not found") {
				filterID, err = ethRPC.NewFilter(rpcURL, query)
				if err != nil {
					log.Println(err)
					time.Sleep(10 * time.Second)
					continue
				}

				log.Println("NEW POLL FILTER ID:", filterID)
			}

			time.Sleep(3 * time.Second)
			continue
		}

		log.Println("POLL FILTER LENGTH:", len(logs))

		for _, vLog := range logs {
			greeterEventsHandler(nftCommandService, vLog, false)
		}

		time.Sleep(20 * time.Minute) // short polling interval
	}
}

// Handle greeter contract events
func greeterEventsHandler(nftCommandService *service.NFTCommandService, vLog types.Log, isFromWS bool) {
	// get topics, topic 0 is signature of event, topic 1 is first indexed
	var topics [4]string
	for i := range vLog.Topics {
		topics[i] = vLog.Topics[i].Hex()
	}

	txHash := vLog.TxHash.Hex()
	eventSignature := topics[0]

	/// LogSetGreeting event
	mintTopic := crypto.Keccak256Hash([]byte("LogSetGreeting(string,uint256)"))
	if eventSignature == mintTopic.Hex() {
		eventName := "LogSetGreeting"
		eventData := greeter.GreeterLogSetGreeting{}

		err := GreeterContractABI.UnpackIntoInterface(&eventData, eventName, vLog.Data)
		if err == nil {
			event := nftServiceTypes.CreateNFTLogSetGreeting{
				Greeting:  eventData.Greeting,
				Timestamp: eventData.Timestamp.Uint64(),
				IsFromWS:  isFromWS,
			}

			err := nftCommandService.CreateNFTLogSetGreeting(context.TODO(), txHash, GreeterContractAddress.Hex(), event)
			if err != nil {
				log.Println("[error] LogSetGreeting cannot update", err)
			}
		}
	}
}

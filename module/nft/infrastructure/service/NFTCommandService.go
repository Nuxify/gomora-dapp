package service

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	greeter "gomora-dapp/infrastructures/smartcontracts/greeter"
	"gomora-dapp/module/nft/domain/entity"
	"gomora-dapp/module/nft/domain/repository"
	repositoryTypes "gomora-dapp/module/nft/infrastructure/repository/types"
	"gomora-dapp/module/nft/infrastructure/service/types"
)

// NFTCommandService handles the nft command service logic
type NFTCommandService struct {
	repository.NFTCommandRepositoryInterface
	GreeterContractInstance *greeter.Greeter
}

// --------- transaction logs methods

// CreateNFTLogSetGreeting create nft log set greeting event
func (service *NFTCommandService) CreateNFTLogSetGreeting(ctx context.Context, txHash string, logIndex uint, contractAddress string, data types.CreateNFTLogSetGreeting) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// optionally store it in a persistent storage like database to log the logs
	// for now we just print it
	fmt.Println("tx hash:", txHash, "log index:", logIndex, "contract:", contractAddress, "greeting:", data.Greeting, "block timestamp:", data.Timestamp)
	fmt.Println("Metadata:", string(output))

	// insert to event logs
	err = service.NFTCommandRepositoryInterface.InsertGreeterContractEventLog(repositoryTypes.CreateGreeterContractEventLog{
		TxHash:          txHash,
		LogIndex:        logIndex,
		ContractAddress: contractAddress,
		Event:           entity.LogSetGreeting,
		Metadata:        string(output),
		BlockTimestamp:  time.Unix(int64(data.Timestamp), 0),
	})
	if err != nil {
		return err
	}

	// you can also send notifications here and be sure it will be only triggered one-time
	if data.IsFromWS {
		fmt.Println("sending notifications")
	}

	return nil
}

// TODO: helpful when you want to save logs
// func saveToUploadLogs(text string) {
// 	rootPath, _ := os.Getwd()
// 	f, err := os.OpenFile(fmt.Sprintf("%s/%s", rootPath, constants.NFTMintUploadLogPath), os.O_APPEND|os.O_WRONLY, 0644)
// 	if err != nil {
// 		log.Println("[Error] error opening nft upload mint log file", err)
// 	}

// 	_, err = f.WriteString(fmt.Sprintf("%s\n", text))
// 	if err != nil {
// 		log.Println("[Error] error writing log in nft upload mint log")
// 	}

// 	if err = f.Close(); err != nil {
// 		log.Println("[Error] error closing nft upload mint file")
// 	}
// }

package service

import (
	"context"
	"encoding/json"
	"fmt"

	greeter "gomora-dapp/infrastructures/smartcontracts/greeter"
	"gomora-dapp/module/nft/infrastructure/service/types"
)

// NFTCommandService handles the nft command service logic
type NFTCommandService struct {
	GreeterContractInstance *greeter.Greeter
}

func (service *NFTCommandService) CreateNFTLogSetGreeting(ctx context.Context, txHash, contractAddress string, data types.CreateNFTLogSetGreeting) error {
	output, err := json.Marshal(data)
	if err != nil {
		return err
	}

	// optionally store it in a persistent storage like database to log the logs
	// for now we just print it
	fmt.Println("tx hash:", txHash, ", contract:", contractAddress, ", greeting:", data.Greeting, ", block timestamp:", data.Timestamp)
	fmt.Println("Metadata:", string(output))

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

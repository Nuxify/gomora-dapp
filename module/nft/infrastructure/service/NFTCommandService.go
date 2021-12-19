package service

import (
	"context"
	"fmt"
	"log"
	"os"

	"gomora-dapp/internal/constants"
)

// NFTCommandService handles the nft query service logic
type NFTCommandService struct{}

func (service *NFTCommandService) UploadMint(ctx context.Context, data map[string]interface{}) error {
	// TODO: logic to handle upload mint event sample
	return nil
}

func saveToUploadLogs(text string) {
	rootPath, _ := os.Getwd()
	f, err := os.OpenFile(fmt.Sprintf("%s/%s", rootPath, constants.NFTMintUploadLogPath), os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Println("[Error] error opening nft upload mint log file", err)
	}

	_, err = f.WriteString(fmt.Sprintf("%s\n", text))
	if err != nil {
		log.Println("[Error] error writing log in nft upload mint log")
	}

	if err = f.Close(); err != nil {
		log.Println("[Error] error closing nft upload mint file")
	}
}

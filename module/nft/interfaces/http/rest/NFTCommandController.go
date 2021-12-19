package rest

import (
	"gomora-dapp/module/nft/application"
)

// NFTCommandController request controller for nft command
type NFTCommandController struct {
	application.NFTCommandServiceInterface
}

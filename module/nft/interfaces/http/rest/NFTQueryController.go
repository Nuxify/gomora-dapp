package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"

	"github.com/go-chi/chi"

	"gomora-dapp/interfaces/http/rest/viewmodels"
	"gomora-dapp/internal/constants"
	"gomora-dapp/internal/errors"
	"gomora-dapp/module/nft/application"
)

// NFTQueryController request controller for nft query
type NFTQueryController struct {
	application.NFTQueryServiceInterface
}

// GetNFTByID retrieves the nft by id
func (controller *NFTQueryController) GetNFTByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "tokenID")

	nftID, err := strconv.Atoi(idStr)
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusUnprocessableEntity,
			Success:   false,
			Message:   "Invalid nft ID",
			ErrorCode: errors.InvalidRequestPayload,
		}

		response.JSON(w)
		return
	}

	metadata := controller.NFTQueryServiceInterface.GetNFTByID(context.TODO(), int64(nftID))

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(metadata)
	return
}

// GetNFTImage get nft image by filename
func (controller *NFTQueryController) GetNFTImage(w http.ResponseWriter, r *http.Request) {
	fileName := chi.URLParam(r, "fileName")
	rootPath, _ := os.Getwd()

	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", rootPath, constants.NFTImagePath, fileName))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Write(fileBytes)
	return
}

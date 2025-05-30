package rest

import (
	"context"
	"encoding/json"
	"net/http"

	"gomora-dapp/interfaces/http/rest/viewmodels"
	"gomora-dapp/internal/errors"
	"gomora-dapp/module/nft/application"
	types "gomora-dapp/module/nft/interfaces/http"
)

// NFTQueryController request controller for nft query
type NFTQueryController struct {
	application.NFTQueryServiceInterface
}

// GetGreeting get greeting message
func (controller *NFTQueryController) GetGreeting(w http.ResponseWriter, r *http.Request) {
	greeting, err := controller.NFTQueryServiceInterface.GetGreeting(context.TODO())
	if err != nil {
		response := viewmodels.HTTPResponseVM{
			Status:    http.StatusInternalServerError,
			Success:   false,
			Message:   "Cannot fetch greeting message.",
			ErrorCode: errors.EthRPCFailed,
		}

		response.JSON(w)
		return
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched greeting message.",
		Data: map[string]interface{}{
			"greeting": greeting,
		},
	}

	response.JSON(w)
}

// GetGreeterContractEventLogs get greeter contract event logs
func (controller *NFTQueryController) GetGreeterContractEventLogs(w http.ResponseWriter, r *http.Request) {
	res, err := controller.NFTQueryServiceInterface.GetGreeterContractEventLogs(context.TODO())
	if err != nil {
		var httpCode int
		var errorMsg string
		errorCode := err.Error()

		switch errorCode {
		case errors.MissingRecord:
			httpCode = http.StatusNotFound
			errorMsg = "No records found."
		default:
			httpCode = http.StatusInternalServerError
			errorMsg = "Database error."
			errorCode = errors.ServerError
		}

		response := viewmodels.HTTPResponseVM{
			Status:    httpCode,
			Success:   false,
			Message:   errorMsg,
			ErrorCode: err.Error(),
		}

		response.JSON(w)
		return
	}

	var logs []types.GreeterContractEventLogResponse

	for _, logRes := range res {
		var metadata map[string]interface{}
		// decode metadata
		_ = json.Unmarshal([]byte(logRes.Metadata), &metadata)

		logs = append(logs, types.GreeterContractEventLogResponse{
			TxHash:          logRes.TxHash,
			LogIndex:        logRes.LogIndex,
			ContractAddress: logRes.ContractAddress,
			Event:           logRes.Event,
			Metadata:        metadata,
			BlockTimestamp:  uint64(logRes.BlockTimestamp.Unix()),
		})
	}

	response := viewmodels.HTTPResponseVM{
		Status:  http.StatusOK,
		Success: true,
		Message: "Successfully fetched greeter contract event logs.",
		Data:    logs,
	}

	response.JSON(w)
}

// GetNFTByID retrieves the nft by id
// TODO: Example code when fetching nft metadata
// func (controller *NFTQueryController) GetNFTByID(w http.ResponseWriter, r *http.Request) {
// 	idStr := chi.URLParam(r, "tokenID")

// 	nftID, err := strconv.Atoi(idStr)
// 	if err != nil {
// 		response := viewmodels.HTTPResponseVM{
// 			Status:    http.StatusUnprocessableEntity,
// 			Success:   false,
// 			Message:   "Invalid nft ID",
// 			ErrorCode: errors.InvalidRequestPayload,
// 		}

// 		response.JSON(w)
// 		return
// 	}

// 	metadata := controller.NFTQueryServiceInterface.GetNFTByID(context.TODO(), int64(nftID))

// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	_ = json.NewEncoder(w).Encode(metadata)
// 	return
// }

// GetNFTImage get nft image by filename
// TODO: Example code for fetching nft image
// func (controller *NFTQueryController) GetNFTImage(w http.ResponseWriter, r *http.Request) {
// 	fileName := chi.URLParam(r, "fileName")
// 	rootPath, _ := os.Getwd()

// 	fileBytes, err := ioutil.ReadFile(fmt.Sprintf("%s/%s/%s", rootPath, constants.NFTImagePath, fileName))
// 	if err != nil {
// 		w.WriteHeader(http.StatusNotFound)
// 		return
// 	}

// 	w.WriteHeader(http.StatusOK)
// 	w.Header().Set("Content-Type", "application/octet-stream")
// 	w.Write(fileBytes)
// 	return
// }

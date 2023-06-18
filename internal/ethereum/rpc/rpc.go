package ethereum

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/core/types"
)

var client = &http.Client{
	Timeout: 60 * time.Second,
}

// GetFilterChanges implementation of eth_getFilterChanges
func GetFilterChanges(nodeURL string, filterID string) ([]types.Log, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_getFilterChanges",
		"params":  []string{filterID},
		"id":      1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", nodeURL, strings.NewReader(string(body)))
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var response struct {
		Result []types.Log `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	if response.Error != nil {
		return nil, errors.New(fmt.Sprintf("%d: %s", response.Error.Code, response.Error.Message))
	}

	return response.Result, nil
}

// NewFilter implementation of eth_newFilter
func NewFilter(nodeURL string, filter ethereum.FilterQuery) (string, error) {
	payload := map[string]interface{}{
		"jsonrpc": "2.0",
		"method":  "eth_newFilter",
		"params":  []interface{}{filter},
		"id":      1,
	}

	body, err := json.Marshal(payload)
	if err != nil {
		return "", err
	}

	req, err := http.NewRequest("POST", nodeURL, strings.NewReader(string(body)))
	if err != nil {
		return "", err
	}

	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var response struct {
		Result string `json:"result"`
		Error  *struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		} `json:"error"`
	}

	err = json.NewDecoder(resp.Body).Decode(&response)
	if err != nil {
		return "", err
	}

	if response.Error != nil {
		return "", errors.New(fmt.Sprintf("%d: %s", response.Error.Code, response.Error.Message))
	}

	return response.Result, nil
}

// references: https://ethereum.stackexchange.com/questions/131756/verified-go-signature-in-solidity-eip712-typeddata
// https://medium.com/alpineintel/issuing-and-verifying-eip-712-challenges-with-go-32635ca78aaf
package eip712

import (
	"crypto/ecdsa"
	"fmt"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/signer/core/apitypes"
)

// SignHash typed data ready for signing
// accepts [typedData] to be encoded
func SignHash(typedData apitypes.TypedData) ([]byte, error) {
	domainSeparator, err := typedData.HashStruct("EIP712Domain", typedData.Domain.Map())
	if err != nil {
		return []byte{}, err
	}

	typedDataHash, err := typedData.HashStruct(typedData.PrimaryType, typedData.Message)
	if err != nil {
		return []byte{}, err
	}

	rawData := []byte(fmt.Sprintf("\x19\x01%s%s", string(domainSeparator), string(typedDataHash)))
	hashedData := crypto.Keccak256(rawData)

	return hashedData, nil
}

// SignTypedData sign typed data with private key
// accepts [typedData] and [privKey]
func SignTypedData(typedData apitypes.TypedData, privKey *ecdsa.PrivateKey) (string, error) {
	hashedData, err := SignHash(typedData)
	if err != nil {
		return "", err
	}

	sig, err := crypto.Sign(hashedData, privKey)
	if err != nil {
		return "", err
	}

	// https://github.com/ethereum/go-ethereum/blob/55599ee95d4151a2502465e0afc7c47bd1acba77/internal/ethapi/api.go#L442
	sig[64] += 27

	return hexutil.Encode(sig), nil
}

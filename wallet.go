package elementsrpc

import (
	"encoding/json"
	"strings"

	"github.com/rddl-network/elements-rpc/types"
)

// BlindRawTransaction converts one or more outputs of a raw transaction into
// confidential ones using only wallet inputs.
func BlindRawTransaction(url string, params []string) (hex string, err error) {
	result, err := SendRequest(url, types.MethodBlindRawTransaction, params)
	if err != nil {
		return
	}
	hex = strings.ReplaceAll(string(result), "\"", "")
	return
}

// GetAddressInfo returns information about the given address.
func GetAddressInfo(url string, params []string) (transactionResult types.GetAddressInfoResult, err error) {
	result, err := SendRequest(url, types.MethodGetAddressInfo, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	return
}

// GetNewAddress returns a new address for receiving payments.
func GetNewAddress(url string, params []string) (address string, err error) {
	result, err := SendRequest(url, types.MethodGetNewAddress, params)
	if err != nil {
		return
	}
	address = strings.ReplaceAll(string(result), "\"", "")
	return
}

// GetTransaction retrieves a transaction from the chain.
func GetTransaction(url string, params []string) (transactionResult types.GetTransactionResult, err error) {
	result, err := SendRequest(url, types.MethodGetTransaction, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	return
}

// ReissueAsset creates more of an already issued asset. Must have reissuance
// token in wallet to do so.
func ReissueAsset(url string, params []string) (transactionResult types.ReissueAssetResult, err error) {
	result, err := SendRequest(url, types.MethodReissueAsset, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	return
}

// SendToAddress sends an amount to a given address.
func SendToAddress(url string, params []string) (hex string, err error) {
	result, err := SendRequest(url, types.MethodSendToAddress, params)
	if err != nil {
		return
	}
	hex = strings.ReplaceAll(string(result), "\"", "")
	return
}

// SignRawTransactionWithWallet signs inputs for raw transaction (serialized,
// hex-encoded).
func SignRawTransactionWithWallet(url string, params []string) (transactionResult types.SignRawTransactionWithWalletResult, err error) {
	result, err := SendRequest(url, types.MethodSignRawTransactionWithWallet, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	if !transactionResult.Complete {
		err = types.ErrMissingSignatures
		return
	}
	return
}

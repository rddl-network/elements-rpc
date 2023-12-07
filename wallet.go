package elementsrpc

import (
	"encoding/json"
	"strings"

	"github.com/rddl-network/elements-rpc/types"
)

// BlindRawTransaction converts one or more outputs of a raw transaction into
// confidential ones using only wallet inputs.
func BlindRawTransaction(url, params string) (hex string, err error) {
	result, err := SendRequest(url, "blindrawtransaction", params)
	if err != nil {
		return
	}
	hex = strings.ReplaceAll(string(result), "\"", "")
	return
}

// GetAddressInfo returns information about the given address.
func GetAddressInfo(url, params string) (transactionResult types.GetAddressInfoResult, err error) {
	result, err := SendRequest(url, "getaddressinfo", params)
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
func GetNewAddress(url, params string) (address string, err error) {
	result, err := SendRequest(url, "getnewaddress", params)
	if err != nil {
		return
	}
	address = strings.ReplaceAll(string(result), "\"", "")
	return
}

// GetTransaction retrieves a transaction from the chain.
func GetTransaction(url, params string) (transactionResult types.GetTransactionResult, err error) {
	result, err := SendRequest(url, "gettransaction", params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	return
}

// SignRawTransactionWithWallet signs inputs for raw transaction (serialized,
// hex-encoded).
func SignRawTransactionWithWallet(url, params string) (transactionResult types.SignRawTransactionWithWalletResult, err error) {
	result, err := SendRequest(url, "signrawtransactionwithwallet", params)
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

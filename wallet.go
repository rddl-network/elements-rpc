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
func GetAddressInfo(url string, params []string) (addressInfoResult types.GetAddressInfoResult, err error) {
	result, err := SendRequest(url, types.MethodGetAddressInfo, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &addressInfoResult)
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

// LoadWallet loads a wallet from a wallet file or directory.
func LoadWallet(url string, params []string) (loadWalletResult types.LoadWalletResult, err error) {
	result, err := SendRequest(url, types.MethodLoadWallet, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &loadWalletResult)
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

// UnloadWallet unloads the wallet referenced by the request endpoint otherwise
// unloads the wallet specified in the argument.
func UnloadWallet(url string, params []string) (unloadWalletResult types.UnloadWalletResult, err error) {
	result, err := SendRequest(url, types.MethodUnloadWallet, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &unloadWalletResult)
	if err != nil {
		return
	}
	return
}

// Walletpassphrase stores the wallet decryption key in memory for 'timeout'
// seconds.
func Walletpassphrase(url string, params []string) (walletpassphraseResult types.WalletpassphraseResult, err error) {
	resultBytes, err := SendRequest(url, types.MethodWalletpassphrase, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(resultBytes, &walletpassphraseResult)
	if err != nil {
		return
	}
	return
}

// Listwallets returns the list of loaded wallets.
func ListWallets(url string, params []string) (wallets []string, err error) {
	byteWallets, err := SendRequest(url, types.MethodListWallets, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(byteWallets, &wallets)
	if err != nil {
		return
	}
	return
}

// Listreceivedbyaddress returns a list of transactions
func ListReceivedByAddress(url string, params []string) (listReceivedByAddressResults []types.ListReceivedByAddressResult, err error) {
	result, err := SendRequest(url, types.MethodListReceivedByAddress, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &listReceivedByAddressResults)
	if err != nil {
		return
	}
	return
}

func GetBalance(url string, params []string) (getBalanceResult types.GetBalanceResult, err error) {
	result, err := SendRequest(url, types.MethodGetBalance, params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &getBalanceResult)
	if err != nil {
		return
	}
	return
}

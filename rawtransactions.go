package elementsrpc

import (
	"encoding/json"
	"strings"

	"github.com/rddl-network/elements-rpc/types"
)

// CreateRawTransaction creates a transaction spending the given inputs and
// creating new outputs.
func CreateRawTransaction(url, params string) (hex string, err error) {
	result, err := SendRequest(url, "createrawtransaction", params)
	if err != nil {
		return
	}
	hex = strings.ReplaceAll(string(result), "\"", "")
	return
}

// FundRawTransaction funds a raw transaction.
func FundRawTransaction(url, params string) (transactionResult types.FundRawTransactionResult, err error) {
	result, err := SendRequest(url, "fundrawtransaction", params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResult)
	if err != nil {
		return
	}
	return
}

// RawIssueAsset creates an asset by attaching issuances to transaction inputs.
func RawIssueAsset(url, params string) (transactionResults []types.RawIssueAssetResult, err error) {
	result, err := SendRequest(url, "rawissueasset", params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResults)
	if err != nil {
		return
	}
	return
}

// SendRawTransaction submits a raw transaction (serialized, hex-encoded) to
// local node and network.
func SendRawTransaction(url, params string) (hex string, err error) {
	result, err := SendRequest(url, "sendrawtransaction", params)
	if err != nil {
		return
	}
	hex = strings.ReplaceAll(string(result), "\"", "")
	return
}

// TestMempoolAccept returns result of mempool acceptance tests indicating if
// raw transaction(s) (serialized, hex-encoded) would be accepted by mempool.
func TestMempoolAccept(url, params string) (transactionResults []types.TestMempoolAcceptResult, err error) {
	result, err := SendRequest(url, "testmempoolaccept", params)
	if err != nil {
		return
	}
	err = json.Unmarshal(result, &transactionResults)
	if err != nil {
		return
	}
	return
}

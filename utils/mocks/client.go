package mocks

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rddl-network/elements-rpc/types"
)

// Body mocks the request sent to the elements' RPC
type Body struct {
	Jsonrpc string `json:"jsonrpc"`
	Method  string `json:"method"`
	Params  []any  `json:"params"`
}

// MockClient is the mock client
type MockClient struct {
	DoFunc func(req *http.Request) (*http.Response, error)
}

// Note: not all fields in the results are set
var (
	address                            = "tlq1qqdhf9taz5ftpxcl0qphdurs9csneacznlazmqqdjs9te0jh7nld0w3rmvzfuggr7l508ahclr69te88exl4rjp8z8etu5d35t"
	addressInfoResult                  = types.GetAddressInfoResult{Address: address, Confidential: address, Pubkey: "02" + zeros}
	fundRawTransactionResult           = types.FundRawTransactionResult{Hex: zeros, Fee: 0.0, Changepos: 0}
	rawIssueAssetResult                = types.RawIssueAssetResult{Hex: zeros, Vin: 0, Entropy: "", Asset: zeros, Token: ""}
	rawIssueAssetResults               = []types.RawIssueAssetResult{rawIssueAssetResult}
	reissueAssetResult                 = types.ReissueAssetResult{TxID: zeros, Vin: 0}
	signRawTransactionWithWalletResult = types.SignRawTransactionWithWalletResult{Hex: zeros, Complete: true}
	testMempoolAcceptResult            = types.TestMempoolAcceptResult{Allowed: true, Txid: zeros, Vsize: 0, Wtxid: ""}
	testMempoolAcceptResults           = []types.TestMempoolAcceptResult{testMempoolAcceptResult}
	transactionResult                  = types.GetTransactionResult{Hex: zeros, TxID: zeros, Amount: map[string]float64{
		"7add40beb27df701e02ee85089c5bc0021bc813823fedb5f1dcb5debda7f3da9": 10000,
	}}
	testDeriveAddressesResult = types.DeriveAddressesResult{address}
	zeros                     = "0000000000000000000000000000000000000000000000000000000000000000"
	testLoadWallet            = types.LoadWalletResult{Name: "testwallet4", Warning: ""}
	testUnloadWallet          = types.UnloadWalletResult{Warning: ""}
	testListWallets           = []string{"testwallet4"}
)

// GetDoFunc fetches the mock client's `Do` func
func GetDoFunc(req *http.Request) (*http.Response, error) {
	var body Body
	bodyBytes, err := io.ReadAll(req.Body)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(bodyBytes, &body)
	if err != nil {
		return nil, err
	}

	var response types.Response
	response.Error.Code = 0
	response.Error.Message = ""

	switch body.Method {
	case types.MethodBlindRawTransaction:
		response.Result = zeros
	case types.MethodCreateRawTransaction:
		response.Result = zeros
	case types.MethodFundRawTransaction:
		response.Result = fundRawTransactionResult
	case types.MethodGetAddressInfo:
		response.Result = addressInfoResult
	case types.MethodGetNewAddress:
		response.Result = address
	case types.MethodGetTransaction:
		response.Result = transactionResult
	case types.MethodRawIssueAsset:
		response.Result = rawIssueAssetResults
	case types.MethodReissueAsset:
		response.Result = reissueAssetResult
	case types.MethodSendRawTransaction:
		response.Result = zeros
	case types.MethodSendToAddress:
		response.Result = zeros
	case types.MethodSignRawTransactionWithWallet:
		response.Result = signRawTransactionWithWalletResult
	case types.MethodTestMempoolAccept:
		response.Result = testMempoolAcceptResults
	case types.MethodDeriveAddresses:
		response.Result = testDeriveAddressesResult
	case types.MethodLoadWallet:
		response.Result = testLoadWallet
	case types.MethodUnloadWallet:
		response.Result = testUnloadWallet
	case types.MethodListWallets:
		response.Result = testListWallets
	default:
		response.Result = nil
		response.Error.Code = -1337
		response.Error.Message = fmt.Sprintf("method '%s' not implemented", body.Method)
	}
	respBytes, err := json.Marshal(&response)
	if err != nil {
		return nil, err
	}
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewReader(respBytes)),
	}
	return resp, nil
}

// Do is the mock client's `Do` func
func (m *MockClient) Do(req *http.Request) (*http.Response, error) {
	return GetDoFunc(req)
}

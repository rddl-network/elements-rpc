package mocks

import (
	"bytes"
	"encoding/json"
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
	addressInfoResult                  = types.GetAddressInfoResult{Address: address, Confidential: address}
	fundRawTransactionResult           = types.FundRawTransactionResult{Hex: zeros, Fee: 0.0, Changepos: 0}
	rawIssueAssetResult                = types.RawIssueAssetResult{Hex: zeros, Vin: 0, Entropy: "", Asset: "", Token: ""}
	rawIssueAssetResults               = []types.RawIssueAssetResult{rawIssueAssetResult}
	reissueAssetResult                 = types.ReissueAssetResult{TxID: zeros, Vin: 0}
	signRawTransactionWithWalletResult = types.SignRawTransactionWithWalletResult{Hex: zeros, Complete: true}
	testMempoolAcceptResult            = types.TestMempoolAcceptResult{Allowed: true, Txid: zeros, Vsize: 0, Wtxid: ""}
	testMempoolAcceptResults           = []types.TestMempoolAcceptResult{testMempoolAcceptResult}
	transactionResult                  = types.GetTransactionResult{Hex: zeros, TxID: zeros}
	zeros                              = "0000000000000000000000000000000000000000000000000000000000000000"
	zerosWithQuotes                    = `"` + zeros + `"`
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
		response.Result = zerosWithQuotes
	case types.MethodCreateRawTransaction:
		response.Result = zerosWithQuotes
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
		response.Result = zerosWithQuotes
	case types.MethodSendToAddress:
		response.Result = zerosWithQuotes
	case types.MethodSignRawTransactionWithWallet:
		response.Result = signRawTransactionWithWalletResult
	case types.MethodTestMempoolAccept:
		response.Result = testMempoolAcceptResults
	default:
		response.Result = nil
		response.Error.Code = -1337
		response.Error.Message = "method not implemented"
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

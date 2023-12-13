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

var (
	txID = "0000000000000000000000000000000000000000000000000000000000000000"
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
	switch body.Method {
	case "reissueasset":
		reissueAssetResult := types.ReissueAssetResult{TxID: txID, Vin: 0}
		response.Result = reissueAssetResult
		response.Error.Code = 0
		response.Error.Message = ""
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

package types

// FundRawTransactionResult models the result of a "fundrawtransaction"
// transaction.
type FundRawTransactionResult struct {
	Hex       string  `json:"hex"`
	Fee       float64 `json:"fee"`
	Changepos int64   `json:"changepos"`
}

// RawIssueAssetResult models the result of a "rawissueasset" transaction.
type RawIssueAssetResult struct {
	Hex     string `json:"hex"`
	Vin     int64  `json:"vin"`
	Entropy string `json:"entropy"`
	Asset   string `json:"asset"`
	Token   string `json:"token"`
}

// TestMempoolAcceptResult models the result of a "testmempoolaccept" request.
type TestMempoolAcceptResult struct {
	Allowed bool `json:"allowed"`
	Fees    struct {
		Base float64 `json:"base"`
	} `json:"fees"`
	Txid  string `json:"txid"`
	Vsize int    `json:"vsize"`
	Wtxid string `json:"wtxid"`
}

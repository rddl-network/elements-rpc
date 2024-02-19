package types

// FundRawTransactionResult models the result of a "fundrawtransaction"
// transaction.
type FundRawTransactionResult struct {
	Hex       string  `json:"hex"`
	Fee       float64 `json:"fee"`
	Changepos int64   `json:"changepos"`
}

// GetRawTransactionResult models the result of a "getrawtransaction"
// transaction.
type GetRawTransactionResult struct {
	Txid     string `json:"txid"`
	Hash     string `json:"hash"`
	Wtxid    string `json:"wtxid"`
	Withash  string `json:"withash"`
	Version  int    `json:"version"`
	Size     int    `json:"size"`
	Vsize    int    `json:"vsize"`
	Weight   int    `json:"weight"`
	Locktime int    `json:"locktime"`
	Vin      []struct {
		Txid      string `json:"txid"`
		Vout      int    `json:"vout"`
		ScriptSig struct {
			Asm string `json:"asm"`
			Hex string `json:"hex"`
		} `json:"scriptSig"`
		IsPegin     bool     `json:"is_pegin"`
		Sequence    int64    `json:"sequence"`
		Txinwitness []string `json:"txinwitness"`
	} `json:"vin"`
	Vout []struct {
		ValueMinimum              float64 `json:"value-minimum,omitempty"`
		ValueMaximum              float64 `json:"value-maximum,omitempty"`
		CtExponent                int     `json:"ct-exponent,omitempty"`
		CtBits                    int     `json:"ct-bits,omitempty"`
		Surjectionproof           string  `json:"surjectionproof,omitempty"`
		Valuecommitment           string  `json:"valuecommitment,omitempty"`
		Assetcommitment           string  `json:"assetcommitment,omitempty"`
		Commitmentnonce           string  `json:"commitmentnonce"`
		CommitmentnonceFullyValid bool    `json:"commitmentnonce_fully_valid"`
		N                         int     `json:"n"`
		ScriptPubKey              struct {
			Asm     string `json:"asm"`
			Hex     string `json:"hex"`
			Address string `json:"address"`
			Type    string `json:"type"`
		} `json:"scriptPubKey,omitempty"`
	} `json:"vout"`
	Hex           string `json:"hex"`
	Blockhash     string `json:"blockhash"`
	Confirmations int    `json:"confirmations"`
	Time          int    `json:"time"`
	Blocktime     int    `json:"blocktime"`
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

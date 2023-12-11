package types

// GetAddressInfoResult models the result of a "getaddressinfo" request.
type GetAddressInfoResult struct {
	Address             string   `json:"address"`
	Confidential        string   `json:"confidential"`
	ConfidentialKey     string   `json:"confidential_key"`
	Desc                string   `json:"desc"`
	Hdkeypath           string   `json:"hdkeypath"`
	Hdmasterfingerprint string   `json:"hdmasterfingerprint"`
	Hdseedid            string   `json:"hdseedid"`
	Ischange            bool     `json:"ischange"`
	Ismine              bool     `json:"ismine"`
	Isscript            bool     `json:"isscript"`
	Iswatchonly         bool     `json:"iswatchonly"`
	Iswitness           bool     `json:"iswitness"`
	Labels              []string `json:"labels"`
	Pubkey              string   `json:"pubkey"`
	ScriptPubKey        string   `json:"scriptPubKey"`
	Solvable            bool     `json:"solvable"`
	Timestamp           int      `json:"timestamp"`
	Unconfidential      string   `json:"unconfidential"`
	WitnessProgram      string   `json:"witness_program"`
	WitnessVersion      int      `json:"witness_version"`
}

// This models the "short" version of the ListTransactionsResult type, which
// excludes fields common to the transaction.  These common fields are instead
// part of the GetTransactionResult.
type GetTransactionDetailsResult struct {
	Account           string   `json:"account"`
	Address           string   `json:"address,omitempty"`
	Amount            float64  `json:"amount"`
	Category          string   `json:"category"`
	InvolvesWatchOnly bool     `json:"involveswatchonly,omitempty"`
	Fee               *float64 `json:"fee,omitempty"`
	Vout              uint32   `json:"vout"`
}

// GetTransactionResult models the data from the gettransaction command.
type GetTransactionResult struct {
	Amount          map[string]float64            `json:"amount"`
	Fee             map[string]float64            `json:"fee,omitempty"`
	Confirmations   int64                         `json:"confirmations"`
	BlockHash       string                        `json:"blockhash"`
	BlockIndex      int64                         `json:"blockindex"`
	BlockTime       int64                         `json:"blocktime"`
	TxID            string                        `json:"txid"`
	WalletConflicts []string                      `json:"walletconflicts"`
	Time            int64                         `json:"time"`
	TimeReceived    int64                         `json:"timereceived"`
	Details         []GetTransactionDetailsResult `json:"details"`
	Hex             string                        `json:"hex"`
}

// ReissueAssetResult models the result of a "reissueasset" transaction.
type ReissueAssetResult struct {
	TxID string `json:"txid"`
	Vin  int64  `json:"vin"`
}

// SignRawTransactionWithWalletResult models the result of a
// "signrawtransactionwithwallet" transaction.
type SignRawTransactionWithWalletResult struct {
	Hex      string `json:"hex"`
	Complete bool   `json:"complete"`
}

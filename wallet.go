package elementsrpc

import (
	"encoding/json"

	"github.com/rddl-network/elements-rpc/types"
)

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

// GetWalletTx retrieves a transaction from the chain.
//
// Deprecated: Only for backward compatibility. Use GetTransaction instead.
func GetWalletTx(url string, txhash string) (tx types.GetTransactionResult, err error) {
	tx, err = GetTransaction(url, `"`+txhash+`"`)
	return
}

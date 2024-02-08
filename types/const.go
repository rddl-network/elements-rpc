package types

import (
	"errors"
)

var (
	ErrMissingSignatures = errors.New("transaction does not have the complete set of signatures")

	// elements RPC methods
	MethodBlindRawTransaction          = "blindrawtransaction"
	MethodCreateRawTransaction         = "createrawtransaction"
	MethodFundRawTransaction           = "fundrawtransaction"
	MethodGetAddressInfo               = "getaddressinfo"
	MethodGetNewAddress                = "getnewaddress"
	MethodGetTransaction               = "gettransaction"
	MethodRawIssueAsset                = "rawissueasset"
	MethodReissueAsset                 = "reissueasset"
	MethodSendRawTransaction           = "sendrawtransaction"
	MethodSendToAddress                = "sendtoaddress"
	MethodSignRawTransactionWithWallet = "signrawtransactionwithwallet"
	MethodTestMempoolAccept            = "testmempoolaccept"
	MethodDeriveAddresses              = "deriveaddresses"
)

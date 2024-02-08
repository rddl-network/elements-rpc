package elementsrpc

import (
	"github.com/rddl-network/elements-rpc/types"
	"github.com/rddl-network/elements-rpc/utils"
)

func DeriveAddresses(url string, params []string) (addresses types.DeriveAddressesResult, err error) {
	result, err := SendRequest(url, types.MethodDeriveAddresses, params)
	if err != nil {
		return
	}

	addresses = utils.ParseResultToArray(result)
	return
}

package elementsrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	types "github.com/rddl-network/elements-rpc/types"
)

func GetWalletTx(url string, txhash string) (tx types.GetTransactionResult, err error) {
	jsonStr := []byte(fmt.Sprintf(`{"jsonrpc":"1.0","method":"gettransaction","params":["%s"]}`, txhash))

	request, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}

	response, err := client.Do(request)
	if err != nil {
		return tx, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return tx, err
	}

	var res types.Result
	if err := json.Unmarshal(body, &res); err != nil {
		return tx, err
	}

	tx = res.Result

	return
}

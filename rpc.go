package elementsrpc

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/rddl-network/elements-rpc/types"
)

func SendRequest(url, method, params string) (result []byte, err error) {
	jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","method":"%s","params":[%s]}`, method, params)

	request, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer([]byte(jsonStr)))
	request.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(request)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	var response types.Response
	if err = json.Unmarshal(body, &response); err != nil {
		return
	}

	if response.Error.Code != 0 {
		err = fmt.Errorf("%s: %d", response.Error.Message, response.Error.Code)
		return
	}

	result, err = json.Marshal(response.Result)
	if err != nil {
		return
	}
	return
}

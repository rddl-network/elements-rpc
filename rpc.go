package elementsrpc

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/rddl-network/elements-rpc/types"
)

// HTTPClient interface
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

var (
	Client HTTPClient
)

func init() {
	Client = &http.Client{}
}

func parse(params []string) (param string) {
	if len(params) == 0 {
		param = ""
		return
	}

	param = strings.Join(params, ",")
	return
}

func SendRequest(url, method string, params []string) (result []byte, err error) {
	param := parse(params)
	jsonStr := fmt.Sprintf(`{"jsonrpc":"1.0","method":"%s","params":[%s]}`, method, param)

	ctx := context.Background()
	request, err := http.NewRequestWithContext(ctx, http.MethodPost, url, bytes.NewBufferString(jsonStr))
	if err != nil {
		return
	}
	request.Header.Set("Content-Type", "application/json")

	resp, err := Client.Do(request)
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

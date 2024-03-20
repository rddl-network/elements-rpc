package elementsrpc_test

import (
	"testing"

	elementsrpc "github.com/rddl-network/elements-rpc"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	t.Parallel()
	for _, tc := range []struct {
		input          []string
		expectedResult string
	}{
		{
			[]string{"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d"},
			"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d",
		},
		{
			[]string{"\"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d\""},
			"\"1075db55d416d3ca199f55b6084e2115b9345e16c5cf302fc80e9d5fbf5d48d\"",
		},
		{
			[]string{
				"10",
				"false",
				"true",
				"tex1qr4p5fwsttwsfqvk832gx2pqq44d9pdccwcfjtg",
				"06c20c8de513527f1ae6c901f74a05126525ac2d7e89306f4a7fd5ec4e674403",
			},
			"10,false,true,tex1qr4p5fwsttwsfqvk832gx2pqq44d9pdccwcfjtg,06c20c8de513527f1ae6c901f74a05126525ac2d7e89306f4a7fd5ec4e674403",
		},
		{
			[]string{`[]`, `[{"data":"00"}]`},
			"[],[{\"data\":\"00\"}]",
		},
		{
			[]string{`"abc"`, `[{"data":"00"}]`},
			"\"abc\",[{\"data\":\"00\"}]",
		},
	} {
		result := elementsrpc.Parse(tc.input)
		assert.Equal(t, tc.expectedResult, result)
	}
}

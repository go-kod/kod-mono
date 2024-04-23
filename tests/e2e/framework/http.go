package framework

import (
	"context"
	"encoding/json"
	"time"

	"github.com/expr-lang/expr"
	"github.com/go-resty/resty/v2"
	"github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/require"
)

type HTTPTestCase struct {
	BaseURL string
	Method  string
	Path    string
	Body    string
	Header  map[string]string
	Query   string
	Timeout time.Duration

	Expect string
}

// RunHTTPTestCase runs a test case against the given handler.
func RunHTTPTestCase(tc HTTPTestCase) {
	t := ginkgo.GinkgoT()

	if tc.Timeout == 0 {
		tc.Timeout = time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), tc.Timeout)
	defer cancel()

	req := resty.New().SetBaseURL(tc.BaseURL).R()
	req.SetQueryString(tc.Query)
	req.SetBody(tc.Body)
	req.SetHeaders(tc.Header)
	req.SetContext(ctx)

	res, err := req.Execute(tc.Method, tc.Path)

	require.Nil(t, err, "error: %s", err)

	if tc.Expect != "" {

		var resMap map[string]interface{}
		err = json.Unmarshal([]byte(res.String()), &resMap)
		require.Nil(t, err)

		data := map[string]interface{}{
			"status": res.StatusCode(),
			"header": res.Header(),
			"body":   resMap,
		}

		program, err := expr.Compile(tc.Expect, expr.Env(data))
		require.Nil(t, err, "error: %s", err)

		val, err := expr.Run(program, data)
		require.Nil(t, err, "data", data)
		require.True(t, val.(bool), "data", data)
	}
}

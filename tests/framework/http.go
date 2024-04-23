package framework

import (
	"context"
	"encoding/json"
	"time"

	"github.com/go-resty/resty/v2"
	"github.com/google/cel-go/cel"
	"github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/require"
)

type HTTPTestCase struct {
	BaseURL   string
	Method    string
	Path      string
	Body      string
	Timeout   time.Duration
	Header    map[string]string
	Query     string
	ExpectCEL string
}

// RunHTTPTestCase runs a test case against the given handler.
func RunHTTPTestCase(htc HTTPTestCase) {
	ginkgoT := ginkgo.GinkgoT()

	if htc.Timeout == 0 {
		htc.Timeout = time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), htc.Timeout)
	defer cancel()

	req := resty.New().SetBaseURL(htc.BaseURL).R()
	req.SetQueryString(htc.Query)
	req.SetBody(htc.Body)
	req.SetHeaders(htc.Header)
	req.SetContext(ctx)

	res, err := req.Execute(htc.Method, htc.Path)

	require.Nil(ginkgoT, err, "error: %s", err)

	if htc.ExpectCEL != "" {
		env, err := cel.NewEnv(
			cel.Variable("status", cel.IntType),
			cel.Variable("header", cel.MapType(cel.StringType, cel.ListType(cel.StringType))),
			cel.Variable("response", cel.MapType(cel.StringType, cel.AnyType)),
		)
		require.Nil(ginkgoT, err)

		ast, issues := env.Compile(htc.ExpectCEL)

		require.False(ginkgoT, issues != nil && issues.Err() != nil, "type-check error: %s", issues.Err())

		prg, err := env.Program(ast)
		require.Nil(ginkgoT, err)

		var resMap map[string]interface{}
		err = json.Unmarshal([]byte(res.String()), &resMap)
		require.Nil(ginkgoT, err)

		data := map[string]interface{}{
			"status":   res.StatusCode(),
			"header":   res.Header(),
			"response": resMap,
		}
		val, _, err := prg.Eval(data)
		require.Nil(ginkgoT, err, "data", data)
		require.True(ginkgoT, val.Value().(bool), "data", data)
	}
}

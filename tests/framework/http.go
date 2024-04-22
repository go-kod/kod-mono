package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/go-kod/kod/ext/client/kresty"
	"github.com/onsi/ginkgo/v2"
	"github.com/stretchr/testify/assert"
)

type HTTPTestCase struct {
	Conf               *kresty.Config
	Method             string
	Path               string
	Body               string
	Timeout            time.Duration
	Header             map[string]string
	Query              string
	ExpectHeader       http.Header
	ExpectStatus       int
	ExpectBody         string
	ExpectContainsBody string
}

// RunHTTPTestCase runs a test case against the given handler.
func RunHTTPTestCase(htc HTTPTestCase) {
	ginkgoT := ginkgo.GinkgoT()

	if htc.Timeout == 0 {
		htc.Timeout = time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), htc.Timeout)
	defer cancel()

	req := htc.Conf.Build().R()
	req.SetQueryString(htc.Query)
	req.SetBody(htc.Body)
	req.SetHeaders(htc.Header)
	req.SetContext(ctx)

	res, err := req.Execute(htc.Method, htc.Path)

	assert.Nil(ginkgoT, err, "error: %s", err)

	if htc.ExpectStatus > 0 {
		assert.Equal(ginkgoT, htc.ExpectStatus, res.StatusCode(),
			"expected: %d\nactually: %d", htc.ExpectStatus, res.StatusCode())
	}

	if len(htc.ExpectHeader) > 0 {
		assert.EqualValues(ginkgoT, htc.ExpectHeader, res.Header(),
			"expected: %s\nactually: %s", htc.ExpectHeader, res.Header())
	}

	if len(htc.ExpectContainsBody) > 0 {
		assert.Contains(ginkgoT, res.String(), htc.ExpectContainsBody,
			"expected: %s\nactually: %s", htc.ExpectContainsBody, res.String())
	}

	if len(htc.ExpectBody) > 0 {
		var body bytes.Buffer
		err = json.Compact(&body, []byte(res.String()))
		// 如果Compact失败，则说明不是json格式
		if err != nil {
			body.WriteString(res.String())
		}

		assert.Equal(ginkgoT, htc.ExpectBody, body.String(),
			"expected: %s\nactually: %s", htc.ExpectBody, body.String())
	}
}

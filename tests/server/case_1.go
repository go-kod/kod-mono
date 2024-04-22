package server

import (
	"net/http"

	"github.com/go-kod/kod-mono/tests/framework"
	"github.com/go-kod/kod/ext/client/kresty"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("server", func() {
	DescribeTable("uniqueId", func(htc framework.HTTPTestCase) {
		framework.RunHTTPTestCase(htc)
	},
		Entry("SayHello", framework.HTTPTestCase{
			Conf: &kresty.Config{
				BaseURL: "http://localhost:9527",
			},
			Method:             http.MethodGet,
			Path:               "/uniqueId",
			Header:             map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			Query:              "name=bob",
			ExpectContainsBody: `{"Uuid":"`,
			ExpectStatus:       http.StatusOK,
		}),
	)
})

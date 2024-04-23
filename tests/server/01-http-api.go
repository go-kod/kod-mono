package server

import (
	"net/http"

	"github.com/go-kod/kod-mono/tests/framework"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("01-http-api", func() {
	DescribeTable("uniqueId", func(htc framework.HTTPTestCase) {
		framework.RunHTTPTestCase(htc)
	},
		Entry("ok", framework.HTTPTestCase{
			BaseURL:   "http://localhost:9527",
			Method:    http.MethodGet,
			Path:      "/uniqueId",
			Header:    map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			Query:     "name=bob",
			ExpectCEL: "status==200 && response.Uuid.size()==19",
		}),
		Entry("invalid argument", framework.HTTPTestCase{
			BaseURL:   "http://localhost:9527",
			Method:    http.MethodGet,
			Path:      "/uniqueId",
			Header:    map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			Query:     "",
			ExpectCEL: `status==200 && response.code==3 && response.message=="name is empty" && header["Content-Type"]==["application/json; charset=utf-8"]`,
		}),
	)
})

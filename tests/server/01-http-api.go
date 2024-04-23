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
			BaseURL: "http://localhost:9527",
			Method:  http.MethodGet,
			Path:    "/uniqueId",
			Header:  map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			Query:   "name=bob",
			Expect:  "status==200 && len(body.Uuid)==19",
		}),
		Entry("invalid argument", framework.HTTPTestCase{
			BaseURL: "http://localhost:9527",
			Method:  http.MethodGet,
			Path:    "/uniqueId",
			Header:  map[string]string{"Content-Type": "application/x-www-form-urlencoded"},
			Query:   "",
			Expect:  `status==200 && body.code==3 && body.message=="name is empty" && header["Content-Type"]==["application/json; charset=utf-8"]`,
		}),
	)
})

package server

import (
	snowflakev1 "github.com/go-kod/kod-mono/api/gen/go/snowflake/v1"
	"github.com/go-kod/kod-mono/tests/e2e/framework"
	. "github.com/onsi/ginkgo/v2"
)

var _ = Describe("01-grpc-api", func() {
	DescribeTable("uniqueId", func(tc framework.GRPCTestCase) {
		framework.RunGRPCTestCase(tc)
	},
		Entry("ok", framework.GRPCTestCase{
			Addr:   "passthrough:///localhost:9528",
			Method: snowflakev1.SnowflakeService_UniqueId_FullMethodName,
			Body:   `{"name":"bob"}`,
			Expect: `int(status.Code())==0 && metadata["content-type"] == ["application/grpc"] && len(body.uuid)==19`,
		}),
	)
})

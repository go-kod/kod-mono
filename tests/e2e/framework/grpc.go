package framework

import (
	"bytes"
	"context"
	"encoding/json"
	"strings"
	"time"

	"github.com/expr-lang/expr"
	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/desc"
	"github.com/jhump/protoreflect/grpcreflect"
	"github.com/onsi/ginkgo/v2"
	"github.com/samber/lo"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/runtime/protoiface"
)

type GRPCTestCase struct {
	Addr    string
	Method  string
	Header  map[string]string
	Body    string
	Timeout time.Duration

	Expect string
}

// RunGRPCTestCase runs a test case against the given handler.
func RunGRPCTestCase(tc GRPCTestCase) {
	t := ginkgo.GinkgoT()

	if tc.Timeout == 0 {
		tc.Timeout = time.Second
	}

	ctx, cancel := context.WithTimeout(context.Background(), tc.Timeout)
	defer cancel()

	cc, err := grpc.NewClient(tc.Addr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	assert.Nil(t, err)
	defer cc.Close()

	refClient := grpcreflect.NewClientV1Alpha(ctx, reflectpb.NewServerReflectionClient(cc))
	descSource := grpcurl.DescriptorSourceFromServer(ctx, refClient)

	rf, _, err := grpcurl.RequestParserAndFormatter(grpcurl.FormatJSON, descSource, bytes.NewBufferString(tc.Body), grpcurl.FormatOptions{
		EmitJSONDefaultFields: true,
	})
	require.Nil(t, err)

	h := &eventHandler{}

	headers := lo.MapToSlice(tc.Header, func(key string, value string) string {
		return strings.Join([]string{key, value}, ":")
	})

	err = grpcurl.InvokeRPC(ctx, descSource, cc, strings.TrimPrefix(tc.Method, "/"), headers, h, rf.Next)
	require.Nil(t, err)

	if tc.Expect != "" {

		data := map[string]interface{}{
			"status":   h.Status,
			"metadata": h.Metadata,
			"body":     h.Response,
		}

		program, err := expr.Compile(tc.Expect, expr.Env(data))
		require.Nil(t, err, "error: %s, data:%+v", err, data)

		val, err := expr.Run(program, data)

		require.Nil(t, err, "error: %s, data:%+v", err, data)
		require.True(t, val.(bool), "data:%+v, val:%+v", data, val)
	}
}

type eventHandler struct {
	Status   *status.Status
	Response map[string]interface{}
	Metadata metadata.MD
}

func (h *eventHandler) OnResolveMethod(md *desc.MethodDescriptor) {
}

func (h *eventHandler) OnSendHeaders(md metadata.MD) {
}

func (h *eventHandler) OnReceiveHeaders(md metadata.MD) {
	h.Metadata = md
}

func (h *eventHandler) OnReceiveResponse(resp protoiface.MessageV1) {
	reply := make(map[string]interface{})
	lo.Must0(json.Unmarshal(lo.Must(json.Marshal(resp)), &reply))
	h.Response = reply
}

func (h *eventHandler) OnReceiveTrailers(stat *status.Status, md metadata.MD) {
	h.Status = stat
}

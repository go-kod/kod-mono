// Code generated by kod struct2interface; DO NOT EDIT.

package example

import (
	"context"
)

// component is a component that implements Service.
type Service interface {
	UniqueID(ctx context.Context, req *TestReq) (*TestRes, error)
}

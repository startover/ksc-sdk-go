// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package knadiface provides an interface to enable mocking the knad service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package knadiface

import (
	"github.com/KscSDK/ksc-sdk-go/service/knad"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
)

// KnadAPI provides an interface to enable mocking the
// knad.Knad service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// knad.
//	func myFunc(svc knadiface.KnadAPI) bool {
//	    // Make svc.CreateKnad request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := knad.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockKnadClient struct {
//	    knadiface.KnadAPI
//	}
//	func (m *mockKnadClient) CreateKnad(input *map[string]interface{}) (*map[string]interface{}, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockKnadClient{}
//
//	    myfunc(mockSvc)
//
//	    // Verify myFunc's functionality
//	}
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type KnadAPI interface {
	CreateKnad(*map[string]interface{}) (*map[string]interface{}, error)
	CreateKnadWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	CreateKnadRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})

	DescribeKnad(*map[string]interface{}) (*map[string]interface{}, error)
	DescribeKnadWithContext(aws.Context, *map[string]interface{}, ...request.Option) (*map[string]interface{}, error)
	DescribeKnadRequest(*map[string]interface{}) (*request.Request, *map[string]interface{})
}

var _ KnadAPI = (*knad.Knad)(nil)

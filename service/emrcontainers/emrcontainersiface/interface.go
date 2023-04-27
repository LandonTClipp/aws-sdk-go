// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package emrcontainersiface provides an interface to enable mocking the Amazon EMR Containers service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package emrcontainersiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/emrcontainers"
)

// EMRContainersAPI provides an interface to enable mocking the
// emrcontainers.EMRContainers service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon EMR Containers.
//	func myFunc(svc emrcontainersiface.EMRContainersAPI) bool {
//	    // Make svc.CancelJobRun request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := emrcontainers.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockEMRContainersClient struct {
//	    emrcontainersiface.EMRContainersAPI
//	}
//	func (m *mockEMRContainersClient) CancelJobRun(input *emrcontainers.CancelJobRunInput) (*emrcontainers.CancelJobRunOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockEMRContainersClient{}
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
type EMRContainersAPI interface {
	CancelJobRun(*emrcontainers.CancelJobRunInput) (*emrcontainers.CancelJobRunOutput, error)
	CancelJobRunWithContext(aws.Context, *emrcontainers.CancelJobRunInput, ...request.Option) (*emrcontainers.CancelJobRunOutput, error)
	CancelJobRunRequest(*emrcontainers.CancelJobRunInput) (*request.Request, *emrcontainers.CancelJobRunOutput)

	CreateJobTemplate(*emrcontainers.CreateJobTemplateInput) (*emrcontainers.CreateJobTemplateOutput, error)
	CreateJobTemplateWithContext(aws.Context, *emrcontainers.CreateJobTemplateInput, ...request.Option) (*emrcontainers.CreateJobTemplateOutput, error)
	CreateJobTemplateRequest(*emrcontainers.CreateJobTemplateInput) (*request.Request, *emrcontainers.CreateJobTemplateOutput)

	CreateManagedEndpoint(*emrcontainers.CreateManagedEndpointInput) (*emrcontainers.CreateManagedEndpointOutput, error)
	CreateManagedEndpointWithContext(aws.Context, *emrcontainers.CreateManagedEndpointInput, ...request.Option) (*emrcontainers.CreateManagedEndpointOutput, error)
	CreateManagedEndpointRequest(*emrcontainers.CreateManagedEndpointInput) (*request.Request, *emrcontainers.CreateManagedEndpointOutput)

	CreateVirtualCluster(*emrcontainers.CreateVirtualClusterInput) (*emrcontainers.CreateVirtualClusterOutput, error)
	CreateVirtualClusterWithContext(aws.Context, *emrcontainers.CreateVirtualClusterInput, ...request.Option) (*emrcontainers.CreateVirtualClusterOutput, error)
	CreateVirtualClusterRequest(*emrcontainers.CreateVirtualClusterInput) (*request.Request, *emrcontainers.CreateVirtualClusterOutput)

	DeleteJobTemplate(*emrcontainers.DeleteJobTemplateInput) (*emrcontainers.DeleteJobTemplateOutput, error)
	DeleteJobTemplateWithContext(aws.Context, *emrcontainers.DeleteJobTemplateInput, ...request.Option) (*emrcontainers.DeleteJobTemplateOutput, error)
	DeleteJobTemplateRequest(*emrcontainers.DeleteJobTemplateInput) (*request.Request, *emrcontainers.DeleteJobTemplateOutput)

	DeleteManagedEndpoint(*emrcontainers.DeleteManagedEndpointInput) (*emrcontainers.DeleteManagedEndpointOutput, error)
	DeleteManagedEndpointWithContext(aws.Context, *emrcontainers.DeleteManagedEndpointInput, ...request.Option) (*emrcontainers.DeleteManagedEndpointOutput, error)
	DeleteManagedEndpointRequest(*emrcontainers.DeleteManagedEndpointInput) (*request.Request, *emrcontainers.DeleteManagedEndpointOutput)

	DeleteVirtualCluster(*emrcontainers.DeleteVirtualClusterInput) (*emrcontainers.DeleteVirtualClusterOutput, error)
	DeleteVirtualClusterWithContext(aws.Context, *emrcontainers.DeleteVirtualClusterInput, ...request.Option) (*emrcontainers.DeleteVirtualClusterOutput, error)
	DeleteVirtualClusterRequest(*emrcontainers.DeleteVirtualClusterInput) (*request.Request, *emrcontainers.DeleteVirtualClusterOutput)

	DescribeJobRun(*emrcontainers.DescribeJobRunInput) (*emrcontainers.DescribeJobRunOutput, error)
	DescribeJobRunWithContext(aws.Context, *emrcontainers.DescribeJobRunInput, ...request.Option) (*emrcontainers.DescribeJobRunOutput, error)
	DescribeJobRunRequest(*emrcontainers.DescribeJobRunInput) (*request.Request, *emrcontainers.DescribeJobRunOutput)

	DescribeJobTemplate(*emrcontainers.DescribeJobTemplateInput) (*emrcontainers.DescribeJobTemplateOutput, error)
	DescribeJobTemplateWithContext(aws.Context, *emrcontainers.DescribeJobTemplateInput, ...request.Option) (*emrcontainers.DescribeJobTemplateOutput, error)
	DescribeJobTemplateRequest(*emrcontainers.DescribeJobTemplateInput) (*request.Request, *emrcontainers.DescribeJobTemplateOutput)

	DescribeManagedEndpoint(*emrcontainers.DescribeManagedEndpointInput) (*emrcontainers.DescribeManagedEndpointOutput, error)
	DescribeManagedEndpointWithContext(aws.Context, *emrcontainers.DescribeManagedEndpointInput, ...request.Option) (*emrcontainers.DescribeManagedEndpointOutput, error)
	DescribeManagedEndpointRequest(*emrcontainers.DescribeManagedEndpointInput) (*request.Request, *emrcontainers.DescribeManagedEndpointOutput)

	DescribeVirtualCluster(*emrcontainers.DescribeVirtualClusterInput) (*emrcontainers.DescribeVirtualClusterOutput, error)
	DescribeVirtualClusterWithContext(aws.Context, *emrcontainers.DescribeVirtualClusterInput, ...request.Option) (*emrcontainers.DescribeVirtualClusterOutput, error)
	DescribeVirtualClusterRequest(*emrcontainers.DescribeVirtualClusterInput) (*request.Request, *emrcontainers.DescribeVirtualClusterOutput)

	GetManagedEndpointSessionCredentials(*emrcontainers.GetManagedEndpointSessionCredentialsInput) (*emrcontainers.GetManagedEndpointSessionCredentialsOutput, error)
	GetManagedEndpointSessionCredentialsWithContext(aws.Context, *emrcontainers.GetManagedEndpointSessionCredentialsInput, ...request.Option) (*emrcontainers.GetManagedEndpointSessionCredentialsOutput, error)
	GetManagedEndpointSessionCredentialsRequest(*emrcontainers.GetManagedEndpointSessionCredentialsInput) (*request.Request, *emrcontainers.GetManagedEndpointSessionCredentialsOutput)

	ListJobRuns(*emrcontainers.ListJobRunsInput) (*emrcontainers.ListJobRunsOutput, error)
	ListJobRunsWithContext(aws.Context, *emrcontainers.ListJobRunsInput, ...request.Option) (*emrcontainers.ListJobRunsOutput, error)
	ListJobRunsRequest(*emrcontainers.ListJobRunsInput) (*request.Request, *emrcontainers.ListJobRunsOutput)

	ListJobRunsPages(*emrcontainers.ListJobRunsInput, func(*emrcontainers.ListJobRunsOutput, bool) bool) error
	ListJobRunsPagesWithContext(aws.Context, *emrcontainers.ListJobRunsInput, func(*emrcontainers.ListJobRunsOutput, bool) bool, ...request.Option) error

	ListJobTemplates(*emrcontainers.ListJobTemplatesInput) (*emrcontainers.ListJobTemplatesOutput, error)
	ListJobTemplatesWithContext(aws.Context, *emrcontainers.ListJobTemplatesInput, ...request.Option) (*emrcontainers.ListJobTemplatesOutput, error)
	ListJobTemplatesRequest(*emrcontainers.ListJobTemplatesInput) (*request.Request, *emrcontainers.ListJobTemplatesOutput)

	ListJobTemplatesPages(*emrcontainers.ListJobTemplatesInput, func(*emrcontainers.ListJobTemplatesOutput, bool) bool) error
	ListJobTemplatesPagesWithContext(aws.Context, *emrcontainers.ListJobTemplatesInput, func(*emrcontainers.ListJobTemplatesOutput, bool) bool, ...request.Option) error

	ListManagedEndpoints(*emrcontainers.ListManagedEndpointsInput) (*emrcontainers.ListManagedEndpointsOutput, error)
	ListManagedEndpointsWithContext(aws.Context, *emrcontainers.ListManagedEndpointsInput, ...request.Option) (*emrcontainers.ListManagedEndpointsOutput, error)
	ListManagedEndpointsRequest(*emrcontainers.ListManagedEndpointsInput) (*request.Request, *emrcontainers.ListManagedEndpointsOutput)

	ListManagedEndpointsPages(*emrcontainers.ListManagedEndpointsInput, func(*emrcontainers.ListManagedEndpointsOutput, bool) bool) error
	ListManagedEndpointsPagesWithContext(aws.Context, *emrcontainers.ListManagedEndpointsInput, func(*emrcontainers.ListManagedEndpointsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*emrcontainers.ListTagsForResourceInput) (*emrcontainers.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *emrcontainers.ListTagsForResourceInput, ...request.Option) (*emrcontainers.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*emrcontainers.ListTagsForResourceInput) (*request.Request, *emrcontainers.ListTagsForResourceOutput)

	ListVirtualClusters(*emrcontainers.ListVirtualClustersInput) (*emrcontainers.ListVirtualClustersOutput, error)
	ListVirtualClustersWithContext(aws.Context, *emrcontainers.ListVirtualClustersInput, ...request.Option) (*emrcontainers.ListVirtualClustersOutput, error)
	ListVirtualClustersRequest(*emrcontainers.ListVirtualClustersInput) (*request.Request, *emrcontainers.ListVirtualClustersOutput)

	ListVirtualClustersPages(*emrcontainers.ListVirtualClustersInput, func(*emrcontainers.ListVirtualClustersOutput, bool) bool) error
	ListVirtualClustersPagesWithContext(aws.Context, *emrcontainers.ListVirtualClustersInput, func(*emrcontainers.ListVirtualClustersOutput, bool) bool, ...request.Option) error

	StartJobRun(*emrcontainers.StartJobRunInput) (*emrcontainers.StartJobRunOutput, error)
	StartJobRunWithContext(aws.Context, *emrcontainers.StartJobRunInput, ...request.Option) (*emrcontainers.StartJobRunOutput, error)
	StartJobRunRequest(*emrcontainers.StartJobRunInput) (*request.Request, *emrcontainers.StartJobRunOutput)

	TagResource(*emrcontainers.TagResourceInput) (*emrcontainers.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *emrcontainers.TagResourceInput, ...request.Option) (*emrcontainers.TagResourceOutput, error)
	TagResourceRequest(*emrcontainers.TagResourceInput) (*request.Request, *emrcontainers.TagResourceOutput)

	UntagResource(*emrcontainers.UntagResourceInput) (*emrcontainers.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *emrcontainers.UntagResourceInput, ...request.Option) (*emrcontainers.UntagResourceOutput, error)
	UntagResourceRequest(*emrcontainers.UntagResourceInput) (*request.Request, *emrcontainers.UntagResourceOutput)
}

var _ EMRContainersAPI = (*emrcontainers.EMRContainers)(nil)

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package cloudhsmv2iface provides an interface to enable mocking the AWS CloudHSM V2 service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package cloudhsmv2iface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/cloudhsmv2"
)

// CloudHSMV2API provides an interface to enable mocking the
// cloudhsmv2.CloudHSMV2 service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS CloudHSM V2.
//	func myFunc(svc cloudhsmv2iface.CloudHSMV2API) bool {
//	    // Make svc.CopyBackupToRegion request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := cloudhsmv2.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockCloudHSMV2Client struct {
//	    cloudhsmv2iface.CloudHSMV2API
//	}
//	func (m *mockCloudHSMV2Client) CopyBackupToRegion(input *cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockCloudHSMV2Client{}
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
type CloudHSMV2API interface {
	CopyBackupToRegion(*cloudhsmv2.CopyBackupToRegionInput) (*cloudhsmv2.CopyBackupToRegionOutput, error)
	CopyBackupToRegionWithContext(aws.Context, *cloudhsmv2.CopyBackupToRegionInput, ...request.Option) (*cloudhsmv2.CopyBackupToRegionOutput, error)
	CopyBackupToRegionRequest(*cloudhsmv2.CopyBackupToRegionInput) (*request.Request, *cloudhsmv2.CopyBackupToRegionOutput)

	CreateCluster(*cloudhsmv2.CreateClusterInput) (*cloudhsmv2.CreateClusterOutput, error)
	CreateClusterWithContext(aws.Context, *cloudhsmv2.CreateClusterInput, ...request.Option) (*cloudhsmv2.CreateClusterOutput, error)
	CreateClusterRequest(*cloudhsmv2.CreateClusterInput) (*request.Request, *cloudhsmv2.CreateClusterOutput)

	CreateHsm(*cloudhsmv2.CreateHsmInput) (*cloudhsmv2.CreateHsmOutput, error)
	CreateHsmWithContext(aws.Context, *cloudhsmv2.CreateHsmInput, ...request.Option) (*cloudhsmv2.CreateHsmOutput, error)
	CreateHsmRequest(*cloudhsmv2.CreateHsmInput) (*request.Request, *cloudhsmv2.CreateHsmOutput)

	DeleteBackup(*cloudhsmv2.DeleteBackupInput) (*cloudhsmv2.DeleteBackupOutput, error)
	DeleteBackupWithContext(aws.Context, *cloudhsmv2.DeleteBackupInput, ...request.Option) (*cloudhsmv2.DeleteBackupOutput, error)
	DeleteBackupRequest(*cloudhsmv2.DeleteBackupInput) (*request.Request, *cloudhsmv2.DeleteBackupOutput)

	DeleteCluster(*cloudhsmv2.DeleteClusterInput) (*cloudhsmv2.DeleteClusterOutput, error)
	DeleteClusterWithContext(aws.Context, *cloudhsmv2.DeleteClusterInput, ...request.Option) (*cloudhsmv2.DeleteClusterOutput, error)
	DeleteClusterRequest(*cloudhsmv2.DeleteClusterInput) (*request.Request, *cloudhsmv2.DeleteClusterOutput)

	DeleteHsm(*cloudhsmv2.DeleteHsmInput) (*cloudhsmv2.DeleteHsmOutput, error)
	DeleteHsmWithContext(aws.Context, *cloudhsmv2.DeleteHsmInput, ...request.Option) (*cloudhsmv2.DeleteHsmOutput, error)
	DeleteHsmRequest(*cloudhsmv2.DeleteHsmInput) (*request.Request, *cloudhsmv2.DeleteHsmOutput)

	DescribeBackups(*cloudhsmv2.DescribeBackupsInput) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeBackupsWithContext(aws.Context, *cloudhsmv2.DescribeBackupsInput, ...request.Option) (*cloudhsmv2.DescribeBackupsOutput, error)
	DescribeBackupsRequest(*cloudhsmv2.DescribeBackupsInput) (*request.Request, *cloudhsmv2.DescribeBackupsOutput)

	DescribeBackupsPages(*cloudhsmv2.DescribeBackupsInput, func(*cloudhsmv2.DescribeBackupsOutput, bool) bool) error
	DescribeBackupsPagesWithContext(aws.Context, *cloudhsmv2.DescribeBackupsInput, func(*cloudhsmv2.DescribeBackupsOutput, bool) bool, ...request.Option) error

	DescribeClusters(*cloudhsmv2.DescribeClustersInput) (*cloudhsmv2.DescribeClustersOutput, error)
	DescribeClustersWithContext(aws.Context, *cloudhsmv2.DescribeClustersInput, ...request.Option) (*cloudhsmv2.DescribeClustersOutput, error)
	DescribeClustersRequest(*cloudhsmv2.DescribeClustersInput) (*request.Request, *cloudhsmv2.DescribeClustersOutput)

	DescribeClustersPages(*cloudhsmv2.DescribeClustersInput, func(*cloudhsmv2.DescribeClustersOutput, bool) bool) error
	DescribeClustersPagesWithContext(aws.Context, *cloudhsmv2.DescribeClustersInput, func(*cloudhsmv2.DescribeClustersOutput, bool) bool, ...request.Option) error

	InitializeCluster(*cloudhsmv2.InitializeClusterInput) (*cloudhsmv2.InitializeClusterOutput, error)
	InitializeClusterWithContext(aws.Context, *cloudhsmv2.InitializeClusterInput, ...request.Option) (*cloudhsmv2.InitializeClusterOutput, error)
	InitializeClusterRequest(*cloudhsmv2.InitializeClusterInput) (*request.Request, *cloudhsmv2.InitializeClusterOutput)

	ListTags(*cloudhsmv2.ListTagsInput) (*cloudhsmv2.ListTagsOutput, error)
	ListTagsWithContext(aws.Context, *cloudhsmv2.ListTagsInput, ...request.Option) (*cloudhsmv2.ListTagsOutput, error)
	ListTagsRequest(*cloudhsmv2.ListTagsInput) (*request.Request, *cloudhsmv2.ListTagsOutput)

	ListTagsPages(*cloudhsmv2.ListTagsInput, func(*cloudhsmv2.ListTagsOutput, bool) bool) error
	ListTagsPagesWithContext(aws.Context, *cloudhsmv2.ListTagsInput, func(*cloudhsmv2.ListTagsOutput, bool) bool, ...request.Option) error

	ModifyBackupAttributes(*cloudhsmv2.ModifyBackupAttributesInput) (*cloudhsmv2.ModifyBackupAttributesOutput, error)
	ModifyBackupAttributesWithContext(aws.Context, *cloudhsmv2.ModifyBackupAttributesInput, ...request.Option) (*cloudhsmv2.ModifyBackupAttributesOutput, error)
	ModifyBackupAttributesRequest(*cloudhsmv2.ModifyBackupAttributesInput) (*request.Request, *cloudhsmv2.ModifyBackupAttributesOutput)

	ModifyCluster(*cloudhsmv2.ModifyClusterInput) (*cloudhsmv2.ModifyClusterOutput, error)
	ModifyClusterWithContext(aws.Context, *cloudhsmv2.ModifyClusterInput, ...request.Option) (*cloudhsmv2.ModifyClusterOutput, error)
	ModifyClusterRequest(*cloudhsmv2.ModifyClusterInput) (*request.Request, *cloudhsmv2.ModifyClusterOutput)

	RestoreBackup(*cloudhsmv2.RestoreBackupInput) (*cloudhsmv2.RestoreBackupOutput, error)
	RestoreBackupWithContext(aws.Context, *cloudhsmv2.RestoreBackupInput, ...request.Option) (*cloudhsmv2.RestoreBackupOutput, error)
	RestoreBackupRequest(*cloudhsmv2.RestoreBackupInput) (*request.Request, *cloudhsmv2.RestoreBackupOutput)

	TagResource(*cloudhsmv2.TagResourceInput) (*cloudhsmv2.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *cloudhsmv2.TagResourceInput, ...request.Option) (*cloudhsmv2.TagResourceOutput, error)
	TagResourceRequest(*cloudhsmv2.TagResourceInput) (*request.Request, *cloudhsmv2.TagResourceOutput)

	UntagResource(*cloudhsmv2.UntagResourceInput) (*cloudhsmv2.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *cloudhsmv2.UntagResourceInput, ...request.Option) (*cloudhsmv2.UntagResourceOutput, error)
	UntagResourceRequest(*cloudhsmv2.UntagResourceInput) (*request.Request, *cloudhsmv2.UntagResourceOutput)
}

var _ CloudHSMV2API = (*cloudhsmv2.CloudHSMV2)(nil)

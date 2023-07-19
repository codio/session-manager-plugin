// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package glacieriface provides an interface to enable mocking the Amazon Glacier service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package glacieriface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/glacier"
)

// GlacierAPI provides an interface to enable mocking the
// glacier.Glacier service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Glacier.
//	func myFunc(svc glacieriface.GlacierAPI) bool {
//	    // Make svc.AbortMultipartUpload request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := glacier.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockGlacierClient struct {
//	    glacieriface.GlacierAPI
//	}
//	func (m *mockGlacierClient) AbortMultipartUpload(input *glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockGlacierClient{}
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
type GlacierAPI interface {
	AbortMultipartUpload(*glacier.AbortMultipartUploadInput) (*glacier.AbortMultipartUploadOutput, error)
	AbortMultipartUploadWithContext(aws.Context, *glacier.AbortMultipartUploadInput, ...request.Option) (*glacier.AbortMultipartUploadOutput, error)
	AbortMultipartUploadRequest(*glacier.AbortMultipartUploadInput) (*request.Request, *glacier.AbortMultipartUploadOutput)

	AbortVaultLock(*glacier.AbortVaultLockInput) (*glacier.AbortVaultLockOutput, error)
	AbortVaultLockWithContext(aws.Context, *glacier.AbortVaultLockInput, ...request.Option) (*glacier.AbortVaultLockOutput, error)
	AbortVaultLockRequest(*glacier.AbortVaultLockInput) (*request.Request, *glacier.AbortVaultLockOutput)

	AddTagsToVault(*glacier.AddTagsToVaultInput) (*glacier.AddTagsToVaultOutput, error)
	AddTagsToVaultWithContext(aws.Context, *glacier.AddTagsToVaultInput, ...request.Option) (*glacier.AddTagsToVaultOutput, error)
	AddTagsToVaultRequest(*glacier.AddTagsToVaultInput) (*request.Request, *glacier.AddTagsToVaultOutput)

	CompleteMultipartUpload(*glacier.CompleteMultipartUploadInput) (*glacier.ArchiveCreationOutput, error)
	CompleteMultipartUploadWithContext(aws.Context, *glacier.CompleteMultipartUploadInput, ...request.Option) (*glacier.ArchiveCreationOutput, error)
	CompleteMultipartUploadRequest(*glacier.CompleteMultipartUploadInput) (*request.Request, *glacier.ArchiveCreationOutput)

	CompleteVaultLock(*glacier.CompleteVaultLockInput) (*glacier.CompleteVaultLockOutput, error)
	CompleteVaultLockWithContext(aws.Context, *glacier.CompleteVaultLockInput, ...request.Option) (*glacier.CompleteVaultLockOutput, error)
	CompleteVaultLockRequest(*glacier.CompleteVaultLockInput) (*request.Request, *glacier.CompleteVaultLockOutput)

	CreateVault(*glacier.CreateVaultInput) (*glacier.CreateVaultOutput, error)
	CreateVaultWithContext(aws.Context, *glacier.CreateVaultInput, ...request.Option) (*glacier.CreateVaultOutput, error)
	CreateVaultRequest(*glacier.CreateVaultInput) (*request.Request, *glacier.CreateVaultOutput)

	DeleteArchive(*glacier.DeleteArchiveInput) (*glacier.DeleteArchiveOutput, error)
	DeleteArchiveWithContext(aws.Context, *glacier.DeleteArchiveInput, ...request.Option) (*glacier.DeleteArchiveOutput, error)
	DeleteArchiveRequest(*glacier.DeleteArchiveInput) (*request.Request, *glacier.DeleteArchiveOutput)

	DeleteVault(*glacier.DeleteVaultInput) (*glacier.DeleteVaultOutput, error)
	DeleteVaultWithContext(aws.Context, *glacier.DeleteVaultInput, ...request.Option) (*glacier.DeleteVaultOutput, error)
	DeleteVaultRequest(*glacier.DeleteVaultInput) (*request.Request, *glacier.DeleteVaultOutput)

	DeleteVaultAccessPolicy(*glacier.DeleteVaultAccessPolicyInput) (*glacier.DeleteVaultAccessPolicyOutput, error)
	DeleteVaultAccessPolicyWithContext(aws.Context, *glacier.DeleteVaultAccessPolicyInput, ...request.Option) (*glacier.DeleteVaultAccessPolicyOutput, error)
	DeleteVaultAccessPolicyRequest(*glacier.DeleteVaultAccessPolicyInput) (*request.Request, *glacier.DeleteVaultAccessPolicyOutput)

	DeleteVaultNotifications(*glacier.DeleteVaultNotificationsInput) (*glacier.DeleteVaultNotificationsOutput, error)
	DeleteVaultNotificationsWithContext(aws.Context, *glacier.DeleteVaultNotificationsInput, ...request.Option) (*glacier.DeleteVaultNotificationsOutput, error)
	DeleteVaultNotificationsRequest(*glacier.DeleteVaultNotificationsInput) (*request.Request, *glacier.DeleteVaultNotificationsOutput)

	DescribeJob(*glacier.DescribeJobInput) (*glacier.JobDescription, error)
	DescribeJobWithContext(aws.Context, *glacier.DescribeJobInput, ...request.Option) (*glacier.JobDescription, error)
	DescribeJobRequest(*glacier.DescribeJobInput) (*request.Request, *glacier.JobDescription)

	DescribeVault(*glacier.DescribeVaultInput) (*glacier.DescribeVaultOutput, error)
	DescribeVaultWithContext(aws.Context, *glacier.DescribeVaultInput, ...request.Option) (*glacier.DescribeVaultOutput, error)
	DescribeVaultRequest(*glacier.DescribeVaultInput) (*request.Request, *glacier.DescribeVaultOutput)

	GetDataRetrievalPolicy(*glacier.GetDataRetrievalPolicyInput) (*glacier.GetDataRetrievalPolicyOutput, error)
	GetDataRetrievalPolicyWithContext(aws.Context, *glacier.GetDataRetrievalPolicyInput, ...request.Option) (*glacier.GetDataRetrievalPolicyOutput, error)
	GetDataRetrievalPolicyRequest(*glacier.GetDataRetrievalPolicyInput) (*request.Request, *glacier.GetDataRetrievalPolicyOutput)

	GetJobOutput(*glacier.GetJobOutputInput) (*glacier.GetJobOutputOutput, error)
	GetJobOutputWithContext(aws.Context, *glacier.GetJobOutputInput, ...request.Option) (*glacier.GetJobOutputOutput, error)
	GetJobOutputRequest(*glacier.GetJobOutputInput) (*request.Request, *glacier.GetJobOutputOutput)

	GetVaultAccessPolicy(*glacier.GetVaultAccessPolicyInput) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultAccessPolicyWithContext(aws.Context, *glacier.GetVaultAccessPolicyInput, ...request.Option) (*glacier.GetVaultAccessPolicyOutput, error)
	GetVaultAccessPolicyRequest(*glacier.GetVaultAccessPolicyInput) (*request.Request, *glacier.GetVaultAccessPolicyOutput)

	GetVaultLock(*glacier.GetVaultLockInput) (*glacier.GetVaultLockOutput, error)
	GetVaultLockWithContext(aws.Context, *glacier.GetVaultLockInput, ...request.Option) (*glacier.GetVaultLockOutput, error)
	GetVaultLockRequest(*glacier.GetVaultLockInput) (*request.Request, *glacier.GetVaultLockOutput)

	GetVaultNotifications(*glacier.GetVaultNotificationsInput) (*glacier.GetVaultNotificationsOutput, error)
	GetVaultNotificationsWithContext(aws.Context, *glacier.GetVaultNotificationsInput, ...request.Option) (*glacier.GetVaultNotificationsOutput, error)
	GetVaultNotificationsRequest(*glacier.GetVaultNotificationsInput) (*request.Request, *glacier.GetVaultNotificationsOutput)

	InitiateJob(*glacier.InitiateJobInput) (*glacier.InitiateJobOutput, error)
	InitiateJobWithContext(aws.Context, *glacier.InitiateJobInput, ...request.Option) (*glacier.InitiateJobOutput, error)
	InitiateJobRequest(*glacier.InitiateJobInput) (*request.Request, *glacier.InitiateJobOutput)

	InitiateMultipartUpload(*glacier.InitiateMultipartUploadInput) (*glacier.InitiateMultipartUploadOutput, error)
	InitiateMultipartUploadWithContext(aws.Context, *glacier.InitiateMultipartUploadInput, ...request.Option) (*glacier.InitiateMultipartUploadOutput, error)
	InitiateMultipartUploadRequest(*glacier.InitiateMultipartUploadInput) (*request.Request, *glacier.InitiateMultipartUploadOutput)

	InitiateVaultLock(*glacier.InitiateVaultLockInput) (*glacier.InitiateVaultLockOutput, error)
	InitiateVaultLockWithContext(aws.Context, *glacier.InitiateVaultLockInput, ...request.Option) (*glacier.InitiateVaultLockOutput, error)
	InitiateVaultLockRequest(*glacier.InitiateVaultLockInput) (*request.Request, *glacier.InitiateVaultLockOutput)

	ListJobs(*glacier.ListJobsInput) (*glacier.ListJobsOutput, error)
	ListJobsWithContext(aws.Context, *glacier.ListJobsInput, ...request.Option) (*glacier.ListJobsOutput, error)
	ListJobsRequest(*glacier.ListJobsInput) (*request.Request, *glacier.ListJobsOutput)

	ListJobsPages(*glacier.ListJobsInput, func(*glacier.ListJobsOutput, bool) bool) error
	ListJobsPagesWithContext(aws.Context, *glacier.ListJobsInput, func(*glacier.ListJobsOutput, bool) bool, ...request.Option) error

	ListMultipartUploads(*glacier.ListMultipartUploadsInput) (*glacier.ListMultipartUploadsOutput, error)
	ListMultipartUploadsWithContext(aws.Context, *glacier.ListMultipartUploadsInput, ...request.Option) (*glacier.ListMultipartUploadsOutput, error)
	ListMultipartUploadsRequest(*glacier.ListMultipartUploadsInput) (*request.Request, *glacier.ListMultipartUploadsOutput)

	ListMultipartUploadsPages(*glacier.ListMultipartUploadsInput, func(*glacier.ListMultipartUploadsOutput, bool) bool) error
	ListMultipartUploadsPagesWithContext(aws.Context, *glacier.ListMultipartUploadsInput, func(*glacier.ListMultipartUploadsOutput, bool) bool, ...request.Option) error

	ListParts(*glacier.ListPartsInput) (*glacier.ListPartsOutput, error)
	ListPartsWithContext(aws.Context, *glacier.ListPartsInput, ...request.Option) (*glacier.ListPartsOutput, error)
	ListPartsRequest(*glacier.ListPartsInput) (*request.Request, *glacier.ListPartsOutput)

	ListPartsPages(*glacier.ListPartsInput, func(*glacier.ListPartsOutput, bool) bool) error
	ListPartsPagesWithContext(aws.Context, *glacier.ListPartsInput, func(*glacier.ListPartsOutput, bool) bool, ...request.Option) error

	ListProvisionedCapacity(*glacier.ListProvisionedCapacityInput) (*glacier.ListProvisionedCapacityOutput, error)
	ListProvisionedCapacityWithContext(aws.Context, *glacier.ListProvisionedCapacityInput, ...request.Option) (*glacier.ListProvisionedCapacityOutput, error)
	ListProvisionedCapacityRequest(*glacier.ListProvisionedCapacityInput) (*request.Request, *glacier.ListProvisionedCapacityOutput)

	ListTagsForVault(*glacier.ListTagsForVaultInput) (*glacier.ListTagsForVaultOutput, error)
	ListTagsForVaultWithContext(aws.Context, *glacier.ListTagsForVaultInput, ...request.Option) (*glacier.ListTagsForVaultOutput, error)
	ListTagsForVaultRequest(*glacier.ListTagsForVaultInput) (*request.Request, *glacier.ListTagsForVaultOutput)

	ListVaults(*glacier.ListVaultsInput) (*glacier.ListVaultsOutput, error)
	ListVaultsWithContext(aws.Context, *glacier.ListVaultsInput, ...request.Option) (*glacier.ListVaultsOutput, error)
	ListVaultsRequest(*glacier.ListVaultsInput) (*request.Request, *glacier.ListVaultsOutput)

	ListVaultsPages(*glacier.ListVaultsInput, func(*glacier.ListVaultsOutput, bool) bool) error
	ListVaultsPagesWithContext(aws.Context, *glacier.ListVaultsInput, func(*glacier.ListVaultsOutput, bool) bool, ...request.Option) error

	PurchaseProvisionedCapacity(*glacier.PurchaseProvisionedCapacityInput) (*glacier.PurchaseProvisionedCapacityOutput, error)
	PurchaseProvisionedCapacityWithContext(aws.Context, *glacier.PurchaseProvisionedCapacityInput, ...request.Option) (*glacier.PurchaseProvisionedCapacityOutput, error)
	PurchaseProvisionedCapacityRequest(*glacier.PurchaseProvisionedCapacityInput) (*request.Request, *glacier.PurchaseProvisionedCapacityOutput)

	RemoveTagsFromVault(*glacier.RemoveTagsFromVaultInput) (*glacier.RemoveTagsFromVaultOutput, error)
	RemoveTagsFromVaultWithContext(aws.Context, *glacier.RemoveTagsFromVaultInput, ...request.Option) (*glacier.RemoveTagsFromVaultOutput, error)
	RemoveTagsFromVaultRequest(*glacier.RemoveTagsFromVaultInput) (*request.Request, *glacier.RemoveTagsFromVaultOutput)

	SetDataRetrievalPolicy(*glacier.SetDataRetrievalPolicyInput) (*glacier.SetDataRetrievalPolicyOutput, error)
	SetDataRetrievalPolicyWithContext(aws.Context, *glacier.SetDataRetrievalPolicyInput, ...request.Option) (*glacier.SetDataRetrievalPolicyOutput, error)
	SetDataRetrievalPolicyRequest(*glacier.SetDataRetrievalPolicyInput) (*request.Request, *glacier.SetDataRetrievalPolicyOutput)

	SetVaultAccessPolicy(*glacier.SetVaultAccessPolicyInput) (*glacier.SetVaultAccessPolicyOutput, error)
	SetVaultAccessPolicyWithContext(aws.Context, *glacier.SetVaultAccessPolicyInput, ...request.Option) (*glacier.SetVaultAccessPolicyOutput, error)
	SetVaultAccessPolicyRequest(*glacier.SetVaultAccessPolicyInput) (*request.Request, *glacier.SetVaultAccessPolicyOutput)

	SetVaultNotifications(*glacier.SetVaultNotificationsInput) (*glacier.SetVaultNotificationsOutput, error)
	SetVaultNotificationsWithContext(aws.Context, *glacier.SetVaultNotificationsInput, ...request.Option) (*glacier.SetVaultNotificationsOutput, error)
	SetVaultNotificationsRequest(*glacier.SetVaultNotificationsInput) (*request.Request, *glacier.SetVaultNotificationsOutput)

	UploadArchive(*glacier.UploadArchiveInput) (*glacier.ArchiveCreationOutput, error)
	UploadArchiveWithContext(aws.Context, *glacier.UploadArchiveInput, ...request.Option) (*glacier.ArchiveCreationOutput, error)
	UploadArchiveRequest(*glacier.UploadArchiveInput) (*request.Request, *glacier.ArchiveCreationOutput)

	UploadMultipartPart(*glacier.UploadMultipartPartInput) (*glacier.UploadMultipartPartOutput, error)
	UploadMultipartPartWithContext(aws.Context, *glacier.UploadMultipartPartInput, ...request.Option) (*glacier.UploadMultipartPartOutput, error)
	UploadMultipartPartRequest(*glacier.UploadMultipartPartInput) (*request.Request, *glacier.UploadMultipartPartOutput)

	WaitUntilVaultExists(*glacier.DescribeVaultInput) error
	WaitUntilVaultExistsWithContext(aws.Context, *glacier.DescribeVaultInput, ...request.WaiterOption) error

	WaitUntilVaultNotExists(*glacier.DescribeVaultInput) error
	WaitUntilVaultNotExistsWithContext(aws.Context, *glacier.DescribeVaultInput, ...request.WaiterOption) error
}

var _ GlacierAPI = (*glacier.Glacier)(nil)

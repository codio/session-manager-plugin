// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package mgniface provides an interface to enable mocking the Application Migration Service service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package mgniface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/mgn"
)

// MgnAPI provides an interface to enable mocking the
// mgn.Mgn service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Application Migration Service.
//	func myFunc(svc mgniface.MgnAPI) bool {
//	    // Make svc.ArchiveApplication request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := mgn.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockMgnClient struct {
//	    mgniface.MgnAPI
//	}
//	func (m *mockMgnClient) ArchiveApplication(input *mgn.ArchiveApplicationInput) (*mgn.ArchiveApplicationOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockMgnClient{}
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
type MgnAPI interface {
	ArchiveApplication(*mgn.ArchiveApplicationInput) (*mgn.ArchiveApplicationOutput, error)
	ArchiveApplicationWithContext(aws.Context, *mgn.ArchiveApplicationInput, ...request.Option) (*mgn.ArchiveApplicationOutput, error)
	ArchiveApplicationRequest(*mgn.ArchiveApplicationInput) (*request.Request, *mgn.ArchiveApplicationOutput)

	ArchiveWave(*mgn.ArchiveWaveInput) (*mgn.ArchiveWaveOutput, error)
	ArchiveWaveWithContext(aws.Context, *mgn.ArchiveWaveInput, ...request.Option) (*mgn.ArchiveWaveOutput, error)
	ArchiveWaveRequest(*mgn.ArchiveWaveInput) (*request.Request, *mgn.ArchiveWaveOutput)

	AssociateApplications(*mgn.AssociateApplicationsInput) (*mgn.AssociateApplicationsOutput, error)
	AssociateApplicationsWithContext(aws.Context, *mgn.AssociateApplicationsInput, ...request.Option) (*mgn.AssociateApplicationsOutput, error)
	AssociateApplicationsRequest(*mgn.AssociateApplicationsInput) (*request.Request, *mgn.AssociateApplicationsOutput)

	AssociateSourceServers(*mgn.AssociateSourceServersInput) (*mgn.AssociateSourceServersOutput, error)
	AssociateSourceServersWithContext(aws.Context, *mgn.AssociateSourceServersInput, ...request.Option) (*mgn.AssociateSourceServersOutput, error)
	AssociateSourceServersRequest(*mgn.AssociateSourceServersInput) (*request.Request, *mgn.AssociateSourceServersOutput)

	ChangeServerLifeCycleState(*mgn.ChangeServerLifeCycleStateInput) (*mgn.ChangeServerLifeCycleStateOutput, error)
	ChangeServerLifeCycleStateWithContext(aws.Context, *mgn.ChangeServerLifeCycleStateInput, ...request.Option) (*mgn.ChangeServerLifeCycleStateOutput, error)
	ChangeServerLifeCycleStateRequest(*mgn.ChangeServerLifeCycleStateInput) (*request.Request, *mgn.ChangeServerLifeCycleStateOutput)

	CreateApplication(*mgn.CreateApplicationInput) (*mgn.CreateApplicationOutput, error)
	CreateApplicationWithContext(aws.Context, *mgn.CreateApplicationInput, ...request.Option) (*mgn.CreateApplicationOutput, error)
	CreateApplicationRequest(*mgn.CreateApplicationInput) (*request.Request, *mgn.CreateApplicationOutput)

	CreateLaunchConfigurationTemplate(*mgn.CreateLaunchConfigurationTemplateInput) (*mgn.CreateLaunchConfigurationTemplateOutput, error)
	CreateLaunchConfigurationTemplateWithContext(aws.Context, *mgn.CreateLaunchConfigurationTemplateInput, ...request.Option) (*mgn.CreateLaunchConfigurationTemplateOutput, error)
	CreateLaunchConfigurationTemplateRequest(*mgn.CreateLaunchConfigurationTemplateInput) (*request.Request, *mgn.CreateLaunchConfigurationTemplateOutput)

	CreateReplicationConfigurationTemplate(*mgn.CreateReplicationConfigurationTemplateInput) (*mgn.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateWithContext(aws.Context, *mgn.CreateReplicationConfigurationTemplateInput, ...request.Option) (*mgn.CreateReplicationConfigurationTemplateOutput, error)
	CreateReplicationConfigurationTemplateRequest(*mgn.CreateReplicationConfigurationTemplateInput) (*request.Request, *mgn.CreateReplicationConfigurationTemplateOutput)

	CreateWave(*mgn.CreateWaveInput) (*mgn.CreateWaveOutput, error)
	CreateWaveWithContext(aws.Context, *mgn.CreateWaveInput, ...request.Option) (*mgn.CreateWaveOutput, error)
	CreateWaveRequest(*mgn.CreateWaveInput) (*request.Request, *mgn.CreateWaveOutput)

	DeleteApplication(*mgn.DeleteApplicationInput) (*mgn.DeleteApplicationOutput, error)
	DeleteApplicationWithContext(aws.Context, *mgn.DeleteApplicationInput, ...request.Option) (*mgn.DeleteApplicationOutput, error)
	DeleteApplicationRequest(*mgn.DeleteApplicationInput) (*request.Request, *mgn.DeleteApplicationOutput)

	DeleteJob(*mgn.DeleteJobInput) (*mgn.DeleteJobOutput, error)
	DeleteJobWithContext(aws.Context, *mgn.DeleteJobInput, ...request.Option) (*mgn.DeleteJobOutput, error)
	DeleteJobRequest(*mgn.DeleteJobInput) (*request.Request, *mgn.DeleteJobOutput)

	DeleteLaunchConfigurationTemplate(*mgn.DeleteLaunchConfigurationTemplateInput) (*mgn.DeleteLaunchConfigurationTemplateOutput, error)
	DeleteLaunchConfigurationTemplateWithContext(aws.Context, *mgn.DeleteLaunchConfigurationTemplateInput, ...request.Option) (*mgn.DeleteLaunchConfigurationTemplateOutput, error)
	DeleteLaunchConfigurationTemplateRequest(*mgn.DeleteLaunchConfigurationTemplateInput) (*request.Request, *mgn.DeleteLaunchConfigurationTemplateOutput)

	DeleteReplicationConfigurationTemplate(*mgn.DeleteReplicationConfigurationTemplateInput) (*mgn.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateWithContext(aws.Context, *mgn.DeleteReplicationConfigurationTemplateInput, ...request.Option) (*mgn.DeleteReplicationConfigurationTemplateOutput, error)
	DeleteReplicationConfigurationTemplateRequest(*mgn.DeleteReplicationConfigurationTemplateInput) (*request.Request, *mgn.DeleteReplicationConfigurationTemplateOutput)

	DeleteSourceServer(*mgn.DeleteSourceServerInput) (*mgn.DeleteSourceServerOutput, error)
	DeleteSourceServerWithContext(aws.Context, *mgn.DeleteSourceServerInput, ...request.Option) (*mgn.DeleteSourceServerOutput, error)
	DeleteSourceServerRequest(*mgn.DeleteSourceServerInput) (*request.Request, *mgn.DeleteSourceServerOutput)

	DeleteVcenterClient(*mgn.DeleteVcenterClientInput) (*mgn.DeleteVcenterClientOutput, error)
	DeleteVcenterClientWithContext(aws.Context, *mgn.DeleteVcenterClientInput, ...request.Option) (*mgn.DeleteVcenterClientOutput, error)
	DeleteVcenterClientRequest(*mgn.DeleteVcenterClientInput) (*request.Request, *mgn.DeleteVcenterClientOutput)

	DeleteWave(*mgn.DeleteWaveInput) (*mgn.DeleteWaveOutput, error)
	DeleteWaveWithContext(aws.Context, *mgn.DeleteWaveInput, ...request.Option) (*mgn.DeleteWaveOutput, error)
	DeleteWaveRequest(*mgn.DeleteWaveInput) (*request.Request, *mgn.DeleteWaveOutput)

	DescribeJobLogItems(*mgn.DescribeJobLogItemsInput) (*mgn.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsWithContext(aws.Context, *mgn.DescribeJobLogItemsInput, ...request.Option) (*mgn.DescribeJobLogItemsOutput, error)
	DescribeJobLogItemsRequest(*mgn.DescribeJobLogItemsInput) (*request.Request, *mgn.DescribeJobLogItemsOutput)

	DescribeJobLogItemsPages(*mgn.DescribeJobLogItemsInput, func(*mgn.DescribeJobLogItemsOutput, bool) bool) error
	DescribeJobLogItemsPagesWithContext(aws.Context, *mgn.DescribeJobLogItemsInput, func(*mgn.DescribeJobLogItemsOutput, bool) bool, ...request.Option) error

	DescribeJobs(*mgn.DescribeJobsInput) (*mgn.DescribeJobsOutput, error)
	DescribeJobsWithContext(aws.Context, *mgn.DescribeJobsInput, ...request.Option) (*mgn.DescribeJobsOutput, error)
	DescribeJobsRequest(*mgn.DescribeJobsInput) (*request.Request, *mgn.DescribeJobsOutput)

	DescribeJobsPages(*mgn.DescribeJobsInput, func(*mgn.DescribeJobsOutput, bool) bool) error
	DescribeJobsPagesWithContext(aws.Context, *mgn.DescribeJobsInput, func(*mgn.DescribeJobsOutput, bool) bool, ...request.Option) error

	DescribeLaunchConfigurationTemplates(*mgn.DescribeLaunchConfigurationTemplatesInput) (*mgn.DescribeLaunchConfigurationTemplatesOutput, error)
	DescribeLaunchConfigurationTemplatesWithContext(aws.Context, *mgn.DescribeLaunchConfigurationTemplatesInput, ...request.Option) (*mgn.DescribeLaunchConfigurationTemplatesOutput, error)
	DescribeLaunchConfigurationTemplatesRequest(*mgn.DescribeLaunchConfigurationTemplatesInput) (*request.Request, *mgn.DescribeLaunchConfigurationTemplatesOutput)

	DescribeLaunchConfigurationTemplatesPages(*mgn.DescribeLaunchConfigurationTemplatesInput, func(*mgn.DescribeLaunchConfigurationTemplatesOutput, bool) bool) error
	DescribeLaunchConfigurationTemplatesPagesWithContext(aws.Context, *mgn.DescribeLaunchConfigurationTemplatesInput, func(*mgn.DescribeLaunchConfigurationTemplatesOutput, bool) bool, ...request.Option) error

	DescribeReplicationConfigurationTemplates(*mgn.DescribeReplicationConfigurationTemplatesInput) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesWithContext(aws.Context, *mgn.DescribeReplicationConfigurationTemplatesInput, ...request.Option) (*mgn.DescribeReplicationConfigurationTemplatesOutput, error)
	DescribeReplicationConfigurationTemplatesRequest(*mgn.DescribeReplicationConfigurationTemplatesInput) (*request.Request, *mgn.DescribeReplicationConfigurationTemplatesOutput)

	DescribeReplicationConfigurationTemplatesPages(*mgn.DescribeReplicationConfigurationTemplatesInput, func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool) error
	DescribeReplicationConfigurationTemplatesPagesWithContext(aws.Context, *mgn.DescribeReplicationConfigurationTemplatesInput, func(*mgn.DescribeReplicationConfigurationTemplatesOutput, bool) bool, ...request.Option) error

	DescribeSourceServers(*mgn.DescribeSourceServersInput) (*mgn.DescribeSourceServersOutput, error)
	DescribeSourceServersWithContext(aws.Context, *mgn.DescribeSourceServersInput, ...request.Option) (*mgn.DescribeSourceServersOutput, error)
	DescribeSourceServersRequest(*mgn.DescribeSourceServersInput) (*request.Request, *mgn.DescribeSourceServersOutput)

	DescribeSourceServersPages(*mgn.DescribeSourceServersInput, func(*mgn.DescribeSourceServersOutput, bool) bool) error
	DescribeSourceServersPagesWithContext(aws.Context, *mgn.DescribeSourceServersInput, func(*mgn.DescribeSourceServersOutput, bool) bool, ...request.Option) error

	DescribeVcenterClients(*mgn.DescribeVcenterClientsInput) (*mgn.DescribeVcenterClientsOutput, error)
	DescribeVcenterClientsWithContext(aws.Context, *mgn.DescribeVcenterClientsInput, ...request.Option) (*mgn.DescribeVcenterClientsOutput, error)
	DescribeVcenterClientsRequest(*mgn.DescribeVcenterClientsInput) (*request.Request, *mgn.DescribeVcenterClientsOutput)

	DescribeVcenterClientsPages(*mgn.DescribeVcenterClientsInput, func(*mgn.DescribeVcenterClientsOutput, bool) bool) error
	DescribeVcenterClientsPagesWithContext(aws.Context, *mgn.DescribeVcenterClientsInput, func(*mgn.DescribeVcenterClientsOutput, bool) bool, ...request.Option) error

	DisassociateApplications(*mgn.DisassociateApplicationsInput) (*mgn.DisassociateApplicationsOutput, error)
	DisassociateApplicationsWithContext(aws.Context, *mgn.DisassociateApplicationsInput, ...request.Option) (*mgn.DisassociateApplicationsOutput, error)
	DisassociateApplicationsRequest(*mgn.DisassociateApplicationsInput) (*request.Request, *mgn.DisassociateApplicationsOutput)

	DisassociateSourceServers(*mgn.DisassociateSourceServersInput) (*mgn.DisassociateSourceServersOutput, error)
	DisassociateSourceServersWithContext(aws.Context, *mgn.DisassociateSourceServersInput, ...request.Option) (*mgn.DisassociateSourceServersOutput, error)
	DisassociateSourceServersRequest(*mgn.DisassociateSourceServersInput) (*request.Request, *mgn.DisassociateSourceServersOutput)

	DisconnectFromService(*mgn.DisconnectFromServiceInput) (*mgn.DisconnectFromServiceOutput, error)
	DisconnectFromServiceWithContext(aws.Context, *mgn.DisconnectFromServiceInput, ...request.Option) (*mgn.DisconnectFromServiceOutput, error)
	DisconnectFromServiceRequest(*mgn.DisconnectFromServiceInput) (*request.Request, *mgn.DisconnectFromServiceOutput)

	FinalizeCutover(*mgn.FinalizeCutoverInput) (*mgn.FinalizeCutoverOutput, error)
	FinalizeCutoverWithContext(aws.Context, *mgn.FinalizeCutoverInput, ...request.Option) (*mgn.FinalizeCutoverOutput, error)
	FinalizeCutoverRequest(*mgn.FinalizeCutoverInput) (*request.Request, *mgn.FinalizeCutoverOutput)

	GetLaunchConfiguration(*mgn.GetLaunchConfigurationInput) (*mgn.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationWithContext(aws.Context, *mgn.GetLaunchConfigurationInput, ...request.Option) (*mgn.GetLaunchConfigurationOutput, error)
	GetLaunchConfigurationRequest(*mgn.GetLaunchConfigurationInput) (*request.Request, *mgn.GetLaunchConfigurationOutput)

	GetReplicationConfiguration(*mgn.GetReplicationConfigurationInput) (*mgn.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationWithContext(aws.Context, *mgn.GetReplicationConfigurationInput, ...request.Option) (*mgn.GetReplicationConfigurationOutput, error)
	GetReplicationConfigurationRequest(*mgn.GetReplicationConfigurationInput) (*request.Request, *mgn.GetReplicationConfigurationOutput)

	InitializeService(*mgn.InitializeServiceInput) (*mgn.InitializeServiceOutput, error)
	InitializeServiceWithContext(aws.Context, *mgn.InitializeServiceInput, ...request.Option) (*mgn.InitializeServiceOutput, error)
	InitializeServiceRequest(*mgn.InitializeServiceInput) (*request.Request, *mgn.InitializeServiceOutput)

	ListApplications(*mgn.ListApplicationsInput) (*mgn.ListApplicationsOutput, error)
	ListApplicationsWithContext(aws.Context, *mgn.ListApplicationsInput, ...request.Option) (*mgn.ListApplicationsOutput, error)
	ListApplicationsRequest(*mgn.ListApplicationsInput) (*request.Request, *mgn.ListApplicationsOutput)

	ListApplicationsPages(*mgn.ListApplicationsInput, func(*mgn.ListApplicationsOutput, bool) bool) error
	ListApplicationsPagesWithContext(aws.Context, *mgn.ListApplicationsInput, func(*mgn.ListApplicationsOutput, bool) bool, ...request.Option) error

	ListExportErrors(*mgn.ListExportErrorsInput) (*mgn.ListExportErrorsOutput, error)
	ListExportErrorsWithContext(aws.Context, *mgn.ListExportErrorsInput, ...request.Option) (*mgn.ListExportErrorsOutput, error)
	ListExportErrorsRequest(*mgn.ListExportErrorsInput) (*request.Request, *mgn.ListExportErrorsOutput)

	ListExportErrorsPages(*mgn.ListExportErrorsInput, func(*mgn.ListExportErrorsOutput, bool) bool) error
	ListExportErrorsPagesWithContext(aws.Context, *mgn.ListExportErrorsInput, func(*mgn.ListExportErrorsOutput, bool) bool, ...request.Option) error

	ListExports(*mgn.ListExportsInput) (*mgn.ListExportsOutput, error)
	ListExportsWithContext(aws.Context, *mgn.ListExportsInput, ...request.Option) (*mgn.ListExportsOutput, error)
	ListExportsRequest(*mgn.ListExportsInput) (*request.Request, *mgn.ListExportsOutput)

	ListExportsPages(*mgn.ListExportsInput, func(*mgn.ListExportsOutput, bool) bool) error
	ListExportsPagesWithContext(aws.Context, *mgn.ListExportsInput, func(*mgn.ListExportsOutput, bool) bool, ...request.Option) error

	ListImportErrors(*mgn.ListImportErrorsInput) (*mgn.ListImportErrorsOutput, error)
	ListImportErrorsWithContext(aws.Context, *mgn.ListImportErrorsInput, ...request.Option) (*mgn.ListImportErrorsOutput, error)
	ListImportErrorsRequest(*mgn.ListImportErrorsInput) (*request.Request, *mgn.ListImportErrorsOutput)

	ListImportErrorsPages(*mgn.ListImportErrorsInput, func(*mgn.ListImportErrorsOutput, bool) bool) error
	ListImportErrorsPagesWithContext(aws.Context, *mgn.ListImportErrorsInput, func(*mgn.ListImportErrorsOutput, bool) bool, ...request.Option) error

	ListImports(*mgn.ListImportsInput) (*mgn.ListImportsOutput, error)
	ListImportsWithContext(aws.Context, *mgn.ListImportsInput, ...request.Option) (*mgn.ListImportsOutput, error)
	ListImportsRequest(*mgn.ListImportsInput) (*request.Request, *mgn.ListImportsOutput)

	ListImportsPages(*mgn.ListImportsInput, func(*mgn.ListImportsOutput, bool) bool) error
	ListImportsPagesWithContext(aws.Context, *mgn.ListImportsInput, func(*mgn.ListImportsOutput, bool) bool, ...request.Option) error

	ListManagedAccounts(*mgn.ListManagedAccountsInput) (*mgn.ListManagedAccountsOutput, error)
	ListManagedAccountsWithContext(aws.Context, *mgn.ListManagedAccountsInput, ...request.Option) (*mgn.ListManagedAccountsOutput, error)
	ListManagedAccountsRequest(*mgn.ListManagedAccountsInput) (*request.Request, *mgn.ListManagedAccountsOutput)

	ListManagedAccountsPages(*mgn.ListManagedAccountsInput, func(*mgn.ListManagedAccountsOutput, bool) bool) error
	ListManagedAccountsPagesWithContext(aws.Context, *mgn.ListManagedAccountsInput, func(*mgn.ListManagedAccountsOutput, bool) bool, ...request.Option) error

	ListSourceServerActions(*mgn.ListSourceServerActionsInput) (*mgn.ListSourceServerActionsOutput, error)
	ListSourceServerActionsWithContext(aws.Context, *mgn.ListSourceServerActionsInput, ...request.Option) (*mgn.ListSourceServerActionsOutput, error)
	ListSourceServerActionsRequest(*mgn.ListSourceServerActionsInput) (*request.Request, *mgn.ListSourceServerActionsOutput)

	ListSourceServerActionsPages(*mgn.ListSourceServerActionsInput, func(*mgn.ListSourceServerActionsOutput, bool) bool) error
	ListSourceServerActionsPagesWithContext(aws.Context, *mgn.ListSourceServerActionsInput, func(*mgn.ListSourceServerActionsOutput, bool) bool, ...request.Option) error

	ListTagsForResource(*mgn.ListTagsForResourceInput) (*mgn.ListTagsForResourceOutput, error)
	ListTagsForResourceWithContext(aws.Context, *mgn.ListTagsForResourceInput, ...request.Option) (*mgn.ListTagsForResourceOutput, error)
	ListTagsForResourceRequest(*mgn.ListTagsForResourceInput) (*request.Request, *mgn.ListTagsForResourceOutput)

	ListTemplateActions(*mgn.ListTemplateActionsInput) (*mgn.ListTemplateActionsOutput, error)
	ListTemplateActionsWithContext(aws.Context, *mgn.ListTemplateActionsInput, ...request.Option) (*mgn.ListTemplateActionsOutput, error)
	ListTemplateActionsRequest(*mgn.ListTemplateActionsInput) (*request.Request, *mgn.ListTemplateActionsOutput)

	ListTemplateActionsPages(*mgn.ListTemplateActionsInput, func(*mgn.ListTemplateActionsOutput, bool) bool) error
	ListTemplateActionsPagesWithContext(aws.Context, *mgn.ListTemplateActionsInput, func(*mgn.ListTemplateActionsOutput, bool) bool, ...request.Option) error

	ListWaves(*mgn.ListWavesInput) (*mgn.ListWavesOutput, error)
	ListWavesWithContext(aws.Context, *mgn.ListWavesInput, ...request.Option) (*mgn.ListWavesOutput, error)
	ListWavesRequest(*mgn.ListWavesInput) (*request.Request, *mgn.ListWavesOutput)

	ListWavesPages(*mgn.ListWavesInput, func(*mgn.ListWavesOutput, bool) bool) error
	ListWavesPagesWithContext(aws.Context, *mgn.ListWavesInput, func(*mgn.ListWavesOutput, bool) bool, ...request.Option) error

	MarkAsArchived(*mgn.MarkAsArchivedInput) (*mgn.MarkAsArchivedOutput, error)
	MarkAsArchivedWithContext(aws.Context, *mgn.MarkAsArchivedInput, ...request.Option) (*mgn.MarkAsArchivedOutput, error)
	MarkAsArchivedRequest(*mgn.MarkAsArchivedInput) (*request.Request, *mgn.MarkAsArchivedOutput)

	PauseReplication(*mgn.PauseReplicationInput) (*mgn.PauseReplicationOutput, error)
	PauseReplicationWithContext(aws.Context, *mgn.PauseReplicationInput, ...request.Option) (*mgn.PauseReplicationOutput, error)
	PauseReplicationRequest(*mgn.PauseReplicationInput) (*request.Request, *mgn.PauseReplicationOutput)

	PutSourceServerAction(*mgn.PutSourceServerActionInput) (*mgn.PutSourceServerActionOutput, error)
	PutSourceServerActionWithContext(aws.Context, *mgn.PutSourceServerActionInput, ...request.Option) (*mgn.PutSourceServerActionOutput, error)
	PutSourceServerActionRequest(*mgn.PutSourceServerActionInput) (*request.Request, *mgn.PutSourceServerActionOutput)

	PutTemplateAction(*mgn.PutTemplateActionInput) (*mgn.PutTemplateActionOutput, error)
	PutTemplateActionWithContext(aws.Context, *mgn.PutTemplateActionInput, ...request.Option) (*mgn.PutTemplateActionOutput, error)
	PutTemplateActionRequest(*mgn.PutTemplateActionInput) (*request.Request, *mgn.PutTemplateActionOutput)

	RemoveSourceServerAction(*mgn.RemoveSourceServerActionInput) (*mgn.RemoveSourceServerActionOutput, error)
	RemoveSourceServerActionWithContext(aws.Context, *mgn.RemoveSourceServerActionInput, ...request.Option) (*mgn.RemoveSourceServerActionOutput, error)
	RemoveSourceServerActionRequest(*mgn.RemoveSourceServerActionInput) (*request.Request, *mgn.RemoveSourceServerActionOutput)

	RemoveTemplateAction(*mgn.RemoveTemplateActionInput) (*mgn.RemoveTemplateActionOutput, error)
	RemoveTemplateActionWithContext(aws.Context, *mgn.RemoveTemplateActionInput, ...request.Option) (*mgn.RemoveTemplateActionOutput, error)
	RemoveTemplateActionRequest(*mgn.RemoveTemplateActionInput) (*request.Request, *mgn.RemoveTemplateActionOutput)

	ResumeReplication(*mgn.ResumeReplicationInput) (*mgn.ResumeReplicationOutput, error)
	ResumeReplicationWithContext(aws.Context, *mgn.ResumeReplicationInput, ...request.Option) (*mgn.ResumeReplicationOutput, error)
	ResumeReplicationRequest(*mgn.ResumeReplicationInput) (*request.Request, *mgn.ResumeReplicationOutput)

	RetryDataReplication(*mgn.RetryDataReplicationInput) (*mgn.RetryDataReplicationOutput, error)
	RetryDataReplicationWithContext(aws.Context, *mgn.RetryDataReplicationInput, ...request.Option) (*mgn.RetryDataReplicationOutput, error)
	RetryDataReplicationRequest(*mgn.RetryDataReplicationInput) (*request.Request, *mgn.RetryDataReplicationOutput)

	StartCutover(*mgn.StartCutoverInput) (*mgn.StartCutoverOutput, error)
	StartCutoverWithContext(aws.Context, *mgn.StartCutoverInput, ...request.Option) (*mgn.StartCutoverOutput, error)
	StartCutoverRequest(*mgn.StartCutoverInput) (*request.Request, *mgn.StartCutoverOutput)

	StartExport(*mgn.StartExportInput) (*mgn.StartExportOutput, error)
	StartExportWithContext(aws.Context, *mgn.StartExportInput, ...request.Option) (*mgn.StartExportOutput, error)
	StartExportRequest(*mgn.StartExportInput) (*request.Request, *mgn.StartExportOutput)

	StartImport(*mgn.StartImportInput) (*mgn.StartImportOutput, error)
	StartImportWithContext(aws.Context, *mgn.StartImportInput, ...request.Option) (*mgn.StartImportOutput, error)
	StartImportRequest(*mgn.StartImportInput) (*request.Request, *mgn.StartImportOutput)

	StartReplication(*mgn.StartReplicationInput) (*mgn.StartReplicationOutput, error)
	StartReplicationWithContext(aws.Context, *mgn.StartReplicationInput, ...request.Option) (*mgn.StartReplicationOutput, error)
	StartReplicationRequest(*mgn.StartReplicationInput) (*request.Request, *mgn.StartReplicationOutput)

	StartTest(*mgn.StartTestInput) (*mgn.StartTestOutput, error)
	StartTestWithContext(aws.Context, *mgn.StartTestInput, ...request.Option) (*mgn.StartTestOutput, error)
	StartTestRequest(*mgn.StartTestInput) (*request.Request, *mgn.StartTestOutput)

	StopReplication(*mgn.StopReplicationInput) (*mgn.StopReplicationOutput, error)
	StopReplicationWithContext(aws.Context, *mgn.StopReplicationInput, ...request.Option) (*mgn.StopReplicationOutput, error)
	StopReplicationRequest(*mgn.StopReplicationInput) (*request.Request, *mgn.StopReplicationOutput)

	TagResource(*mgn.TagResourceInput) (*mgn.TagResourceOutput, error)
	TagResourceWithContext(aws.Context, *mgn.TagResourceInput, ...request.Option) (*mgn.TagResourceOutput, error)
	TagResourceRequest(*mgn.TagResourceInput) (*request.Request, *mgn.TagResourceOutput)

	TerminateTargetInstances(*mgn.TerminateTargetInstancesInput) (*mgn.TerminateTargetInstancesOutput, error)
	TerminateTargetInstancesWithContext(aws.Context, *mgn.TerminateTargetInstancesInput, ...request.Option) (*mgn.TerminateTargetInstancesOutput, error)
	TerminateTargetInstancesRequest(*mgn.TerminateTargetInstancesInput) (*request.Request, *mgn.TerminateTargetInstancesOutput)

	UnarchiveApplication(*mgn.UnarchiveApplicationInput) (*mgn.UnarchiveApplicationOutput, error)
	UnarchiveApplicationWithContext(aws.Context, *mgn.UnarchiveApplicationInput, ...request.Option) (*mgn.UnarchiveApplicationOutput, error)
	UnarchiveApplicationRequest(*mgn.UnarchiveApplicationInput) (*request.Request, *mgn.UnarchiveApplicationOutput)

	UnarchiveWave(*mgn.UnarchiveWaveInput) (*mgn.UnarchiveWaveOutput, error)
	UnarchiveWaveWithContext(aws.Context, *mgn.UnarchiveWaveInput, ...request.Option) (*mgn.UnarchiveWaveOutput, error)
	UnarchiveWaveRequest(*mgn.UnarchiveWaveInput) (*request.Request, *mgn.UnarchiveWaveOutput)

	UntagResource(*mgn.UntagResourceInput) (*mgn.UntagResourceOutput, error)
	UntagResourceWithContext(aws.Context, *mgn.UntagResourceInput, ...request.Option) (*mgn.UntagResourceOutput, error)
	UntagResourceRequest(*mgn.UntagResourceInput) (*request.Request, *mgn.UntagResourceOutput)

	UpdateApplication(*mgn.UpdateApplicationInput) (*mgn.UpdateApplicationOutput, error)
	UpdateApplicationWithContext(aws.Context, *mgn.UpdateApplicationInput, ...request.Option) (*mgn.UpdateApplicationOutput, error)
	UpdateApplicationRequest(*mgn.UpdateApplicationInput) (*request.Request, *mgn.UpdateApplicationOutput)

	UpdateLaunchConfiguration(*mgn.UpdateLaunchConfigurationInput) (*mgn.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationWithContext(aws.Context, *mgn.UpdateLaunchConfigurationInput, ...request.Option) (*mgn.UpdateLaunchConfigurationOutput, error)
	UpdateLaunchConfigurationRequest(*mgn.UpdateLaunchConfigurationInput) (*request.Request, *mgn.UpdateLaunchConfigurationOutput)

	UpdateLaunchConfigurationTemplate(*mgn.UpdateLaunchConfigurationTemplateInput) (*mgn.UpdateLaunchConfigurationTemplateOutput, error)
	UpdateLaunchConfigurationTemplateWithContext(aws.Context, *mgn.UpdateLaunchConfigurationTemplateInput, ...request.Option) (*mgn.UpdateLaunchConfigurationTemplateOutput, error)
	UpdateLaunchConfigurationTemplateRequest(*mgn.UpdateLaunchConfigurationTemplateInput) (*request.Request, *mgn.UpdateLaunchConfigurationTemplateOutput)

	UpdateReplicationConfiguration(*mgn.UpdateReplicationConfigurationInput) (*mgn.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationWithContext(aws.Context, *mgn.UpdateReplicationConfigurationInput, ...request.Option) (*mgn.UpdateReplicationConfigurationOutput, error)
	UpdateReplicationConfigurationRequest(*mgn.UpdateReplicationConfigurationInput) (*request.Request, *mgn.UpdateReplicationConfigurationOutput)

	UpdateReplicationConfigurationTemplate(*mgn.UpdateReplicationConfigurationTemplateInput) (*mgn.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateWithContext(aws.Context, *mgn.UpdateReplicationConfigurationTemplateInput, ...request.Option) (*mgn.UpdateReplicationConfigurationTemplateOutput, error)
	UpdateReplicationConfigurationTemplateRequest(*mgn.UpdateReplicationConfigurationTemplateInput) (*request.Request, *mgn.UpdateReplicationConfigurationTemplateOutput)

	UpdateSourceServerReplicationType(*mgn.UpdateSourceServerReplicationTypeInput) (*mgn.UpdateSourceServerReplicationTypeOutput, error)
	UpdateSourceServerReplicationTypeWithContext(aws.Context, *mgn.UpdateSourceServerReplicationTypeInput, ...request.Option) (*mgn.UpdateSourceServerReplicationTypeOutput, error)
	UpdateSourceServerReplicationTypeRequest(*mgn.UpdateSourceServerReplicationTypeInput) (*request.Request, *mgn.UpdateSourceServerReplicationTypeOutput)

	UpdateWave(*mgn.UpdateWaveInput) (*mgn.UpdateWaveOutput, error)
	UpdateWaveWithContext(aws.Context, *mgn.UpdateWaveInput, ...request.Option) (*mgn.UpdateWaveOutput, error)
	UpdateWaveRequest(*mgn.UpdateWaveInput) (*request.Request, *mgn.UpdateWaveOutput)
}

var _ MgnAPI = (*mgn.Mgn)(nil)

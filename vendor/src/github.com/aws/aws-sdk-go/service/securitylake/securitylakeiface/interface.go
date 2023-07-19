// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package securitylakeiface provides an interface to enable mocking the Amazon Security Lake service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package securitylakeiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/securitylake"
)

// SecurityLakeAPI provides an interface to enable mocking the
// securitylake.SecurityLake service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// Amazon Security Lake.
//	func myFunc(svc securitylakeiface.SecurityLakeAPI) bool {
//	    // Make svc.CreateAwsLogSource request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := securitylake.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockSecurityLakeClient struct {
//	    securitylakeiface.SecurityLakeAPI
//	}
//	func (m *mockSecurityLakeClient) CreateAwsLogSource(input *securitylake.CreateAwsLogSourceInput) (*securitylake.CreateAwsLogSourceOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockSecurityLakeClient{}
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
type SecurityLakeAPI interface {
	CreateAwsLogSource(*securitylake.CreateAwsLogSourceInput) (*securitylake.CreateAwsLogSourceOutput, error)
	CreateAwsLogSourceWithContext(aws.Context, *securitylake.CreateAwsLogSourceInput, ...request.Option) (*securitylake.CreateAwsLogSourceOutput, error)
	CreateAwsLogSourceRequest(*securitylake.CreateAwsLogSourceInput) (*request.Request, *securitylake.CreateAwsLogSourceOutput)

	CreateCustomLogSource(*securitylake.CreateCustomLogSourceInput) (*securitylake.CreateCustomLogSourceOutput, error)
	CreateCustomLogSourceWithContext(aws.Context, *securitylake.CreateCustomLogSourceInput, ...request.Option) (*securitylake.CreateCustomLogSourceOutput, error)
	CreateCustomLogSourceRequest(*securitylake.CreateCustomLogSourceInput) (*request.Request, *securitylake.CreateCustomLogSourceOutput)

	CreateDataLake(*securitylake.CreateDataLakeInput) (*securitylake.CreateDataLakeOutput, error)
	CreateDataLakeWithContext(aws.Context, *securitylake.CreateDataLakeInput, ...request.Option) (*securitylake.CreateDataLakeOutput, error)
	CreateDataLakeRequest(*securitylake.CreateDataLakeInput) (*request.Request, *securitylake.CreateDataLakeOutput)

	CreateDataLakeExceptionSubscription(*securitylake.CreateDataLakeExceptionSubscriptionInput) (*securitylake.CreateDataLakeExceptionSubscriptionOutput, error)
	CreateDataLakeExceptionSubscriptionWithContext(aws.Context, *securitylake.CreateDataLakeExceptionSubscriptionInput, ...request.Option) (*securitylake.CreateDataLakeExceptionSubscriptionOutput, error)
	CreateDataLakeExceptionSubscriptionRequest(*securitylake.CreateDataLakeExceptionSubscriptionInput) (*request.Request, *securitylake.CreateDataLakeExceptionSubscriptionOutput)

	CreateDataLakeOrganizationConfiguration(*securitylake.CreateDataLakeOrganizationConfigurationInput) (*securitylake.CreateDataLakeOrganizationConfigurationOutput, error)
	CreateDataLakeOrganizationConfigurationWithContext(aws.Context, *securitylake.CreateDataLakeOrganizationConfigurationInput, ...request.Option) (*securitylake.CreateDataLakeOrganizationConfigurationOutput, error)
	CreateDataLakeOrganizationConfigurationRequest(*securitylake.CreateDataLakeOrganizationConfigurationInput) (*request.Request, *securitylake.CreateDataLakeOrganizationConfigurationOutput)

	CreateSubscriber(*securitylake.CreateSubscriberInput) (*securitylake.CreateSubscriberOutput, error)
	CreateSubscriberWithContext(aws.Context, *securitylake.CreateSubscriberInput, ...request.Option) (*securitylake.CreateSubscriberOutput, error)
	CreateSubscriberRequest(*securitylake.CreateSubscriberInput) (*request.Request, *securitylake.CreateSubscriberOutput)

	CreateSubscriberNotification(*securitylake.CreateSubscriberNotificationInput) (*securitylake.CreateSubscriberNotificationOutput, error)
	CreateSubscriberNotificationWithContext(aws.Context, *securitylake.CreateSubscriberNotificationInput, ...request.Option) (*securitylake.CreateSubscriberNotificationOutput, error)
	CreateSubscriberNotificationRequest(*securitylake.CreateSubscriberNotificationInput) (*request.Request, *securitylake.CreateSubscriberNotificationOutput)

	DeleteAwsLogSource(*securitylake.DeleteAwsLogSourceInput) (*securitylake.DeleteAwsLogSourceOutput, error)
	DeleteAwsLogSourceWithContext(aws.Context, *securitylake.DeleteAwsLogSourceInput, ...request.Option) (*securitylake.DeleteAwsLogSourceOutput, error)
	DeleteAwsLogSourceRequest(*securitylake.DeleteAwsLogSourceInput) (*request.Request, *securitylake.DeleteAwsLogSourceOutput)

	DeleteCustomLogSource(*securitylake.DeleteCustomLogSourceInput) (*securitylake.DeleteCustomLogSourceOutput, error)
	DeleteCustomLogSourceWithContext(aws.Context, *securitylake.DeleteCustomLogSourceInput, ...request.Option) (*securitylake.DeleteCustomLogSourceOutput, error)
	DeleteCustomLogSourceRequest(*securitylake.DeleteCustomLogSourceInput) (*request.Request, *securitylake.DeleteCustomLogSourceOutput)

	DeleteDataLake(*securitylake.DeleteDataLakeInput) (*securitylake.DeleteDataLakeOutput, error)
	DeleteDataLakeWithContext(aws.Context, *securitylake.DeleteDataLakeInput, ...request.Option) (*securitylake.DeleteDataLakeOutput, error)
	DeleteDataLakeRequest(*securitylake.DeleteDataLakeInput) (*request.Request, *securitylake.DeleteDataLakeOutput)

	DeleteDataLakeExceptionSubscription(*securitylake.DeleteDataLakeExceptionSubscriptionInput) (*securitylake.DeleteDataLakeExceptionSubscriptionOutput, error)
	DeleteDataLakeExceptionSubscriptionWithContext(aws.Context, *securitylake.DeleteDataLakeExceptionSubscriptionInput, ...request.Option) (*securitylake.DeleteDataLakeExceptionSubscriptionOutput, error)
	DeleteDataLakeExceptionSubscriptionRequest(*securitylake.DeleteDataLakeExceptionSubscriptionInput) (*request.Request, *securitylake.DeleteDataLakeExceptionSubscriptionOutput)

	DeleteDataLakeOrganizationConfiguration(*securitylake.DeleteDataLakeOrganizationConfigurationInput) (*securitylake.DeleteDataLakeOrganizationConfigurationOutput, error)
	DeleteDataLakeOrganizationConfigurationWithContext(aws.Context, *securitylake.DeleteDataLakeOrganizationConfigurationInput, ...request.Option) (*securitylake.DeleteDataLakeOrganizationConfigurationOutput, error)
	DeleteDataLakeOrganizationConfigurationRequest(*securitylake.DeleteDataLakeOrganizationConfigurationInput) (*request.Request, *securitylake.DeleteDataLakeOrganizationConfigurationOutput)

	DeleteSubscriber(*securitylake.DeleteSubscriberInput) (*securitylake.DeleteSubscriberOutput, error)
	DeleteSubscriberWithContext(aws.Context, *securitylake.DeleteSubscriberInput, ...request.Option) (*securitylake.DeleteSubscriberOutput, error)
	DeleteSubscriberRequest(*securitylake.DeleteSubscriberInput) (*request.Request, *securitylake.DeleteSubscriberOutput)

	DeleteSubscriberNotification(*securitylake.DeleteSubscriberNotificationInput) (*securitylake.DeleteSubscriberNotificationOutput, error)
	DeleteSubscriberNotificationWithContext(aws.Context, *securitylake.DeleteSubscriberNotificationInput, ...request.Option) (*securitylake.DeleteSubscriberNotificationOutput, error)
	DeleteSubscriberNotificationRequest(*securitylake.DeleteSubscriberNotificationInput) (*request.Request, *securitylake.DeleteSubscriberNotificationOutput)

	DeregisterDataLakeDelegatedAdministrator(*securitylake.DeregisterDataLakeDelegatedAdministratorInput) (*securitylake.DeregisterDataLakeDelegatedAdministratorOutput, error)
	DeregisterDataLakeDelegatedAdministratorWithContext(aws.Context, *securitylake.DeregisterDataLakeDelegatedAdministratorInput, ...request.Option) (*securitylake.DeregisterDataLakeDelegatedAdministratorOutput, error)
	DeregisterDataLakeDelegatedAdministratorRequest(*securitylake.DeregisterDataLakeDelegatedAdministratorInput) (*request.Request, *securitylake.DeregisterDataLakeDelegatedAdministratorOutput)

	GetDataLakeExceptionSubscription(*securitylake.GetDataLakeExceptionSubscriptionInput) (*securitylake.GetDataLakeExceptionSubscriptionOutput, error)
	GetDataLakeExceptionSubscriptionWithContext(aws.Context, *securitylake.GetDataLakeExceptionSubscriptionInput, ...request.Option) (*securitylake.GetDataLakeExceptionSubscriptionOutput, error)
	GetDataLakeExceptionSubscriptionRequest(*securitylake.GetDataLakeExceptionSubscriptionInput) (*request.Request, *securitylake.GetDataLakeExceptionSubscriptionOutput)

	GetDataLakeOrganizationConfiguration(*securitylake.GetDataLakeOrganizationConfigurationInput) (*securitylake.GetDataLakeOrganizationConfigurationOutput, error)
	GetDataLakeOrganizationConfigurationWithContext(aws.Context, *securitylake.GetDataLakeOrganizationConfigurationInput, ...request.Option) (*securitylake.GetDataLakeOrganizationConfigurationOutput, error)
	GetDataLakeOrganizationConfigurationRequest(*securitylake.GetDataLakeOrganizationConfigurationInput) (*request.Request, *securitylake.GetDataLakeOrganizationConfigurationOutput)

	GetDataLakeSources(*securitylake.GetDataLakeSourcesInput) (*securitylake.GetDataLakeSourcesOutput, error)
	GetDataLakeSourcesWithContext(aws.Context, *securitylake.GetDataLakeSourcesInput, ...request.Option) (*securitylake.GetDataLakeSourcesOutput, error)
	GetDataLakeSourcesRequest(*securitylake.GetDataLakeSourcesInput) (*request.Request, *securitylake.GetDataLakeSourcesOutput)

	GetDataLakeSourcesPages(*securitylake.GetDataLakeSourcesInput, func(*securitylake.GetDataLakeSourcesOutput, bool) bool) error
	GetDataLakeSourcesPagesWithContext(aws.Context, *securitylake.GetDataLakeSourcesInput, func(*securitylake.GetDataLakeSourcesOutput, bool) bool, ...request.Option) error

	GetSubscriber(*securitylake.GetSubscriberInput) (*securitylake.GetSubscriberOutput, error)
	GetSubscriberWithContext(aws.Context, *securitylake.GetSubscriberInput, ...request.Option) (*securitylake.GetSubscriberOutput, error)
	GetSubscriberRequest(*securitylake.GetSubscriberInput) (*request.Request, *securitylake.GetSubscriberOutput)

	ListDataLakeExceptions(*securitylake.ListDataLakeExceptionsInput) (*securitylake.ListDataLakeExceptionsOutput, error)
	ListDataLakeExceptionsWithContext(aws.Context, *securitylake.ListDataLakeExceptionsInput, ...request.Option) (*securitylake.ListDataLakeExceptionsOutput, error)
	ListDataLakeExceptionsRequest(*securitylake.ListDataLakeExceptionsInput) (*request.Request, *securitylake.ListDataLakeExceptionsOutput)

	ListDataLakeExceptionsPages(*securitylake.ListDataLakeExceptionsInput, func(*securitylake.ListDataLakeExceptionsOutput, bool) bool) error
	ListDataLakeExceptionsPagesWithContext(aws.Context, *securitylake.ListDataLakeExceptionsInput, func(*securitylake.ListDataLakeExceptionsOutput, bool) bool, ...request.Option) error

	ListDataLakes(*securitylake.ListDataLakesInput) (*securitylake.ListDataLakesOutput, error)
	ListDataLakesWithContext(aws.Context, *securitylake.ListDataLakesInput, ...request.Option) (*securitylake.ListDataLakesOutput, error)
	ListDataLakesRequest(*securitylake.ListDataLakesInput) (*request.Request, *securitylake.ListDataLakesOutput)

	ListLogSources(*securitylake.ListLogSourcesInput) (*securitylake.ListLogSourcesOutput, error)
	ListLogSourcesWithContext(aws.Context, *securitylake.ListLogSourcesInput, ...request.Option) (*securitylake.ListLogSourcesOutput, error)
	ListLogSourcesRequest(*securitylake.ListLogSourcesInput) (*request.Request, *securitylake.ListLogSourcesOutput)

	ListLogSourcesPages(*securitylake.ListLogSourcesInput, func(*securitylake.ListLogSourcesOutput, bool) bool) error
	ListLogSourcesPagesWithContext(aws.Context, *securitylake.ListLogSourcesInput, func(*securitylake.ListLogSourcesOutput, bool) bool, ...request.Option) error

	ListSubscribers(*securitylake.ListSubscribersInput) (*securitylake.ListSubscribersOutput, error)
	ListSubscribersWithContext(aws.Context, *securitylake.ListSubscribersInput, ...request.Option) (*securitylake.ListSubscribersOutput, error)
	ListSubscribersRequest(*securitylake.ListSubscribersInput) (*request.Request, *securitylake.ListSubscribersOutput)

	ListSubscribersPages(*securitylake.ListSubscribersInput, func(*securitylake.ListSubscribersOutput, bool) bool) error
	ListSubscribersPagesWithContext(aws.Context, *securitylake.ListSubscribersInput, func(*securitylake.ListSubscribersOutput, bool) bool, ...request.Option) error

	RegisterDataLakeDelegatedAdministrator(*securitylake.RegisterDataLakeDelegatedAdministratorInput) (*securitylake.RegisterDataLakeDelegatedAdministratorOutput, error)
	RegisterDataLakeDelegatedAdministratorWithContext(aws.Context, *securitylake.RegisterDataLakeDelegatedAdministratorInput, ...request.Option) (*securitylake.RegisterDataLakeDelegatedAdministratorOutput, error)
	RegisterDataLakeDelegatedAdministratorRequest(*securitylake.RegisterDataLakeDelegatedAdministratorInput) (*request.Request, *securitylake.RegisterDataLakeDelegatedAdministratorOutput)

	UpdateDataLake(*securitylake.UpdateDataLakeInput) (*securitylake.UpdateDataLakeOutput, error)
	UpdateDataLakeWithContext(aws.Context, *securitylake.UpdateDataLakeInput, ...request.Option) (*securitylake.UpdateDataLakeOutput, error)
	UpdateDataLakeRequest(*securitylake.UpdateDataLakeInput) (*request.Request, *securitylake.UpdateDataLakeOutput)

	UpdateDataLakeExceptionSubscription(*securitylake.UpdateDataLakeExceptionSubscriptionInput) (*securitylake.UpdateDataLakeExceptionSubscriptionOutput, error)
	UpdateDataLakeExceptionSubscriptionWithContext(aws.Context, *securitylake.UpdateDataLakeExceptionSubscriptionInput, ...request.Option) (*securitylake.UpdateDataLakeExceptionSubscriptionOutput, error)
	UpdateDataLakeExceptionSubscriptionRequest(*securitylake.UpdateDataLakeExceptionSubscriptionInput) (*request.Request, *securitylake.UpdateDataLakeExceptionSubscriptionOutput)

	UpdateSubscriber(*securitylake.UpdateSubscriberInput) (*securitylake.UpdateSubscriberOutput, error)
	UpdateSubscriberWithContext(aws.Context, *securitylake.UpdateSubscriberInput, ...request.Option) (*securitylake.UpdateSubscriberOutput, error)
	UpdateSubscriberRequest(*securitylake.UpdateSubscriberInput) (*request.Request, *securitylake.UpdateSubscriberOutput)

	UpdateSubscriberNotification(*securitylake.UpdateSubscriberNotificationInput) (*securitylake.UpdateSubscriberNotificationOutput, error)
	UpdateSubscriberNotificationWithContext(aws.Context, *securitylake.UpdateSubscriberNotificationInput, ...request.Option) (*securitylake.UpdateSubscriberNotificationOutput, error)
	UpdateSubscriberNotificationRequest(*securitylake.UpdateSubscriberNotificationInput) (*request.Request, *securitylake.UpdateSubscriberNotificationOutput)
}

var _ SecurityLakeAPI = (*securitylake.SecurityLake)(nil)

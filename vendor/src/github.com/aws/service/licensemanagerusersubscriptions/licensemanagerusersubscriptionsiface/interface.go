// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package licensemanagerusersubscriptionsiface provides an interface to enable mocking the AWS License Manager User Subscriptions service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package licensemanagerusersubscriptionsiface

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/licensemanagerusersubscriptions"
)

// LicenseManagerUserSubscriptionsAPI provides an interface to enable mocking the
// licensemanagerusersubscriptions.LicenseManagerUserSubscriptions service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the SDK's request pipeline.
//
//	// myFunc uses an SDK service client to make a request to
//	// AWS License Manager User Subscriptions.
//	func myFunc(svc licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI) bool {
//	    // Make svc.AssociateUser request
//	}
//
//	func main() {
//	    sess := session.New()
//	    svc := licensemanagerusersubscriptions.New(sess)
//
//	    myFunc(svc)
//	}
//
// In your _test.go file:
//
//	// Define a mock struct to be used in your unit tests of myFunc.
//	type mockLicenseManagerUserSubscriptionsClient struct {
//	    licensemanagerusersubscriptionsiface.LicenseManagerUserSubscriptionsAPI
//	}
//	func (m *mockLicenseManagerUserSubscriptionsClient) AssociateUser(input *licensemanagerusersubscriptions.AssociateUserInput) (*licensemanagerusersubscriptions.AssociateUserOutput, error) {
//	    // mock response/functionality
//	}
//
//	func TestMyFunc(t *testing.T) {
//	    // Setup Test
//	    mockSvc := &mockLicenseManagerUserSubscriptionsClient{}
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
type LicenseManagerUserSubscriptionsAPI interface {
	AssociateUser(*licensemanagerusersubscriptions.AssociateUserInput) (*licensemanagerusersubscriptions.AssociateUserOutput, error)
	AssociateUserWithContext(aws.Context, *licensemanagerusersubscriptions.AssociateUserInput, ...request.Option) (*licensemanagerusersubscriptions.AssociateUserOutput, error)
	AssociateUserRequest(*licensemanagerusersubscriptions.AssociateUserInput) (*request.Request, *licensemanagerusersubscriptions.AssociateUserOutput)

	DeregisterIdentityProvider(*licensemanagerusersubscriptions.DeregisterIdentityProviderInput) (*licensemanagerusersubscriptions.DeregisterIdentityProviderOutput, error)
	DeregisterIdentityProviderWithContext(aws.Context, *licensemanagerusersubscriptions.DeregisterIdentityProviderInput, ...request.Option) (*licensemanagerusersubscriptions.DeregisterIdentityProviderOutput, error)
	DeregisterIdentityProviderRequest(*licensemanagerusersubscriptions.DeregisterIdentityProviderInput) (*request.Request, *licensemanagerusersubscriptions.DeregisterIdentityProviderOutput)

	DisassociateUser(*licensemanagerusersubscriptions.DisassociateUserInput) (*licensemanagerusersubscriptions.DisassociateUserOutput, error)
	DisassociateUserWithContext(aws.Context, *licensemanagerusersubscriptions.DisassociateUserInput, ...request.Option) (*licensemanagerusersubscriptions.DisassociateUserOutput, error)
	DisassociateUserRequest(*licensemanagerusersubscriptions.DisassociateUserInput) (*request.Request, *licensemanagerusersubscriptions.DisassociateUserOutput)

	ListIdentityProviders(*licensemanagerusersubscriptions.ListIdentityProvidersInput) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error)
	ListIdentityProvidersWithContext(aws.Context, *licensemanagerusersubscriptions.ListIdentityProvidersInput, ...request.Option) (*licensemanagerusersubscriptions.ListIdentityProvidersOutput, error)
	ListIdentityProvidersRequest(*licensemanagerusersubscriptions.ListIdentityProvidersInput) (*request.Request, *licensemanagerusersubscriptions.ListIdentityProvidersOutput)

	ListIdentityProvidersPages(*licensemanagerusersubscriptions.ListIdentityProvidersInput, func(*licensemanagerusersubscriptions.ListIdentityProvidersOutput, bool) bool) error
	ListIdentityProvidersPagesWithContext(aws.Context, *licensemanagerusersubscriptions.ListIdentityProvidersInput, func(*licensemanagerusersubscriptions.ListIdentityProvidersOutput, bool) bool, ...request.Option) error

	ListInstances(*licensemanagerusersubscriptions.ListInstancesInput) (*licensemanagerusersubscriptions.ListInstancesOutput, error)
	ListInstancesWithContext(aws.Context, *licensemanagerusersubscriptions.ListInstancesInput, ...request.Option) (*licensemanagerusersubscriptions.ListInstancesOutput, error)
	ListInstancesRequest(*licensemanagerusersubscriptions.ListInstancesInput) (*request.Request, *licensemanagerusersubscriptions.ListInstancesOutput)

	ListInstancesPages(*licensemanagerusersubscriptions.ListInstancesInput, func(*licensemanagerusersubscriptions.ListInstancesOutput, bool) bool) error
	ListInstancesPagesWithContext(aws.Context, *licensemanagerusersubscriptions.ListInstancesInput, func(*licensemanagerusersubscriptions.ListInstancesOutput, bool) bool, ...request.Option) error

	ListProductSubscriptions(*licensemanagerusersubscriptions.ListProductSubscriptionsInput) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error)
	ListProductSubscriptionsWithContext(aws.Context, *licensemanagerusersubscriptions.ListProductSubscriptionsInput, ...request.Option) (*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, error)
	ListProductSubscriptionsRequest(*licensemanagerusersubscriptions.ListProductSubscriptionsInput) (*request.Request, *licensemanagerusersubscriptions.ListProductSubscriptionsOutput)

	ListProductSubscriptionsPages(*licensemanagerusersubscriptions.ListProductSubscriptionsInput, func(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, bool) bool) error
	ListProductSubscriptionsPagesWithContext(aws.Context, *licensemanagerusersubscriptions.ListProductSubscriptionsInput, func(*licensemanagerusersubscriptions.ListProductSubscriptionsOutput, bool) bool, ...request.Option) error

	ListUserAssociations(*licensemanagerusersubscriptions.ListUserAssociationsInput) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error)
	ListUserAssociationsWithContext(aws.Context, *licensemanagerusersubscriptions.ListUserAssociationsInput, ...request.Option) (*licensemanagerusersubscriptions.ListUserAssociationsOutput, error)
	ListUserAssociationsRequest(*licensemanagerusersubscriptions.ListUserAssociationsInput) (*request.Request, *licensemanagerusersubscriptions.ListUserAssociationsOutput)

	ListUserAssociationsPages(*licensemanagerusersubscriptions.ListUserAssociationsInput, func(*licensemanagerusersubscriptions.ListUserAssociationsOutput, bool) bool) error
	ListUserAssociationsPagesWithContext(aws.Context, *licensemanagerusersubscriptions.ListUserAssociationsInput, func(*licensemanagerusersubscriptions.ListUserAssociationsOutput, bool) bool, ...request.Option) error

	RegisterIdentityProvider(*licensemanagerusersubscriptions.RegisterIdentityProviderInput) (*licensemanagerusersubscriptions.RegisterIdentityProviderOutput, error)
	RegisterIdentityProviderWithContext(aws.Context, *licensemanagerusersubscriptions.RegisterIdentityProviderInput, ...request.Option) (*licensemanagerusersubscriptions.RegisterIdentityProviderOutput, error)
	RegisterIdentityProviderRequest(*licensemanagerusersubscriptions.RegisterIdentityProviderInput) (*request.Request, *licensemanagerusersubscriptions.RegisterIdentityProviderOutput)

	StartProductSubscription(*licensemanagerusersubscriptions.StartProductSubscriptionInput) (*licensemanagerusersubscriptions.StartProductSubscriptionOutput, error)
	StartProductSubscriptionWithContext(aws.Context, *licensemanagerusersubscriptions.StartProductSubscriptionInput, ...request.Option) (*licensemanagerusersubscriptions.StartProductSubscriptionOutput, error)
	StartProductSubscriptionRequest(*licensemanagerusersubscriptions.StartProductSubscriptionInput) (*request.Request, *licensemanagerusersubscriptions.StartProductSubscriptionOutput)

	StopProductSubscription(*licensemanagerusersubscriptions.StopProductSubscriptionInput) (*licensemanagerusersubscriptions.StopProductSubscriptionOutput, error)
	StopProductSubscriptionWithContext(aws.Context, *licensemanagerusersubscriptions.StopProductSubscriptionInput, ...request.Option) (*licensemanagerusersubscriptions.StopProductSubscriptionOutput, error)
	StopProductSubscriptionRequest(*licensemanagerusersubscriptions.StopProductSubscriptionInput) (*request.Request, *licensemanagerusersubscriptions.StopProductSubscriptionOutput)

	UpdateIdentityProviderSettings(*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput) (*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput, error)
	UpdateIdentityProviderSettingsWithContext(aws.Context, *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput, ...request.Option) (*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput, error)
	UpdateIdentityProviderSettingsRequest(*licensemanagerusersubscriptions.UpdateIdentityProviderSettingsInput) (*request.Request, *licensemanagerusersubscriptions.UpdateIdentityProviderSettingsOutput)
}

var _ LicenseManagerUserSubscriptionsAPI = (*licensemanagerusersubscriptions.LicenseManagerUserSubscriptions)(nil)

// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package backupstorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/client"
	"github.com/aws/aws-sdk-go/aws/client/metadata"
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/aws/signer/v4"
	"github.com/aws/aws-sdk-go/private/protocol"
	"github.com/aws/aws-sdk-go/private/protocol/restjson"
)

// BackupStorage provides the API operation methods for making requests to
// AWS Backup Storage. See this package's package overview docs
// for details on the service.
//
// BackupStorage methods are safe to use concurrently. It is not safe to
// modify mutate any of the struct's properties though.
type BackupStorage struct {
	*client.Client
}

// Used for custom client initialization logic
var initClient func(*client.Client)

// Used for custom request initialization logic
var initRequest func(*request.Request)

// Service information constants
const (
	ServiceName = "BackupStorage" // Name of service.
	EndpointsID = "backupstorage" // ID to lookup a service endpoint with.
	ServiceID   = "BackupStorage" // ServiceID is a unique identifier of a specific service.
)

// New creates a new instance of the BackupStorage client with a session.
// If additional configuration is needed for the client instance use the optional
// aws.Config parameter to add your extra config.
//
// Example:
//
//	mySession := session.Must(session.NewSession())
//
//	// Create a BackupStorage client from just a session.
//	svc := backupstorage.New(mySession)
//
//	// Create a BackupStorage client with additional configuration
//	svc := backupstorage.New(mySession, aws.NewConfig().WithRegion("us-west-2"))
func New(p client.ConfigProvider, cfgs ...*aws.Config) *BackupStorage {
	c := p.ClientConfig(EndpointsID, cfgs...)
	if c.SigningNameDerived || len(c.SigningName) == 0 {
		c.SigningName = "backup-storage"
	}
	return newClient(*c.Config, c.Handlers, c.PartitionID, c.Endpoint, c.SigningRegion, c.SigningName, c.ResolvedRegion)
}

// newClient creates, initializes and returns a new service client instance.
func newClient(cfg aws.Config, handlers request.Handlers, partitionID, endpoint, signingRegion, signingName, resolvedRegion string) *BackupStorage {
	svc := &BackupStorage{
		Client: client.New(
			cfg,
			metadata.ClientInfo{
				ServiceName:    ServiceName,
				ServiceID:      ServiceID,
				SigningName:    signingName,
				SigningRegion:  signingRegion,
				PartitionID:    partitionID,
				Endpoint:       endpoint,
				APIVersion:     "2018-04-10",
				ResolvedRegion: resolvedRegion,
			},
			handlers,
		),
	}

	// Handlers
	svc.Handlers.Sign.PushBackNamed(v4.SignRequestHandler)
	svc.Handlers.Build.PushBackNamed(restjson.BuildHandler)
	svc.Handlers.Unmarshal.PushBackNamed(restjson.UnmarshalHandler)
	svc.Handlers.UnmarshalMeta.PushBackNamed(restjson.UnmarshalMetaHandler)
	svc.Handlers.UnmarshalError.PushBackNamed(
		protocol.NewUnmarshalErrorHandler(restjson.NewUnmarshalTypedError(exceptionFromCode)).NamedHandler(),
	)

	// Run custom client initialization if present
	if initClient != nil {
		initClient(svc.Client)
	}

	return svc
}

// newRequest creates a new request for a BackupStorage operation and runs any
// custom request initialization.
func (c *BackupStorage) newRequest(op *request.Operation, params, data interface{}) *request.Request {
	req := c.NewRequest(op, params, data)

	// Run custom request initialization if present
	if initRequest != nil {
		initRequest(req)
	}

	return req
}

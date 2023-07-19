// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

// Package ivsrealtime provides the client and types for making API
// requests to Amazon Interactive Video Service RealTime.
//
// # Introduction
//
// The Amazon Interactive Video Service (IVS) stage API is REST compatible,
// using a standard HTTP API and an AWS EventBridge event stream for responses.
// JSON is used for both requests and responses, including errors.
//
// Terminology:
//
//   - The IVS stage API sometimes is referred to as the IVS RealTime API.
//
//   - A participant token is an authorization token used to publish/subscribe
//     to a stage.
//
//   - A participant object represents participants (people) in the stage and
//     contains information about them. When a token is created, it includes
//     a participant ID; when a participant uses that token to join a stage,
//     the participant is associated with that participant ID There is a 1:1
//     mapping between participant tokens and participants.
//
// # Resources
//
// The following resources contain information about your IVS live stream (see
// Getting Started with Amazon IVS (https://docs.aws.amazon.com/ivs/latest/userguide/getting-started.html)):
//
//   - Stage — A stage is a virtual space where multiple participants can
//     exchange audio and video in real time.
//
// # Tagging
//
// A tag is a metadata label that you assign to an AWS resource. A tag comprises
// a key and a value, both set by you. For example, you might set a tag as topic:nature
// to label a particular video category. See Tagging AWS Resources (https://docs.aws.amazon.com/general/latest/gr/aws_tagging.html)
// for more information, including restrictions that apply to tags and "Tag
// naming limits and requirements"; Amazon IVS stages has no service-specific
// constraints beyond what is documented there.
//
// Tags can help you identify and organize your AWS resources. For example,
// you can use the same tag for different resources to indicate that they are
// related. You can also use tags to manage access (see Access Tags (https://docs.aws.amazon.com/IAM/latest/UserGuide/access_tags.html)).
//
// The Amazon IVS stage API has these tag-related endpoints: TagResource, UntagResource,
// and ListTagsForResource. The following resource supports tagging: Stage.
//
// At most 50 tags can be applied to a resource.
//
// Stages Endpoints
//
//   - CreateParticipantToken — Creates an additional token for a specified
//     stage. This can be done after stage creation or when tokens expire.
//
//   - CreateStage — Creates a new stage (and optionally participant tokens).
//
//   - DeleteStage — Shuts down and deletes the specified stage (disconnecting
//     all participants).
//
//   - DisconnectParticipant — Disconnects a specified participant and revokes
//     the participant permanently from a specified stage.
//
//   - GetParticipant — Gets information about the specified participant
//     token.
//
//   - GetStage — Gets information for the specified stage.
//
//   - GetStageSession — Gets information for the specified stage session.
//
//   - ListParticipantEvents — Lists events for a specified participant that
//     occurred during a specified stage session.
//
//   - ListParticipants — Lists all participants in a specified stage session.
//
//   - ListStages — Gets summary information about all stages in your account,
//     in the AWS region where the API request is processed.
//
//   - ListStageSessions — Gets all sessions for a specified stage.
//
//   - UpdateStage — Updates a stage’s configuration.
//
// Tags Endpoints
//
//   - ListTagsForResource — Gets information about AWS tags for the specified
//     ARN.
//
//   - TagResource — Adds or updates tags for the AWS resource with the specified
//     ARN.
//
//   - UntagResource — Removes tags from the resource with the specified
//     ARN.
//
// See https://docs.aws.amazon.com/goto/WebAPI/ivs-realtime-2020-07-14 for more information on this service.
//
// See ivsrealtime package documentation for more information.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ivsrealtime/
//
// # Using the Client
//
// To contact Amazon Interactive Video Service RealTime with the SDK use the New function to create
// a new service client. With that client you can make API requests to the service.
// These clients are safe to use concurrently.
//
// See the SDK's documentation for more information on how to use the SDK.
// https://docs.aws.amazon.com/sdk-for-go/api/
//
// See aws.Config documentation for more information on configuring SDK clients.
// https://docs.aws.amazon.com/sdk-for-go/api/aws/#Config
//
// See the Amazon Interactive Video Service RealTime client IVSRealTime for more
// information on creating client for this service.
// https://docs.aws.amazon.com/sdk-for-go/api/service/ivsrealtime/#New
package ivsrealtime

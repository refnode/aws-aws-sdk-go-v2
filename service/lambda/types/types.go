// Code generated by smithy-go-codegen DO NOT EDIT.

package types

import (
	smithydocument "github.com/aws/smithy-go/document"
	"time"
)

// Limits that are related to concurrency and storage. All file and storage sizes
// are in bytes.
type AccountLimit struct {

	// The maximum size of a function's deployment package and layers when they're
	// extracted.
	CodeSizeUnzipped int64

	// The maximum size of a deployment package when it's uploaded directly to Lambda.
	// Use Amazon S3 for larger files.
	CodeSizeZipped int64

	// The maximum number of simultaneous function executions.
	ConcurrentExecutions int32

	// The amount of storage space that you can use for all deployment packages and
	// layer archives.
	TotalCodeSize int64

	// The maximum number of simultaneous function executions, minus the capacity
	// that's reserved for individual functions with PutFunctionConcurrency .
	UnreservedConcurrentExecutions *int32

	noSmithyDocumentSerde
}

// The number of functions and amount of storage in use.
type AccountUsage struct {

	// The number of Lambda functions.
	FunctionCount int64

	// The amount of storage space, in bytes, that's being used by deployment packages
	// and layer archives.
	TotalCodeSize int64

	noSmithyDocumentSerde
}

// Provides configuration information about a Lambda function alias (https://docs.aws.amazon.com/lambda/latest/dg/configuration-aliases.html)
// .
type AliasConfiguration struct {

	// The Amazon Resource Name (ARN) of the alias.
	AliasArn *string

	// A description of the alias.
	Description *string

	// The function version that the alias invokes.
	FunctionVersion *string

	// The name of the alias.
	Name *string

	// A unique identifier that changes when you update the alias.
	RevisionId *string

	// The routing configuration (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
	// of the alias.
	RoutingConfig *AliasRoutingConfiguration

	noSmithyDocumentSerde
}

// The traffic-shifting (https://docs.aws.amazon.com/lambda/latest/dg/lambda-traffic-shifting-using-aliases.html)
// configuration of a Lambda function alias.
type AliasRoutingConfiguration struct {

	// The second version, and the percentage of traffic that's routed to it.
	AdditionalVersionWeights map[string]float64

	noSmithyDocumentSerde
}

// List of signing profiles that can sign a code package.
type AllowedPublishers struct {

	// The Amazon Resource Name (ARN) for each of the signing profiles. A signing
	// profile defines a trusted user who can sign a code package.
	//
	// This member is required.
	SigningProfileVersionArns []string

	noSmithyDocumentSerde
}

// Specific configuration settings for an Amazon Managed Streaming for Apache
// Kafka (Amazon MSK) event source.
type AmazonManagedKafkaEventSourceConfig struct {

	// The identifier for the Kafka consumer group to join. The consumer group ID must
	// be unique among all your Kafka event sources. After creating a Kafka event
	// source mapping with the consumer group ID specified, you cannot update this
	// value. For more information, see Customizable consumer group ID (https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id)
	// .
	ConsumerGroupId *string

	noSmithyDocumentSerde
}

// Details about a Code signing configuration (https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html)
// .
type CodeSigningConfig struct {

	// List of allowed publishers.
	//
	// This member is required.
	AllowedPublishers *AllowedPublishers

	// The Amazon Resource Name (ARN) of the Code signing configuration.
	//
	// This member is required.
	CodeSigningConfigArn *string

	// Unique identifer for the Code signing configuration.
	//
	// This member is required.
	CodeSigningConfigId *string

	// The code signing policy controls the validation failure action for signature
	// mismatch or expiry.
	//
	// This member is required.
	CodeSigningPolicies *CodeSigningPolicies

	// The date and time that the Code signing configuration was last modified, in
	// ISO-8601 format (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// This member is required.
	LastModified *string

	// Code signing configuration description.
	Description *string

	noSmithyDocumentSerde
}

// Code signing configuration policies (https://docs.aws.amazon.com/lambda/latest/dg/configuration-codesigning.html#config-codesigning-policies)
// specify the validation failure action for signature mismatch or expiry.
type CodeSigningPolicies struct {

	// Code signing configuration policy for deployment validation failure. If you set
	// the policy to Enforce , Lambda blocks the deployment request if signature
	// validation checks fail. If you set the policy to Warn , Lambda allows the
	// deployment and creates a CloudWatch log. Default value: Warn
	UntrustedArtifactOnDeployment CodeSigningPolicy

	noSmithyDocumentSerde
}

type Concurrency struct {

	// The number of concurrent executions that are reserved for this function. For
	// more information, see Managing Lambda reserved concurrency (https://docs.aws.amazon.com/lambda/latest/dg/configuration-concurrency.html)
	// .
	ReservedConcurrentExecutions *int32

	noSmithyDocumentSerde
}

// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
// settings for your Lambda function URL. Use CORS to grant access to your function
// URL from any origin. You can also use CORS to control access for specific HTTP
// headers and methods in requests to your function URL.
type Cors struct {

	// Whether to allow cookies or other credentials in requests to your function URL.
	// The default is false .
	AllowCredentials *bool

	// The HTTP headers that origins can include in requests to your function URL. For
	// example: Date , Keep-Alive , X-Custom-Header .
	AllowHeaders []string

	// The HTTP methods that are allowed when calling your function URL. For example:
	// GET , POST , DELETE , or the wildcard character ( * ).
	AllowMethods []string

	// The origins that can access your function URL. You can list any number of
	// specific origins, separated by a comma. For example: https://www.example.com ,
	// http://localhost:60905 . Alternatively, you can grant access to all origins
	// using the wildcard character ( * ).
	AllowOrigins []string

	// The HTTP headers in your function response that you want to expose to origins
	// that call your function URL. For example: Date , Keep-Alive , X-Custom-Header .
	ExposeHeaders []string

	// The maximum amount of time, in seconds, that web browsers can cache results of
	// a preflight request. By default, this is set to 0 , which means that the browser
	// doesn't cache results.
	MaxAge *int32

	noSmithyDocumentSerde
}

// The dead-letter queue (https://docs.aws.amazon.com/lambda/latest/dg/invocation-async.html#dlq)
// for failed asynchronous invocations.
type DeadLetterConfig struct {

	// The Amazon Resource Name (ARN) of an Amazon SQS queue or Amazon SNS topic.
	TargetArn *string

	noSmithyDocumentSerde
}

// A configuration object that specifies the destination of an event after Lambda
// processes it.
type DestinationConfig struct {

	// The destination configuration for failed invocations.
	OnFailure *OnFailure

	// The destination configuration for successful invocations.
	OnSuccess *OnSuccess

	noSmithyDocumentSerde
}

// Specific configuration settings for a DocumentDB event source.
type DocumentDBEventSourceConfig struct {

	// The name of the collection to consume within the database. If you do not
	// specify a collection, Lambda consumes all collections.
	CollectionName *string

	// The name of the database to consume within the DocumentDB cluster.
	DatabaseName *string

	// Determines what DocumentDB sends to your event stream during document update
	// operations. If set to UpdateLookup, DocumentDB sends a delta describing the
	// changes, along with a copy of the entire document. Otherwise, DocumentDB sends
	// only a partial document that contains the changes.
	FullDocument FullDocument

	noSmithyDocumentSerde
}

// A function's environment variable settings. You can use environment variables
// to adjust your function's behavior without updating code. An environment
// variable is a pair of strings that are stored in a function's version-specific
// configuration.
type Environment struct {

	// Environment variable key-value pairs. For more information, see Using Lambda
	// environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html)
	// .
	Variables map[string]string

	noSmithyDocumentSerde
}

// Error messages for environment variables that couldn't be applied.
type EnvironmentError struct {

	// The error code.
	ErrorCode *string

	// The error message.
	Message *string

	noSmithyDocumentSerde
}

// The results of an operation to update or read environment variables. If the
// operation succeeds, the response contains the environment variables. If it
// fails, the response contains details about the error.
type EnvironmentResponse struct {

	// Error messages for environment variables that couldn't be applied.
	Error *EnvironmentError

	// Environment variable key-value pairs. Omitted from CloudTrail logs.
	Variables map[string]string

	noSmithyDocumentSerde
}

// The size of the function's /tmp directory in MB. The default value is 512, but
// it can be any whole number between 512 and 10,240 MB.
type EphemeralStorage struct {

	// The size of the function's /tmp directory.
	//
	// This member is required.
	Size *int32

	noSmithyDocumentSerde
}

// A mapping between an Amazon Web Services resource and a Lambda function. For
// details, see CreateEventSourceMapping .
type EventSourceMappingConfiguration struct {

	// Specific configuration settings for an Amazon Managed Streaming for Apache
	// Kafka (Amazon MSK) event source.
	AmazonManagedKafkaEventSourceConfig *AmazonManagedKafkaEventSourceConfig

	// The maximum number of records in each batch that Lambda pulls from your stream
	// or queue and sends to your function. Lambda passes all of the records in the
	// batch to the function in a single call, up to the payload limit for synchronous
	// invocation (6 MB). Default value: Varies by service. For Amazon SQS, the default
	// is 10. For all other services, the default is 100. Related setting: When you set
	// BatchSize to a value greater than 10, you must set
	// MaximumBatchingWindowInSeconds to at least 1.
	BatchSize *int32

	// (Kinesis and DynamoDB Streams only) If the function returns an error, split the
	// batch in two and retry. The default value is false.
	BisectBatchOnFunctionError *bool

	// (Kinesis and DynamoDB Streams only) An Amazon SQS queue or Amazon SNS topic
	// destination for discarded records.
	DestinationConfig *DestinationConfig

	// Specific configuration settings for a DocumentDB event source.
	DocumentDBEventSourceConfig *DocumentDBEventSourceConfig

	// The Amazon Resource Name (ARN) of the event source.
	EventSourceArn *string

	// An object that defines the filter criteria that determine whether Lambda should
	// process an event. For more information, see Lambda event filtering (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html)
	// .
	FilterCriteria *FilterCriteria

	// The ARN of the Lambda function.
	FunctionArn *string

	// (Kinesis, DynamoDB Streams, and Amazon SQS) A list of current response type
	// enums applied to the event source mapping.
	FunctionResponseTypes []FunctionResponseType

	// The date that the event source mapping was last updated or that its state
	// changed.
	LastModified *time.Time

	// The result of the last Lambda invocation of your function.
	LastProcessingResult *string

	// The maximum amount of time, in seconds, that Lambda spends gathering records
	// before invoking the function. You can configure MaximumBatchingWindowInSeconds
	// to any value from 0 seconds to 300 seconds in increments of seconds. For streams
	// and Amazon SQS event sources, the default batching window is 0 seconds. For
	// Amazon MSK, Self-managed Apache Kafka, Amazon MQ, and DocumentDB event sources,
	// the default batching window is 500 ms. Note that because you can only change
	// MaximumBatchingWindowInSeconds in increments of seconds, you cannot revert back
	// to the 500 ms default batching window after you have changed it. To restore the
	// default batching window, you must create a new event source mapping. Related
	// setting: For streams and Amazon SQS event sources, when you set BatchSize to a
	// value greater than 10, you must set MaximumBatchingWindowInSeconds to at least
	// 1.
	MaximumBatchingWindowInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records older than the specified
	// age. The default value is -1, which sets the maximum age to infinite. When the
	// value is set to infinite, Lambda never discards old records. The minimum valid
	// value for maximum record age is 60s. Although values less than 60 and greater
	// than -1 fall within the parameter's absolute range, they are not allowed
	MaximumRecordAgeInSeconds *int32

	// (Kinesis and DynamoDB Streams only) Discard records after the specified number
	// of retries. The default value is -1, which sets the maximum number of retries to
	// infinite. When MaximumRetryAttempts is infinite, Lambda retries failed records
	// until the record expires in the event source.
	MaximumRetryAttempts *int32

	// (Kinesis and DynamoDB Streams only) The number of batches to process
	// concurrently from each shard. The default value is 1.
	ParallelizationFactor *int32

	// (Amazon MQ) The name of the Amazon MQ broker destination queue to consume.
	Queues []string

	// (Amazon SQS only) The scaling configuration for the event source. For more
	// information, see Configuring maximum concurrency for Amazon SQS event sources (https://docs.aws.amazon.com/lambda/latest/dg/with-sqs.html#events-sqs-max-concurrency)
	// .
	ScalingConfig *ScalingConfig

	// The self-managed Apache Kafka cluster for your event source.
	SelfManagedEventSource *SelfManagedEventSource

	// Specific configuration settings for a self-managed Apache Kafka event source.
	SelfManagedKafkaEventSourceConfig *SelfManagedKafkaEventSourceConfig

	// An array of the authentication protocol, VPC components, or virtual host to
	// secure and define your event source.
	SourceAccessConfigurations []SourceAccessConfiguration

	// The position in a stream from which to start reading. Required for Amazon
	// Kinesis and Amazon DynamoDB Stream event sources. AT_TIMESTAMP is supported
	// only for Amazon Kinesis streams, Amazon DocumentDB, Amazon MSK, and self-managed
	// Apache Kafka.
	StartingPosition EventSourcePosition

	// With StartingPosition set to AT_TIMESTAMP , the time from which to start
	// reading. StartingPositionTimestamp cannot be in the future.
	StartingPositionTimestamp *time.Time

	// The state of the event source mapping. It can be one of the following: Creating
	// , Enabling , Enabled , Disabling , Disabled , Updating , or Deleting .
	State *string

	// Indicates whether a user or Lambda made the last change to the event source
	// mapping.
	StateTransitionReason *string

	// The name of the Kafka topic.
	Topics []string

	// (Kinesis and DynamoDB Streams only) The duration in seconds of a processing
	// window for DynamoDB and Kinesis Streams event sources. A value of 0 seconds
	// indicates no tumbling window.
	TumblingWindowInSeconds *int32

	// The identifier of the event source mapping.
	UUID *string

	noSmithyDocumentSerde
}

// Details about the connection between a Lambda function and an Amazon EFS file
// system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html)
// .
type FileSystemConfig struct {

	// The Amazon Resource Name (ARN) of the Amazon EFS access point that provides
	// access to the file system.
	//
	// This member is required.
	Arn *string

	// The path where the function can access the file system, starting with /mnt/ .
	//
	// This member is required.
	LocalMountPath *string

	noSmithyDocumentSerde
}

// A structure within a FilterCriteria object that defines an event filtering
// pattern.
type Filter struct {

	// A filter pattern. For more information on the syntax of a filter pattern, see
	// Filter rule syntax (https://docs.aws.amazon.com/lambda/latest/dg/invocation-eventfiltering.html#filtering-syntax)
	// .
	Pattern *string

	noSmithyDocumentSerde
}

// An object that contains the filters for an event source.
type FilterCriteria struct {

	// A list of filters.
	Filters []Filter

	noSmithyDocumentSerde
}

// The code for the Lambda function. You can either specify an object in Amazon
// S3, upload a .zip file archive deployment package directly, or specify the URI
// of a container image.
type FunctionCode struct {

	// URI of a container image (https://docs.aws.amazon.com/lambda/latest/dg/lambda-images.html)
	// in the Amazon ECR registry.
	ImageUri *string

	// An Amazon S3 bucket in the same Amazon Web Services Region as your function.
	// The bucket can be in a different Amazon Web Services account.
	S3Bucket *string

	// The Amazon S3 key of the deployment package.
	S3Key *string

	// For versioned objects, the version of the deployment package object to use.
	S3ObjectVersion *string

	// The base64-encoded contents of the deployment package. Amazon Web Services SDK
	// and CLI clients handle the encoding for you.
	ZipFile []byte

	noSmithyDocumentSerde
}

// Details about a function's deployment package.
type FunctionCodeLocation struct {

	// URI of a container image in the Amazon ECR registry.
	ImageUri *string

	// A presigned URL that you can use to download the deployment package.
	Location *string

	// The service that's hosting the file.
	RepositoryType *string

	// The resolved URI for the image.
	ResolvedImageUri *string

	noSmithyDocumentSerde
}

// Details about a function's configuration.
type FunctionConfiguration struct {

	// The instruction set architecture that the function supports. Architecture is a
	// string array with one of the valid values. The default architecture value is
	// x86_64 .
	Architectures []Architecture

	// The SHA256 hash of the function's deployment package.
	CodeSha256 *string

	// The size of the function's deployment package, in bytes.
	CodeSize int64

	// The function's dead letter queue.
	DeadLetterConfig *DeadLetterConfig

	// The function's description.
	Description *string

	// The function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html)
	// . Omitted from CloudTrail logs.
	Environment *EnvironmentResponse

	// The size of the function’s /tmp directory in MB. The default value is 512, but
	// it can be any whole number between 512 and 10,240 MB.
	EphemeralStorage *EphemeralStorage

	// Connection settings for an Amazon EFS file system (https://docs.aws.amazon.com/lambda/latest/dg/configuration-filesystem.html)
	// .
	FileSystemConfigs []FileSystemConfig

	// The function's Amazon Resource Name (ARN).
	FunctionArn *string

	// The name of the function.
	FunctionName *string

	// The function that Lambda calls to begin running your function.
	Handler *string

	// The function's image configuration values.
	ImageConfigResponse *ImageConfigResponse

	// The KMS key that's used to encrypt the function's environment variables (https://docs.aws.amazon.com/lambda/latest/dg/configuration-envvars.html#configuration-envvars-encryption)
	// . When Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart-security.html)
	// is activated, this key is also used to encrypt the function's snapshot. This key
	// is returned only if you've configured a customer managed key.
	KMSKeyArn *string

	// The date and time that the function was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	LastModified *string

	// The status of the last update that was performed on the function. This is first
	// set to Successful after function creation completes.
	LastUpdateStatus LastUpdateStatus

	// The reason for the last update that was performed on the function.
	LastUpdateStatusReason *string

	// The reason code for the last update that was performed on the function.
	LastUpdateStatusReasonCode LastUpdateStatusReasonCode

	// The function's layers (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
	// .
	Layers []Layer

	// For Lambda@Edge functions, the ARN of the main function.
	MasterArn *string

	// The amount of memory available to the function at runtime.
	MemorySize *int32

	// The type of deployment package. Set to Image for container image and set Zip
	// for .zip file archive.
	PackageType PackageType

	// The latest updated revision of the function or alias.
	RevisionId *string

	// The function's execution role.
	Role *string

	// The identifier of the function's runtime (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html)
	// . Runtime is required if the deployment package is a .zip file archive. The
	// following list includes deprecated runtimes. For more information, see Runtime
	// deprecation policy (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy)
	// .
	Runtime Runtime

	// The ARN of the runtime and any errors that occured.
	RuntimeVersionConfig *RuntimeVersionConfig

	// The ARN of the signing job.
	SigningJobArn *string

	// The ARN of the signing profile version.
	SigningProfileVersionArn *string

	// Set ApplyOn to PublishedVersions to create a snapshot of the initialized
	// execution environment when you publish a function version. For more information,
	// see Improving startup performance with Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
	// .
	SnapStart *SnapStartResponse

	// The current state of the function. When the state is Inactive , you can
	// reactivate the function by invoking it.
	State State

	// The reason for the function's current state.
	StateReason *string

	// The reason code for the function's current state. When the code is Creating ,
	// you can't invoke or modify the function.
	StateReasonCode StateReasonCode

	// The amount of time in seconds that Lambda allows a function to run before
	// stopping it.
	Timeout *int32

	// The function's X-Ray tracing configuration.
	TracingConfig *TracingConfigResponse

	// The version of the Lambda function.
	Version *string

	// The function's networking configuration.
	VpcConfig *VpcConfigResponse

	noSmithyDocumentSerde
}

type FunctionEventInvokeConfig struct {

	// A destination for events after they have been sent to a function for
	// processing. Destinations
	//   - Function - The Amazon Resource Name (ARN) of a Lambda function.
	//   - Queue - The ARN of a standard SQS queue.
	//   - Topic - The ARN of a standard SNS topic.
	//   - Event Bus - The ARN of an Amazon EventBridge event bus.
	DestinationConfig *DestinationConfig

	// The Amazon Resource Name (ARN) of the function.
	FunctionArn *string

	// The date and time that the configuration was last updated.
	LastModified *time.Time

	// The maximum age of a request that Lambda sends to a function for processing.
	MaximumEventAgeInSeconds *int32

	// The maximum number of times to retry when the function returns an error.
	MaximumRetryAttempts *int32

	noSmithyDocumentSerde
}

// Details about a Lambda function URL.
type FunctionUrlConfig struct {

	// The type of authentication that your function URL uses. Set to AWS_IAM if you
	// want to restrict access to authenticated users only. Set to NONE if you want to
	// bypass IAM authentication to create a public endpoint. For more information, see
	// Security and auth model for Lambda function URLs (https://docs.aws.amazon.com/lambda/latest/dg/urls-auth.html)
	// .
	//
	// This member is required.
	AuthType FunctionUrlAuthType

	// When the function URL was created, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// This member is required.
	CreationTime *string

	// The Amazon Resource Name (ARN) of your function.
	//
	// This member is required.
	FunctionArn *string

	// The HTTP URL endpoint for your function.
	//
	// This member is required.
	FunctionUrl *string

	// When the function URL configuration was last updated, in ISO-8601 format (https://www.w3.org/TR/NOTE-datetime)
	// (YYYY-MM-DDThh:mm:ss.sTZD).
	//
	// This member is required.
	LastModifiedTime *string

	// The cross-origin resource sharing (CORS) (https://developer.mozilla.org/en-US/docs/Web/HTTP/CORS)
	// settings for your function URL.
	Cors *Cors

	// Use one of the following options:
	//   - BUFFERED – This is the default option. Lambda invokes your function using
	//   the Invoke API operation. Invocation results are available when the payload is
	//   complete. The maximum payload size is 6 MB.
	//   - RESPONSE_STREAM – Your function streams payload results as they become
	//   available. Lambda invokes your function using the InvokeWithResponseStream API
	//   operation. The maximum response payload size is 20 MB, however, you can
	//   request a quota increase (https://docs.aws.amazon.com/servicequotas/latest/userguide/request-quota-increase.html)
	//   .
	InvokeMode InvokeMode

	noSmithyDocumentSerde
}

// Configuration values that override the container image Dockerfile settings. For
// more information, see Container image settings (https://docs.aws.amazon.com/lambda/latest/dg/images-create.html#images-parms)
// .
type ImageConfig struct {

	// Specifies parameters that you want to pass in with ENTRYPOINT.
	Command []string

	// Specifies the entry point to their application, which is typically the location
	// of the runtime executable.
	EntryPoint []string

	// Specifies the working directory.
	WorkingDirectory *string

	noSmithyDocumentSerde
}

// Error response to GetFunctionConfiguration .
type ImageConfigError struct {

	// Error code.
	ErrorCode *string

	// Error message.
	Message *string

	noSmithyDocumentSerde
}

// Response to a GetFunctionConfiguration request.
type ImageConfigResponse struct {

	// Error response to GetFunctionConfiguration .
	Error *ImageConfigError

	// Configuration values that override the container image Dockerfile.
	ImageConfig *ImageConfig

	noSmithyDocumentSerde
}

// A chunk of the streamed response payload.
type InvokeResponseStreamUpdate struct {

	// Data returned by your Lambda function.
	Payload []byte

	noSmithyDocumentSerde
}

// A response confirming that the event stream is complete.
type InvokeWithResponseStreamCompleteEvent struct {

	// An error code.
	ErrorCode *string

	// The details of any returned error.
	ErrorDetails *string

	// The last 4 KB of the execution log, which is base64-encoded.
	LogResult *string

	noSmithyDocumentSerde
}

// An object that includes a chunk of the response payload. When the stream has
// ended, Lambda includes a InvokeComplete object.
//
// The following types satisfy this interface:
//
//	InvokeWithResponseStreamResponseEventMemberInvokeComplete
//	InvokeWithResponseStreamResponseEventMemberPayloadChunk
type InvokeWithResponseStreamResponseEvent interface {
	isInvokeWithResponseStreamResponseEvent()
}

// An object that's returned when the stream has ended and all the payload chunks
// have been returned.
type InvokeWithResponseStreamResponseEventMemberInvokeComplete struct {
	Value InvokeWithResponseStreamCompleteEvent

	noSmithyDocumentSerde
}

func (*InvokeWithResponseStreamResponseEventMemberInvokeComplete) isInvokeWithResponseStreamResponseEvent() {
}

// A chunk of the streamed response payload.
type InvokeWithResponseStreamResponseEventMemberPayloadChunk struct {
	Value InvokeResponseStreamUpdate

	noSmithyDocumentSerde
}

func (*InvokeWithResponseStreamResponseEventMemberPayloadChunk) isInvokeWithResponseStreamResponseEvent() {
}

// An Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// .
type Layer struct {

	// The Amazon Resource Name (ARN) of the function layer.
	Arn *string

	// The size of the layer archive in bytes.
	CodeSize int64

	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn *string

	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn *string

	noSmithyDocumentSerde
}

// Details about an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// .
type LayersListItem struct {

	// The newest version of the layer.
	LatestMatchingVersion *LayerVersionsListItem

	// The Amazon Resource Name (ARN) of the function layer.
	LayerArn *string

	// The name of the layer.
	LayerName *string

	noSmithyDocumentSerde
}

// A ZIP archive that contains the contents of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// . You can specify either an Amazon S3 location, or upload a layer archive
// directly.
type LayerVersionContentInput struct {

	// The Amazon S3 bucket of the layer archive.
	S3Bucket *string

	// The Amazon S3 key of the layer archive.
	S3Key *string

	// For versioned objects, the version of the layer archive object to use.
	S3ObjectVersion *string

	// The base64-encoded contents of the layer archive. Amazon Web Services SDK and
	// Amazon Web Services CLI clients handle the encoding for you.
	ZipFile []byte

	noSmithyDocumentSerde
}

// Details about a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// .
type LayerVersionContentOutput struct {

	// The SHA-256 hash of the layer archive.
	CodeSha256 *string

	// The size of the layer archive in bytes.
	CodeSize int64

	// A link to the layer archive in Amazon S3 that is valid for 10 minutes.
	Location *string

	// The Amazon Resource Name (ARN) of a signing job.
	SigningJobArn *string

	// The Amazon Resource Name (ARN) for a signing profile version.
	SigningProfileVersionArn *string

	noSmithyDocumentSerde
}

// Details about a version of an Lambda layer (https://docs.aws.amazon.com/lambda/latest/dg/configuration-layers.html)
// .
type LayerVersionsListItem struct {

	// A list of compatible instruction set architectures (https://docs.aws.amazon.com/lambda/latest/dg/foundation-arch.html)
	// .
	CompatibleArchitectures []Architecture

	// The layer's compatible runtimes. The following list includes deprecated
	// runtimes. For more information, see Runtime deprecation policy (https://docs.aws.amazon.com/lambda/latest/dg/lambda-runtimes.html#runtime-support-policy)
	// .
	CompatibleRuntimes []Runtime

	// The date that the version was created, in ISO 8601 format. For example,
	// 2018-11-27T15:10:45.123+0000 .
	CreatedDate *string

	// The description of the version.
	Description *string

	// The ARN of the layer version.
	LayerVersionArn *string

	// The layer's open-source license.
	LicenseInfo *string

	// The version number.
	Version int64

	noSmithyDocumentSerde
}

// A destination for events that failed processing.
type OnFailure struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string

	noSmithyDocumentSerde
}

// A destination for events that were processed successfully.
type OnSuccess struct {

	// The Amazon Resource Name (ARN) of the destination resource.
	Destination *string

	noSmithyDocumentSerde
}

// Details about the provisioned concurrency configuration for a function alias or
// version.
type ProvisionedConcurrencyConfigListItem struct {

	// The amount of provisioned concurrency allocated. When a weighted alias is used
	// during linear and canary deployments, this value fluctuates depending on the
	// amount of concurrency that is provisioned for the function versions.
	AllocatedProvisionedConcurrentExecutions *int32

	// The amount of provisioned concurrency available.
	AvailableProvisionedConcurrentExecutions *int32

	// The Amazon Resource Name (ARN) of the alias or version.
	FunctionArn *string

	// The date and time that a user last updated the configuration, in ISO 8601 format (https://www.iso.org/iso-8601-date-and-time-format.html)
	// .
	LastModified *string

	// The amount of provisioned concurrency requested.
	RequestedProvisionedConcurrentExecutions *int32

	// The status of the allocation process.
	Status ProvisionedConcurrencyStatusEnum

	// For failed allocations, the reason that provisioned concurrency could not be
	// allocated.
	StatusReason *string

	noSmithyDocumentSerde
}

// The ARN of the runtime and any errors that occured.
type RuntimeVersionConfig struct {

	// Error response when Lambda is unable to retrieve the runtime version for a
	// function.
	Error *RuntimeVersionError

	// The ARN of the runtime version you want the function to use.
	RuntimeVersionArn *string

	noSmithyDocumentSerde
}

// Any error returned when the runtime version information for the function could
// not be retrieved.
type RuntimeVersionError struct {

	// The error code.
	ErrorCode *string

	// The error message.
	Message *string

	noSmithyDocumentSerde
}

// (Amazon SQS only) The scaling configuration for the event source. To remove the
// configuration, pass an empty value.
type ScalingConfig struct {

	// Limits the number of concurrent instances that the Amazon SQS event source can
	// invoke.
	MaximumConcurrency *int32

	noSmithyDocumentSerde
}

// The self-managed Apache Kafka cluster for your event source.
type SelfManagedEventSource struct {

	// The list of bootstrap servers for your Kafka brokers in the following format:
	// "KAFKA_BOOTSTRAP_SERVERS": ["abc.xyz.com:xxxx","abc2.xyz.com:xxxx"] .
	Endpoints map[string][]string

	noSmithyDocumentSerde
}

// Specific configuration settings for a self-managed Apache Kafka event source.
type SelfManagedKafkaEventSourceConfig struct {

	// The identifier for the Kafka consumer group to join. The consumer group ID must
	// be unique among all your Kafka event sources. After creating a Kafka event
	// source mapping with the consumer group ID specified, you cannot update this
	// value. For more information, see Customizable consumer group ID (https://docs.aws.amazon.com/lambda/latest/dg/with-msk.html#services-msk-consumer-group-id)
	// .
	ConsumerGroupId *string

	noSmithyDocumentSerde
}

// The function's Lambda SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
// setting. Set ApplyOn to PublishedVersions to create a snapshot of the
// initialized execution environment when you publish a function version.
type SnapStart struct {

	// Set to PublishedVersions to create a snapshot of the initialized execution
	// environment when you publish a function version.
	ApplyOn SnapStartApplyOn

	noSmithyDocumentSerde
}

// The function's SnapStart (https://docs.aws.amazon.com/lambda/latest/dg/snapstart.html)
// setting.
type SnapStartResponse struct {

	// When set to PublishedVersions , Lambda creates a snapshot of the execution
	// environment when you publish a function version.
	ApplyOn SnapStartApplyOn

	// When you provide a qualified Amazon Resource Name (ARN) (https://docs.aws.amazon.com/lambda/latest/dg/configuration-versions.html#versioning-versions-using)
	// , this response element indicates whether SnapStart is activated for the
	// specified function version.
	OptimizationStatus SnapStartOptimizationStatus

	noSmithyDocumentSerde
}

// To secure and define access to your event source, you can specify the
// authentication protocol, VPC components, or virtual host.
type SourceAccessConfiguration struct {

	// The type of authentication protocol, VPC components, or virtual host for your
	// event source. For example: "Type":"SASL_SCRAM_512_AUTH" .
	//   - BASIC_AUTH – (Amazon MQ) The Secrets Manager secret that stores your broker
	//   credentials.
	//   - BASIC_AUTH – (Self-managed Apache Kafka) The Secrets Manager ARN of your
	//   secret key used for SASL/PLAIN authentication of your Apache Kafka brokers.
	//   - VPC_SUBNET – (Self-managed Apache Kafka) The subnets associated with your
	//   VPC. Lambda connects to these subnets to fetch data from your self-managed
	//   Apache Kafka cluster.
	//   - VPC_SECURITY_GROUP – (Self-managed Apache Kafka) The VPC security group used
	//   to manage access to your self-managed Apache Kafka brokers.
	//   - SASL_SCRAM_256_AUTH – (Self-managed Apache Kafka) The Secrets Manager ARN of
	//   your secret key used for SASL SCRAM-256 authentication of your self-managed
	//   Apache Kafka brokers.
	//   - SASL_SCRAM_512_AUTH – (Amazon MSK, Self-managed Apache Kafka) The Secrets
	//   Manager ARN of your secret key used for SASL SCRAM-512 authentication of your
	//   self-managed Apache Kafka brokers.
	//   - VIRTUAL_HOST –- (RabbitMQ) The name of the virtual host in your RabbitMQ
	//   broker. Lambda uses this RabbitMQ host as the event source. This property cannot
	//   be specified in an UpdateEventSourceMapping API call.
	//   - CLIENT_CERTIFICATE_TLS_AUTH – (Amazon MSK, self-managed Apache Kafka) The
	//   Secrets Manager ARN of your secret key containing the certificate chain (X.509
	//   PEM), private key (PKCS#8 PEM), and private key password (optional) used for
	//   mutual TLS authentication of your MSK/Apache Kafka brokers.
	//   - SERVER_ROOT_CA_CERTIFICATE – (Self-managed Apache Kafka) The Secrets Manager
	//   ARN of your secret key containing the root CA certificate (X.509 PEM) used for
	//   TLS encryption of your Apache Kafka brokers.
	Type SourceAccessType

	// The value for your chosen configuration in Type . For example: "URI":
	// "arn:aws:secretsmanager:us-east-1:01234567890:secret:MyBrokerSecretName" .
	URI *string

	noSmithyDocumentSerde
}

// The function's X-Ray (https://docs.aws.amazon.com/lambda/latest/dg/services-xray.html)
// tracing configuration. To sample and record incoming requests, set Mode to
// Active .
type TracingConfig struct {

	// The tracing mode.
	Mode TracingMode

	noSmithyDocumentSerde
}

// The function's X-Ray tracing configuration.
type TracingConfigResponse struct {

	// The tracing mode.
	Mode TracingMode

	noSmithyDocumentSerde
}

// The VPC security groups and subnets that are attached to a Lambda function. For
// more information, see Configuring a Lambda function to access resources in a VPC (https://docs.aws.amazon.com/lambda/latest/dg/configuration-vpc.html)
// .
type VpcConfig struct {

	// Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack
	// subnets.
	Ipv6AllowedForDualStack *bool

	// A list of VPC security group IDs.
	SecurityGroupIds []string

	// A list of VPC subnet IDs.
	SubnetIds []string

	noSmithyDocumentSerde
}

// The VPC security groups and subnets that are attached to a Lambda function.
type VpcConfigResponse struct {

	// Allows outbound IPv6 traffic on VPC functions that are connected to dual-stack
	// subnets.
	Ipv6AllowedForDualStack *bool

	// A list of VPC security group IDs.
	SecurityGroupIds []string

	// A list of VPC subnet IDs.
	SubnetIds []string

	// The ID of the VPC.
	VpcId *string

	noSmithyDocumentSerde
}

type noSmithyDocumentSerde = smithydocument.NoSerde

// UnknownUnionMember is returned when a union member is returned over the wire,
// but has an unknown tag.
type UnknownUnionMember struct {
	Tag   string
	Value []byte

	noSmithyDocumentSerde
}

func (*UnknownUnionMember) isInvokeWithResponseStreamResponseEvent() {}

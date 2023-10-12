// Code generated by smithy-go-codegen DO NOT EDIT.

package rekognition

import (
	"context"
	"errors"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	internalauth "github.com/aws/aws-sdk-go-v2/internal/auth"
	"github.com/aws/aws-sdk-go-v2/service/rekognition/types"
	smithyendpoints "github.com/aws/smithy-go/endpoints"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// This operation applies only to Amazon Rekognition Custom Labels. Detects custom
// labels in a supplied image by using an Amazon Rekognition Custom Labels model.
// You specify which version of a model version to use by using the
// ProjectVersionArn input parameter. You pass the input image as base64-encoded
// image bytes or as a reference to an image in an Amazon S3 bucket. If you use the
// AWS CLI to call Amazon Rekognition operations, passing image bytes is not
// supported. The image must be either a PNG or JPEG formatted file. For each
// object that the model version detects on an image, the API returns a (
// CustomLabel ) object in an array ( CustomLabels ). Each CustomLabel object
// provides the label name ( Name ), the level of confidence that the image
// contains the object ( Confidence ), and object location information, if it
// exists, for the label on the image ( Geometry ). To filter labels that are
// returned, specify a value for MinConfidence . DetectCustomLabelsLabels only
// returns labels with a confidence that's higher than the specified value. The
// value of MinConfidence maps to the assumed threshold values created during
// training. For more information, see Assumed threshold in the Amazon Rekognition
// Custom Labels Developer Guide. Amazon Rekognition Custom Labels metrics
// expresses an assumed threshold as a floating point value between 0-1. The range
// of MinConfidence normalizes the threshold value to a percentage value (0-100).
// Confidence responses from DetectCustomLabels are also returned as a percentage.
// You can use MinConfidence to change the precision and recall or your model. For
// more information, see Analyzing an image in the Amazon Rekognition Custom Labels
// Developer Guide. If you don't specify a value for MinConfidence ,
// DetectCustomLabels returns labels based on the assumed threshold of each label.
// This is a stateless API operation. That is, the operation does not persist any
// data. This operation requires permissions to perform the
// rekognition:DetectCustomLabels action. For more information, see Analyzing an
// image in the Amazon Rekognition Custom Labels Developer Guide.
func (c *Client) DetectCustomLabels(ctx context.Context, params *DetectCustomLabelsInput, optFns ...func(*Options)) (*DetectCustomLabelsOutput, error) {
	if params == nil {
		params = &DetectCustomLabelsInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "DetectCustomLabels", params, optFns, c.addOperationDetectCustomLabelsMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*DetectCustomLabelsOutput)
	out.ResultMetadata = metadata
	return out, nil
}

type DetectCustomLabelsInput struct {

	// Provides the input image either as bytes or an S3 object. You pass image bytes
	// to an Amazon Rekognition API operation by using the Bytes property. For
	// example, you would use the Bytes property to pass an image loaded from a local
	// file system. Image bytes passed by using the Bytes property must be
	// base64-encoded. Your code may not need to encode image bytes if you are using an
	// AWS SDK to call Amazon Rekognition API operations. For more information, see
	// Analyzing an Image Loaded from a Local File System in the Amazon Rekognition
	// Developer Guide. You pass images stored in an S3 bucket to an Amazon Rekognition
	// API operation by using the S3Object property. Images stored in an S3 bucket do
	// not need to be base64-encoded. The region for the S3 bucket containing the S3
	// object must match the region you use for Amazon Rekognition operations. If you
	// use the AWS CLI to call Amazon Rekognition operations, passing image bytes using
	// the Bytes property is not supported. You must first upload the image to an
	// Amazon S3 bucket and then call the operation using the S3Object property. For
	// Amazon Rekognition to process an S3 object, the user must have permission to
	// access the S3 object. For more information, see How Amazon Rekognition works
	// with IAM in the Amazon Rekognition Developer Guide.
	//
	// This member is required.
	Image *types.Image

	// The ARN of the model version that you want to use. Only models associated with
	// Custom Labels projects accepted by the operation. If a provided ARN refers to a
	// model version associated with a project for a different feature type, then an
	// InvalidParameterException is returned.
	//
	// This member is required.
	ProjectVersionArn *string

	// Maximum number of results you want the service to return in the response. The
	// service returns the specified number of highest confidence labels ranked from
	// highest confidence to lowest.
	MaxResults *int32

	// Specifies the minimum confidence level for the labels to return.
	// DetectCustomLabels doesn't return any labels with a confidence value that's
	// lower than this specified value. If you specify a value of 0, DetectCustomLabels
	// returns all labels, regardless of the assumed threshold applied to each label.
	// If you don't specify a value for MinConfidence , DetectCustomLabels returns
	// labels based on the assumed threshold of each label.
	MinConfidence *float32

	noSmithyDocumentSerde
}

type DetectCustomLabelsOutput struct {

	// An array of custom labels detected in the input image.
	CustomLabels []types.CustomLabel

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationDetectCustomLabelsMiddlewares(stack *middleware.Stack, options Options) (err error) {
	err = stack.Serialize.Add(&awsAwsjson11_serializeOpDetectCustomLabels{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsjson11_deserializeOpDetectCustomLabels{}, middleware.After)
	if err != nil {
		return err
	}
	if err = addlegacyEndpointContextSetter(stack, options); err != nil {
		return err
	}
	if err = addSetLoggerMiddleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddClientRequestIDMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddComputeContentLengthMiddleware(stack); err != nil {
		return err
	}
	if err = addResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = v4.AddComputePayloadSHA256Middleware(stack); err != nil {
		return err
	}
	if err = addRetryMiddlewares(stack, options); err != nil {
		return err
	}
	if err = addHTTPSignerV4Middleware(stack, options); err != nil {
		return err
	}
	if err = awsmiddleware.AddRawResponseToMetadata(stack); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecordResponseTiming(stack); err != nil {
		return err
	}
	if err = addClientUserAgent(stack, options); err != nil {
		return err
	}
	if err = smithyhttp.AddErrorCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = smithyhttp.AddCloseResponseBodyMiddleware(stack); err != nil {
		return err
	}
	if err = addDetectCustomLabelsResolveEndpointMiddleware(stack, options); err != nil {
		return err
	}
	if err = addOpDetectCustomLabelsValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opDetectCustomLabels(options.Region), middleware.Before); err != nil {
		return err
	}
	if err = awsmiddleware.AddRecursionDetection(stack); err != nil {
		return err
	}
	if err = addRequestIDRetrieverMiddleware(stack); err != nil {
		return err
	}
	if err = addResponseErrorMiddleware(stack); err != nil {
		return err
	}
	if err = addRequestResponseLogging(stack, options); err != nil {
		return err
	}
	if err = addendpointDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opDetectCustomLabels(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		SigningName:   "rekognition",
		OperationName: "DetectCustomLabels",
	}
}

type opDetectCustomLabelsResolveEndpointMiddleware struct {
	EndpointResolver EndpointResolverV2
	BuiltInResolver  builtInParameterResolver
}

func (*opDetectCustomLabelsResolveEndpointMiddleware) ID() string {
	return "ResolveEndpointV2"
}

func (m *opDetectCustomLabelsResolveEndpointMiddleware) HandleSerialize(ctx context.Context, in middleware.SerializeInput, next middleware.SerializeHandler) (
	out middleware.SerializeOutput, metadata middleware.Metadata, err error,
) {
	if awsmiddleware.GetRequiresLegacyEndpoints(ctx) {
		return next.HandleSerialize(ctx, in)
	}

	req, ok := in.Request.(*smithyhttp.Request)
	if !ok {
		return out, metadata, fmt.Errorf("unknown transport type %T", in.Request)
	}

	if m.EndpointResolver == nil {
		return out, metadata, fmt.Errorf("expected endpoint resolver to not be nil")
	}

	params := EndpointParameters{}

	m.BuiltInResolver.ResolveBuiltIns(&params)

	var resolvedEndpoint smithyendpoints.Endpoint
	resolvedEndpoint, err = m.EndpointResolver.ResolveEndpoint(ctx, params)
	if err != nil {
		return out, metadata, fmt.Errorf("failed to resolve service endpoint, %w", err)
	}

	req.URL = &resolvedEndpoint.URI

	for k := range resolvedEndpoint.Headers {
		req.Header.Set(
			k,
			resolvedEndpoint.Headers.Get(k),
		)
	}

	authSchemes, err := internalauth.GetAuthenticationSchemes(&resolvedEndpoint.Properties)
	if err != nil {
		var nfe *internalauth.NoAuthenticationSchemesFoundError
		if errors.As(err, &nfe) {
			// if no auth scheme is found, default to sigv4
			signingName := "rekognition"
			signingRegion := m.BuiltInResolver.(*builtInResolver).Region
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)

		}
		var ue *internalauth.UnSupportedAuthenticationSchemeSpecifiedError
		if errors.As(err, &ue) {
			return out, metadata, fmt.Errorf(
				"This operation requests signer version(s) %v but the client only supports %v",
				ue.UnsupportedSchemes,
				internalauth.SupportedSchemes,
			)
		}
	}

	for _, authScheme := range authSchemes {
		switch authScheme.(type) {
		case *internalauth.AuthenticationSchemeV4:
			v4Scheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4)
			var signingName, signingRegion string
			if v4Scheme.SigningName == nil {
				signingName = "rekognition"
			} else {
				signingName = *v4Scheme.SigningName
			}
			if v4Scheme.SigningRegion == nil {
				signingRegion = m.BuiltInResolver.(*builtInResolver).Region
			} else {
				signingRegion = *v4Scheme.SigningRegion
			}
			if v4Scheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4Scheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, signingName)
			ctx = awsmiddleware.SetSigningRegion(ctx, signingRegion)
			break
		case *internalauth.AuthenticationSchemeV4A:
			v4aScheme, _ := authScheme.(*internalauth.AuthenticationSchemeV4A)
			if v4aScheme.SigningName == nil {
				v4aScheme.SigningName = aws.String("rekognition")
			}
			if v4aScheme.DisableDoubleEncoding != nil {
				// The signer sets an equivalent value at client initialization time.
				// Setting this context value will cause the signer to extract it
				// and override the value set at client initialization time.
				ctx = internalauth.SetDisableDoubleEncoding(ctx, *v4aScheme.DisableDoubleEncoding)
			}
			ctx = awsmiddleware.SetSigningName(ctx, *v4aScheme.SigningName)
			ctx = awsmiddleware.SetSigningRegion(ctx, v4aScheme.SigningRegionSet[0])
			break
		case *internalauth.AuthenticationSchemeNone:
			break
		}
	}

	return next.HandleSerialize(ctx, in)
}

func addDetectCustomLabelsResolveEndpointMiddleware(stack *middleware.Stack, options Options) error {
	return stack.Serialize.Insert(&opDetectCustomLabelsResolveEndpointMiddleware{
		EndpointResolver: options.EndpointResolverV2,
		BuiltInResolver: &builtInResolver{
			Region:       options.Region,
			UseDualStack: options.EndpointOptions.UseDualStackEndpoint,
			UseFIPS:      options.EndpointOptions.UseFIPSEndpoint,
			Endpoint:     options.BaseEndpoint,
		},
	}, "ResolveEndpoint", middleware.After)
}

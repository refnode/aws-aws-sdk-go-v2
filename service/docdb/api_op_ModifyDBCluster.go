// Code generated by smithy-go-codegen DO NOT EDIT.

package docdb

import (
	"context"
	"fmt"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/aws/signer/v4"
	"github.com/aws/aws-sdk-go-v2/service/docdb/types"
	"github.com/aws/smithy-go/middleware"
	smithyhttp "github.com/aws/smithy-go/transport/http"
)

// Modifies a setting for an Amazon DocumentDB cluster. You can change one or more
// database configuration parameters by specifying these parameters and the new
// values in the request.
func (c *Client) ModifyDBCluster(ctx context.Context, params *ModifyDBClusterInput, optFns ...func(*Options)) (*ModifyDBClusterOutput, error) {
	if params == nil {
		params = &ModifyDBClusterInput{}
	}

	result, metadata, err := c.invokeOperation(ctx, "ModifyDBCluster", params, optFns, c.addOperationModifyDBClusterMiddlewares)
	if err != nil {
		return nil, err
	}

	out := result.(*ModifyDBClusterOutput)
	out.ResultMetadata = metadata
	return out, nil
}

// Represents the input to ModifyDBCluster .
type ModifyDBClusterInput struct {

	// The cluster identifier for the cluster that is being modified. This parameter
	// is not case sensitive. Constraints:
	//   - Must match the identifier of an existing DBCluster .
	//
	// This member is required.
	DBClusterIdentifier *string

	// A value that indicates whether major version upgrades are allowed. Constraints:
	// You must allow major version upgrades when specifying a value for the
	// EngineVersion parameter that is a different major version than the DB cluster's
	// current version.
	AllowMajorVersionUpgrade *bool

	// A value that specifies whether the changes in this request and any pending
	// changes are asynchronously applied as soon as possible, regardless of the
	// PreferredMaintenanceWindow setting for the cluster. If this parameter is set to
	// false , changes to the cluster are applied during the next maintenance window.
	// The ApplyImmediately parameter affects only the NewDBClusterIdentifier and
	// MasterUserPassword values. If you set this parameter value to false , the
	// changes to the NewDBClusterIdentifier and MasterUserPassword values are applied
	// during the next maintenance window. All other changes are applied immediately,
	// regardless of the value of the ApplyImmediately parameter. Default: false
	ApplyImmediately *bool

	// The number of days for which automated backups are retained. You must specify a
	// minimum value of 1. Default: 1 Constraints:
	//   - Must be a value from 1 to 35.
	BackupRetentionPeriod *int32

	// The configuration setting for the log types to be enabled for export to Amazon
	// CloudWatch Logs for a specific instance or cluster. The EnableLogTypes and
	// DisableLogTypes arrays determine which logs are exported (or not exported) to
	// CloudWatch Logs.
	CloudwatchLogsExportConfiguration *types.CloudwatchLogsExportConfiguration

	// The name of the cluster parameter group to use for the cluster.
	DBClusterParameterGroupName *string

	// Specifies whether this cluster can be deleted. If DeletionProtection is
	// enabled, the cluster cannot be deleted unless it is modified and
	// DeletionProtection is disabled. DeletionProtection protects clusters from being
	// accidentally deleted.
	DeletionProtection *bool

	// The version number of the database engine to which you want to upgrade.
	// Changing this parameter results in an outage. The change is applied during the
	// next maintenance window unless ApplyImmediately is enabled. To list all of the
	// available engine versions for Amazon DocumentDB use the following command: aws
	// docdb describe-db-engine-versions --engine docdb --query
	// "DBEngineVersions[].EngineVersion"
	EngineVersion *string

	// The password for the master database user. This password can contain any
	// printable ASCII character except forward slash (/), double quote ("), or the
	// "at" symbol (@). Constraints: Must contain from 8 to 100 characters.
	MasterUserPassword *string

	// The new cluster identifier for the cluster when renaming a cluster. This value
	// is stored as a lowercase string. Constraints:
	//   - Must contain from 1 to 63 letters, numbers, or hyphens.
	//   - The first character must be a letter.
	//   - Cannot end with a hyphen or contain two consecutive hyphens.
	// Example: my-cluster2
	NewDBClusterIdentifier *string

	// The port number on which the cluster accepts connections. Constraints: Must be
	// a value from 1150 to 65535 . Default: The same port as the original cluster.
	Port *int32

	// The daily time range during which automated backups are created if automated
	// backups are enabled, using the BackupRetentionPeriod parameter. The default is
	// a 30-minute window selected at random from an 8-hour block of time for each
	// Amazon Web Services Region. Constraints:
	//   - Must be in the format hh24:mi-hh24:mi .
	//   - Must be in Universal Coordinated Time (UTC).
	//   - Must not conflict with the preferred maintenance window.
	//   - Must be at least 30 minutes.
	PreferredBackupWindow *string

	// The weekly time range during which system maintenance can occur, in Universal
	// Coordinated Time (UTC). Format: ddd:hh24:mi-ddd:hh24:mi The default is a
	// 30-minute window selected at random from an 8-hour block of time for each Amazon
	// Web Services Region, occurring on a random day of the week. Valid days: Mon,
	// Tue, Wed, Thu, Fri, Sat, Sun Constraints: Minimum 30-minute window.
	PreferredMaintenanceWindow *string

	// The storage type to associate with the DB cluster. For information on storage
	// types for Amazon DocumentDB clusters, see Cluster storage configurations in the
	// Amazon DocumentDB Developer Guide. Valid values for storage type - standard |
	// iopt1 Default value is standard
	StorageType *string

	// A list of virtual private cloud (VPC) security groups that the cluster will
	// belong to.
	VpcSecurityGroupIds []string

	noSmithyDocumentSerde
}

type ModifyDBClusterOutput struct {

	// Detailed information about a cluster.
	DBCluster *types.DBCluster

	// Metadata pertaining to the operation's result.
	ResultMetadata middleware.Metadata

	noSmithyDocumentSerde
}

func (c *Client) addOperationModifyDBClusterMiddlewares(stack *middleware.Stack, options Options) (err error) {
	if err := stack.Serialize.Add(&setOperationInputMiddleware{}, middleware.After); err != nil {
		return err
	}
	err = stack.Serialize.Add(&awsAwsquery_serializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	err = stack.Deserialize.Add(&awsAwsquery_deserializeOpModifyDBCluster{}, middleware.After)
	if err != nil {
		return err
	}
	if err := addProtocolFinalizerMiddlewares(stack, options, "ModifyDBCluster"); err != nil {
		return fmt.Errorf("add protocol finalizers: %v", err)
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
	if err = addSetLegacyContextSigningOptionsMiddleware(stack); err != nil {
		return err
	}
	if err = addOpModifyDBClusterValidationMiddleware(stack); err != nil {
		return err
	}
	if err = stack.Initialize.Add(newServiceMetadataMiddleware_opModifyDBCluster(options.Region), middleware.Before); err != nil {
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
	if err = addDisableHTTPSMiddleware(stack, options); err != nil {
		return err
	}
	return nil
}

func newServiceMetadataMiddleware_opModifyDBCluster(region string) *awsmiddleware.RegisterServiceMetadata {
	return &awsmiddleware.RegisterServiceMetadata{
		Region:        region,
		ServiceID:     ServiceID,
		OperationName: "ModifyDBCluster",
	}
}

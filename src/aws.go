package pike

import (
	"encoding/json"
	"log"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result ResourceV2) []string {

	var Permissions []string
	if result.TypeName == "resource" {
		Permissions = GetAWSResourcePermissions(result)
	} else {
		Permissions = GetAWSDataPermissions(result)
	}

	return Permissions
}

// GetAWSResourcePermissions looks up permissions required for resources
func GetAWSResourcePermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "aws_s3_bucket":
		Permissions = GetPermissionMap(awsS3Bucket, result.Attributes)
	case "aws_s3_bucket_acl":
		Permissions = GetPermissionMap(awsS3BucketACL, result.Attributes)
	case "aws_s3_bucket_versioning":
		Permissions = GetPermissionMap(awsS3BucketVersioning, result.Attributes)
	case "aws_s3_bucket_server_side_encryption_configuration":
		Permissions = GetPermissionMap(awsS3BucketServerSideEncryptionConfiguration, result.Attributes)
	case "aws_s3_bucket_public_access_block":
		Permissions = GetPermissionMap(awsS3BucketPublicAccessBlock, result.Attributes)
	case "aws_instance":
		Permissions = GetPermissionMap(awsInstance, result.Attributes)
	case "aws_security_group":
		Permissions = GetPermissionMap(awsSecurityGroup, result.Attributes)
	case "aws_security_group_rule":
		Permissions = GetPermissionMap(awsSecurityGroupRule, result.Attributes)
	case "aws_lambda_function":
		Permissions = GetPermissionMap(awsLambdaFunction, result.Attributes)
	case "aws_vpc":
		Permissions = GetPermissionMap(awsVpc, result.Attributes)
	case "aws_subnet":
		Permissions = GetPermissionMap(awsSubnet, result.Attributes)
	case "aws_network_acl":
		Permissions = GetPermissionMap(awsNetworkACL, result.Attributes)
	case "aws_kms_key":
		Permissions = GetPermissionMap(awsKmsKey, result.Attributes)
	case "aws_iam_role":
		Permissions = GetPermissionMap(awsIamRole, result.Attributes)
	case "aws_iam_role_policy":
		Permissions = GetPermissionMap(awsIamRolePolicy, result.Attributes)
	case "aws_iam_role_policy_attachment":
		Permissions = GetPermissionMap(awsIamRolePolicyAttachment, result.Attributes)
	case "aws_iam_policy":
		Permissions = GetPermissionMap(awsIamPolicy, result.Attributes)
	case "aws_iam_instance_profile":
		Permissions = GetPermissionMap(awsIamInstanceProfile, result.Attributes)
	case "aws_iam_access_key":
		Permissions = GetPermissionMap(awsIamAccessKey, result.Attributes)
	case "aws_iam_group":
		Permissions = GetPermissionMap(awsIamGroup, result.Attributes)
	case "aws_iam_group_membership":
		Permissions = GetPermissionMap(awsIamGroupMembership, result.Attributes)
	case "aws_iam_group_policy":
		Permissions = GetPermissionMap(awsIamGroupPolicy, result.Attributes)
	case "aws_iam_group_policy_attachment":
		Permissions = GetPermissionMap(awsIamGroupPolicyAttachment, result.Attributes)
	case "aws_iam_policy_attachment":
		Permissions = GetPermissionMap(awsIamPolicyAttachment, result.Attributes)
	case "aws_iam_service_linked_role":
		Permissions = GetPermissionMap(awsIamServiceLinkedRole, result.Attributes)
	case "aws_iam_user":
		Permissions = GetPermissionMap(awsIamUser, result.Attributes)
	case "aws_iam_user_login_profile":
		Permissions = GetPermissionMap(awsIamUserLoginProfile, result.Attributes)
	case "aws_iam_user_policy":
		Permissions = GetPermissionMap(awsIamUserPolicy, result.Attributes)
	case "aws_iam_user_policy_attachment":
		Permissions = GetPermissionMap(awsIamUserPolicyAttachment, result.Attributes)
	case "aws_mq_broker":
		Permissions = GetPermissionMap(awsMqBroker, result.Attributes)
	case "aws_mq_configuration":
		Permissions = GetPermissionMap(awsMqConfiguration, result.Attributes)
	case "aws_cloudwatch_log_group":
		Permissions = GetPermissionMap(awsCloudwatchLogGroup, result.Attributes)
	case "aws_cloudwatch_event_rule":
		Permissions = GetPermissionMap(awsCloudwatchEventRule, result.Attributes)
	case "aws_cloudwatch_event_target":
		Permissions = GetPermissionMap(awsCloudwatchEventTarget, result.Attributes)
	case "aws_cloudwatch_log_metric_filter":
		Permissions = GetPermissionMap(awsCloudwatchLogMetricFilter, result.Attributes)
	case "aws_cloudwatch_log_resource_policy":
		Permissions = GetPermissionMap(awsCloudwatchLogResourcePolicy, result.Attributes)
	case "aws_cloudwatch_log_subscription_filter":
		Permissions = GetPermissionMap(awsCloudwatchLogSubscriptionFilter, result.Attributes)
	case "aws_cloudwatch_metric_alarm":
		Permissions = GetPermissionMap(awsCloudwatchMetricAlarm, result.Attributes)
	case "aws_route53_record":
		Permissions = GetPermissionMap(awsRoute53Record, result.Attributes)
	case "aws_sns_topic":
		Permissions = GetPermissionMap(awsSnsTopic, result.Attributes)
	case "aws_key_pair":
		Permissions = GetPermissionMap(awsKeyPair, result.Attributes)
	case "aws_db_instance":
		Permissions = GetPermissionMap(awsDbInstance, result.Attributes)
	case "aws_dynamodb_table":
		Permissions = GetPermissionMap(awsDynamodbTable, result.Attributes)
	case "aws_ssm_parameter":
		Permissions = GetPermissionMap(awsSsmParameter, result.Attributes)
	case "aws_route":
		Permissions = GetPermissionMap(awsRoute, result.Attributes)
	default:
		log.Printf("%s not implemented", result.Name)
	}
	return Permissions
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// GetPermissionMap Anonymous parsing
func GetPermissionMap(raw []byte, attributes []string) []string {
	var mappings []interface{}
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		log.Print(err)
	}
	temp := mappings[0].(map[string]interface{})
	myAttributes := temp["attributes"].(map[string]interface{})

	var found []string

	for _, attribute := range attributes {
		if myAttributes[attribute] != nil {
			entries := myAttributes[attribute].([]interface{})
			for _, entry := range entries {
				found = append(found, entry.(string))
			}
		}
	}

	actions := []string{"apply", "plan", "modify", "destroy"}

	for _, action := range actions {
		if temp[action] != nil {
			myentries := temp[action].([]interface{})
			for _, entry := range myentries {
				found = append(found, entry.(string))
			}
		}
	}

	return found
}

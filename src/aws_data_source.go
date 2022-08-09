package pike

import "log"

// GetAWSDataPermissions gets permissions required for datasources
func GetAWSDataPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "aws_vpcs":
		Permissions = GetPermissionMap(dataAwsVpcs, result.Attributes)
	case "aws_subnet_ids", "aws_subnet", "aws_subnets":
		Permissions = GetPermissionMap(dataAwsSubnetIds, result.Attributes)
	case "aws_ami":
		Permissions = GetPermissionMap(dataAwsAmi, result.Attributes)
	case "aws_iam_policy":
		Permissions = GetPermissionMap(dataAwsIamPolicy, result.Attributes)
	case "aws_iam_role":
		Permissions = GetPermissionMap(dataAwsIamRole, result.Attributes)
	case "aws_s3_bucket":
		Permissions = GetPermissionMap(dataAwsS3Bucket, result.Attributes)
	case "aws_vpc":
		Permissions = GetPermissionMap(dataAwsVpc, result.Attributes)
	case "aws_availability_zones":
		Permissions = GetPermissionMap(dataAwsAvailabilityZones, result.Attributes)
	case "aws_caller_identity", "aws_iam_policy_document", "aws_region", "aws_canonical_user_id":
		// do nothing
	default:
		log.Printf("%s.%s not implemented", result.TypeName, result.Name)
	}
	return Permissions
}

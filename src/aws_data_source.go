package pike

import "log"

// GetAWSDataPermissions gets permissions required for datasources
func GetAWSDataPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "aws_vpcs":
		Permissions = GetPermissionMap(data_aws_vpcs, result.Attributes)
	case "aws_subnet_ids","aws_subnet", "aws_subnets":
		Permissions = GetPermissionMap(data_aws_subnet_ids, result.Attributes)
	case "aws_ami":
		Permissions = GetPermissionMap(data_aws_ami, result.Attributes)
	case "aws_vpc":
		Permissions = GetPermissionMap(data_aws_vpc, result.Attributes)
	case "aws_availability_zones":
		Permissions = GetPermissionMap(data_aws_availability_zones, result.Attributes)
	case "aws_caller_identity","aws_iam_policy_document","aws_region":
		//do nothing
	default:
		log.Printf("%s.%s not implemented", result.TypeName, result.Name)
	}
	return Permissions
}

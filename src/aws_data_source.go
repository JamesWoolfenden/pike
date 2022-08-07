package pike

import "log"

// GetAWSDataPermissions gets permissions required for datasources
func GetAWSDataPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	case "aws_vpcs":
		Permissions = GetPermissionMap(data_aws_vpcs, result.Attributes)
	case "aws_subnet_ids":
		Permissions = GetPermissionMap(data_aws_subnet_ids, result.Attributes)
	case "aws_caller_identity","aws_iam_policy_document":
		//do nothing
	default:
		log.Printf("%s.%s not implemented", result.TypeName, result.Name)
	}
	return Permissions
}

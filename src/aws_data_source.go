package pike

import "log"

// GetAWSDataPermissions gets permissions required for datasources
func GetAWSDataPermissions(result ResourceV2) []string {

	TFLookup := map[string]interface{}{
		"aws_vpcs":                dataAwsVpcs,
		"aws_subnet_ids":          dataAwsSubnetIds,
		"aws_subnet":              dataAwsSubnetIds,
		"aws_subnets":             dataAwsSubnetIds,
		"aws_ami":                 dataAwsAmi,
		"aws_iam_policy":          dataAwsIamPolicy,
		"aws_iam_role":            dataAwsIamRole,
		"aws_s3_bucket":           dataAwsS3Bucket,
		"aws_vpc":                 dataAwsVpc,
		"aws_availability_zones":  dataAwsAvailabilityZones,
		"aws_caller_identity":     placeholder,
		"aws_iam_policy_document": placeholder,
		"aws_region":              placeholder,
		"aws_canonical_user_id":   placeholder,
	}

	var Permissions []string

	temp := TFLookup[result.Name]
	if temp != nil {
		Permissions = GetPermissionMap(TFLookup[result.Name].([]byte), result.Attributes)
	} else {
		log.Printf("%s not implemented", result.Name)
	}

	return Permissions
}

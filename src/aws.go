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
		Permissions = GetPermissionMap(aws_s3_bucket, result.Attributes)
	case "aws_instance":
		Permissions = GetPermissionMap(aws_instance, result.Attributes)
	case "aws_security_group":
		Permissions = GetPermissionMap(aws_security_group, result.Attributes)
	case "aws_lambda_function":
		Permissions = GetPermissionMap(aws_lambda_function, result.Attributes)
	case "aws_vpc":
		Permissions = GetPermissionMap(aws_vpc, result.Attributes)
	case "aws_subnet":
		Permissions = GetPermissionMap(aws_subnet, result.Attributes)
	case "aws_network_acl":
		Permissions = GetPermissionMap(aws_network_acl, result.Attributes)
	case "aws_kms_key":
		Permissions = GetPermissionMap(aws_kms_key, result.Attributes)
	case "aws_iam_role":
		Permissions = GetPermissionMap(aws_iam_role, result.Attributes)
	case "aws_mq_broker":
		Permissions = GetPermissionMap(aws_mq_broker, result.Attributes)
	case "aws_mq_configuration":
		Permissions = GetPermissionMap(aws_mq_configuration, result.Attributes)
	default:
		log.Printf("%s not implemented", result.Name)
	}
	return Permissions
}

// GetAWSDataPermissions gets permissions required for datasources
func GetAWSDataPermissions(result ResourceV2) []string {
	var Permissions []string
	switch result.Name {
	//case "aws_s3_bucket":
	//	Permissions = GetPermissionMap(aws_s3_bucket, result.Attributes)
	//case "aws_instance":
	//	Permissions = GetPermissionMap(aws_instance, result.Attributes)
	//case "aws_security_group":
	//	Permissions = GetPermissionMap(aws_security_group, result.Attributes)
	//case "aws_lambda_function":
	//	Permissions = GetPermissionMap(aws_lambda_function, result.Attributes)
	//case "aws_vpc":
	//	Permissions = GetPermissionMap(aws_vpc, result.Attributes)
	//case "aws_subnet":
	//	Permissions = GetPermissionMap(aws_subnet, result.Attributes)
	//case "aws_network_acl":
	//	Permissions = GetPermissionMap(aws_network_acl, result.Attributes)
	//case "aws_kms_key":
	//	Permissions = GetPermissionMap(aws_kms_key, result.Attributes)
	//case "aws_iam_role":
	//	Permissions = GetPermissionMap(aws_iam_role, result.Attributes)
	//case "aws_mq_broker":
	//	Permissions = GetPermissionMap(aws_mq_broker, result.Attributes)
	//case "aws_mq_configuration":
	//	Permissions = GetPermissionMap(aws_mq_configuration, result.Attributes)
	default:
		log.Printf("%s.%s not implemented", result.TypeName, result.Name)
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

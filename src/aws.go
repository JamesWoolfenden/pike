package pike

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/hcl/hcl/ast"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result template) []string {
	myAttributes := GetAttributes(result)
	var Permissions []string
	switch result.Resource.name {
	case "aws_s3_bucket":
		Permissions = GetPermissionMap(aws_s3_bucket, myAttributes)
	case "aws_instance":
		Permissions = GetPermissionMap(aws_instance, myAttributes)
	case "aws_security_group":
		Permissions = GetPermissionMap(aws_security_group, myAttributes)
	case "aws_lambda_function":
		Permissions = GetPermissionMap(aws_lambda_function, myAttributes)
	case "aws_vpc":
		Permissions = GetPermissionMap(aws_vpc, myAttributes)
	case "aws_subnet":
		Permissions = GetPermissionMap(aws_subnet, myAttributes)
	case "aws_network_acl":
		Permissions = GetPermissionMap(aws_network_acl, myAttributes)
	default:
		log.Printf("%s %s not found", result.Template, result.Resource.name)
	}

	return Permissions
}

// GetAttributes gets the name of the important attributes for this resource
func GetAttributes(result template) []string {
	temp := result.Resource.code.Val.(*ast.ObjectType)
	attributes := temp.List.Items
	var myAttributes []string
	for _, item := range attributes {
		mytemp := item.Keys
		myAttributes = append(myAttributes, mytemp[0].Token.Text)
	}
	return myAttributes
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

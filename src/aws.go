package pike

import (
	"encoding/json"
	"log"

	"github.com/hashicorp/hcl/hcl/ast"
)

// GetAWSPermissions for AWS resources
func GetAWSPermissions(result template) []interface{} {
	myAttributes := GetAttributes(result)
	var Permissions []interface{}
	switch result.Resource.name {
	case "aws_s3_bucket":
		Permissions = GetPermissionMap(aws_s3_bucket, myAttributes)
	case "aws_instance":
		Permissions = GetPermissionMap(aws_instance, myAttributes)
	case "aws_security_group":
		Permissions = GetPermissionMap(aws_security_group, myAttributes)
	case "aws_lambda":
		Permissions = GetPermissionMap(aws_lambda, myAttributes)
		
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
func GetPermissionMap(raw []byte, attributes []string) []interface{} {
	var mappings []interface{}
	err := json.Unmarshal(raw, &mappings)
	if err != nil {
		log.Print(err)
	}
	temp := mappings[0].(map[string]interface{})
	myAttributes := temp["attributes"].(map[string]interface{})

	var found []interface{}

	for _, attribute := range attributes {
		if myAttributes[attribute] != nil {
			found = append(found, myAttributes[attribute])
		}
	}

	actions := []string{"apply", "plan", "modify", "destroy"}

	for _, action := range actions {
		if temp[action] != nil {
			found = append(found, temp[action])
		}
	}

	return found
}

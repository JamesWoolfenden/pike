package pike

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// GetAPI determines which APIs a resource targets
func GetAPI(resource string) string {

	m := map[string]string{
		"aws_instance":                           "ec2",
		"aws_s3_bucket":                          "s3",
		"aws_s3_bucket_accelerate_configuration": "s3",
		"aws_s3_bucket_policy":                   "s3",
		"aws_s3_bucket_logging":                  "s3",
		"aws_vpc":                                "vpc",
	}

	api := m[resource]
	if api == "" {
		//log.Printf("Resource:%s does not map", resource)
		return ""
	}

	return api
}

// GetResources retrieves all the resources in a tf file
func GetResources(file fs.FileInfo, dirname string) []Resource {

	var results []Resource

	src, _ := ioutil.ReadFile(dirname + file.Name())

	myCode, _ := hcl.Parse(string(src))
	Tree := myCode.Node.(*ast.ObjectList)

	for _, item := range Tree.Items {
		var temp Resource
		temp.name = strings.Trim(item.Keys[1].Token.Text, "\"")
		temp.path = dirname + file.Name()
		temp.code = *item
		results = append(results, temp)
	}

	// resources, filename, code
	return results
}

// GetProvider retrieves the provider from the resource
func GetProvider(resource string) string {
	return strings.Split(resource, "_")[0]
}

// GetPermission determines the IAM permissions required and returns a list of permission
func GetPermission(result template) []string {
	var myPermission []string
	switch result.Provider {
	case "aws":
		myPermission = GetAWSPermissions(result)
	case "azure":
	case "gcp":
	default:
		log.Fatal("Provider {result.Provider} not found")
	}

	return myPermission
}

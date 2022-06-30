package pike

import (
	"fmt"
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
		log.Printf("Resource:%s does not map", resource)
		return ""
	}

	return api
}

// GetResources retrieves all the resources in a tf file
func GetResources(file fs.FileInfo, dirname string) ([]string, string, string) {
	var result []string

	src, _ := ioutil.ReadFile(dirname + file.Name())

	myCode, _ := hcl.Parse(string(src))

	fmt.Print(myCode)
	Tree := myCode.Node.(*ast.ObjectList)

	for _, item := range Tree.Items {
		result = append(result, strings.Trim(item.Keys[1].Token.Text, "\""))
	}

	fmt.Println(myCode)
	return result, dirname + file.Name(), string(src)
}

// GetProvider retrieves the provider from the resource
func GetProvider(resource string) string {
	return strings.Split(resource, "_")[0]
}

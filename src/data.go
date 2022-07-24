package pike

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// GetResources retrieves all the resources in a tf file
func GetResources(file fs.FileInfo, dirname string) []Resource {

	var results []Resource

	src, err := ioutil.ReadFile(dirname + file.Name())
	if err != nil {
		log.Fatal(err)
	}

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
	if strings.Contains(resource, "_"){
		return strings.Split(resource, "_")[0]
	}
	return ""
}
// GetPermission determines the IAM permissions required and returns a list of permission
func GetPermission(result template) []interface{} {
	var myPermission []interface{}
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

package pike

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
	"os"
	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// GetResources retrieves all the resources in a tf file
func GetResources(file fs.FileInfo, dirname string) ([]Resource, error ){

	var results []Resource

	fullfile:=dirname + string(os.PathSeparator)+ file.Name()
	
	src, err := ioutil.ReadFile(fullfile)

	if err != nil {
		log.Fatal(err)
	}

	myCode, err := hcl.Parse(string(src))

	if err != nil {
		log.Printf("failed to parse %s", fullfile )
		log.Print(err)
		return nil, err
	}

	Tree := myCode.Node.(*ast.ObjectList)

	for _, item := range Tree.Items {
		var temp Resource
		temp.name = strings.Trim(item.Keys[1].Token.Text, "\"")
		temp.path = dirname + file.Name()
		temp.code = *item
		results = append(results, temp)
	}

	// resources, filename, code
	return results, nil
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
		if result.Provider != "" {
			log.Printf("Provider %s was not found", result.Provider)
		}
	}

	return myPermission
}

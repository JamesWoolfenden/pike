package pike

import (
	"errors"
	"io/ioutil"
	"log"
	"strings"

	"github.com/hashicorp/hcl"
	"github.com/hashicorp/hcl/hcl/ast"
)

// GetResources retrieves all the resources in a tf file
func GetResources(file string) []Resource {

	var results []Resource

	src, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	myCode, err := hcl.Parse(string(src))

	if err != nil {
		log.Printf("failed to parse %s", file)
	}

	Tree := myCode.Node.(*ast.ObjectList)

	for _, item := range Tree.Items {
		var temp Resource
		temp.name = strings.Trim(item.Keys[1].Token.Text, "\"")
		temp.path = file
		temp.code = *item
		results = append(results, temp)
	}

	// resources, filename, code
	return results
}

// GetProvider retrieves the provider from the resource
func GetProvider(resource string) string {
	if strings.Contains(resource, "_") {
		return strings.Split(resource, "_")[0]
	}
	return ""
}

// GetPermission determines the IAM permissions required and returns a list of permission
func GetPermission(result template) (Sorted, error) {
	var myPermission Sorted
	switch result.Provider {
	case "aws":
		myPermission.AWS = GetAWSPermissions(result)
	case "azure":
		return myPermission, errors.New("not implemented")
	case "gcp", "google":
		myPermission.GCP = GetGCPPermissions(result)
	default:
		if result.Provider != "" {
			log.Printf("Provider %s was not found", result.Provider)
		}
	}

	return myPermission, nil
}

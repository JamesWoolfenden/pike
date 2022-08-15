package pike

import (
	"log"
	"os"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/hashicorp/hcl/v2/hclparse"
)

// GetResources retrieves all the resources in a tf file
func GetResources(file string) ([]ResourceV2, error) {

	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	parser := hclparse.NewParser()
	parsedFile, fileDiags := parser.ParseHCL(src, file)

	if fileDiags != nil {
		return nil, fileDiags
	}

	var Resources []ResourceV2

	temp := parsedFile.Body.(*hclsyntax.Body)

	for _, block := range temp.Blocks {
		var resource ResourceV2
		resource.TypeName = block.Type

		ignore := []string{"terraform", "output", "provider", "variable", "locals", "module", "template"}

		if stringInSlice(resource.TypeName, ignore) {
			continue
		}

		if block.Labels != nil {
			resource.Name = block.Labels[0]

			if len(block.Labels) > 1 {
				resource.ResourceName = block.Labels[1]
			}
		}

		var attributes []string
		for _, attribute := range block.Body.Attributes {
			attributes = append(attributes, attribute.Name)
		}
		resource.Attributes = attributes
		resource.Provider = GetHCLType(block.Labels[0])
		Resources = append(Resources, resource)
	}

	return Resources, nil
}

// GetProvider retrieves the provider from the resource
func GetProvider(resource string) string {
	if strings.Contains(resource, "_") {
		return strings.Split(resource, "_")[0]
	}
	return ""
}

// GetPermission determines the IAM permissions required and returns a list of permission
func GetPermission(result ResourceV2) (Sorted, error) {
	var myPermission Sorted
	switch result.Provider {
	case "aws":
		myPermission.AWS = GetAWSPermissions(result)
	case "azurerm", "oci", "digitalocean", "linode", "helm":
		log.Printf("Provider %s not yet implemented", result.Provider)
		return myPermission, nil
	case "gcp", "google":
		myPermission.GCP = GetGCPPermissions(result)
	case "provider", "random", "main", "ip", "http", "test", "local",
		"archive", "tls", "template", "null", "time":
		return myPermission, nil
	default:
		if result.Provider != "" {
			log.Printf("Provider %s was not found", result.Provider)
		}
	}

	return myPermission, nil
}

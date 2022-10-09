package pike

import (
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/hashicorp/hcl/v2/hclsyntax"

	"github.com/hashicorp/hcl/v2/hclparse"
)

// GetResources retrieves all the resources in a tf file
func GetResources(file string, dirName string) ([]ResourceV2, error) {
	var Resources []ResourceV2
	temp, err := GetResourceBlocks(file)
	if err != nil {
		return Resources, err
	}

	for _, block := range temp.Blocks {
		var resource ResourceV2
		resource.TypeName = block.Type

		ignore := []string{"terraform", "output", "provider", "variable", "locals", "template"}

		if stringInSlice(resource.TypeName, ignore) {
			continue
		}

		if strings.Contains(resource.TypeName, "module") {
			LocalResources := getLocalModules(block, dirName)
			Resources = append(LocalResources, Resources...)
			if err != nil {
				log.Print(err)
				continue
			}
		}

		if block.Labels != nil {
			resource.Name = block.Labels[0]

			if len(block.Labels) > 1 {
				resource.ResourceName = block.Labels[1]
			}
		}

		var attributes []string

		resource.Attributes = GetBlockAttributes(attributes, block)
		if len(block.Labels) > 0 {
			resource.Provider = GetHCLType(block.Labels[0])
		} else {
			resource.Provider = "unknown"
			log.Print("parsing error for ", block.Type)
		}
		Resources = append(Resources, resource)
	}

	return Resources, nil
}

// GetResourceBlocks breaks down a file into resources
func GetResourceBlocks(file string) (*hclsyntax.Body, error) {
	temp, _ := filepath.Abs(file)
	src, err := os.ReadFile(temp)
	if err != nil {
		return nil, err
	}

	parser := hclparse.NewParser()
	parsedFile, fileDiags := parser.ParseHCL(src, file)

	if fileDiags != nil {
		return nil, fileDiags
	}

	return parsedFile.Body.(*hclsyntax.Body), err
}

func getLocalModules(block *hclsyntax.Block, dirName string) []ResourceV2 {
	var Resources []ResourceV2
	modulePath := GetModulePath(block)

	_, err := os.Stat(modulePath)
	if err != nil {
		// could be totally valid
		return nil
	}

	// have the path to the module
	modulePath = filepath.Join(dirName, "/", modulePath)

	// now process these extras
	ExtraFiles, _ := GetTF(modulePath)
	for _, file := range ExtraFiles {
		resource, err := GetResources(file, dirName)
		if err == nil {
			Resources = append(Resources, resource...)
		}
	}
	return Resources
}

// GetModulePath extracts the source location from a module
func GetModulePath(block *hclsyntax.Block) string {
	var modulePath string

	attributes := block.Body.Attributes
	value := attributes["source"].Expr

	castValue := value.(*hclsyntax.TemplateExpr)
	parts := castValue.Parts
	for _, part := range parts {
		myPart := part.(*hclsyntax.LiteralValueExpr)
		modulePath = myPart.Val.AsString()
	}

	return modulePath
}

// GetBlockAttributes walks through a blocks getting all blocks and attributes
func GetBlockAttributes(attributes []string, block *hclsyntax.Block) []string {
	for _, attribute := range block.Body.Attributes {
		attributes = append(attributes, attribute.Name)
	}
	for _, block := range block.Body.Blocks {
		// Also add in block names
		if block.Type != "resource" {
			attributes = append(attributes, block.Type)
		}

		attributes = GetBlockAttributes(attributes, block)
	}
	return attributes
}

// GetPermission determines the IAM permissions required and returns a list of permission
func GetPermission(result ResourceV2) (Sorted, error) {
	var err error
	var myPermission Sorted
	switch result.Provider {
	case "aws":
		myPermission.AWS, err = GetAWSPermissions(result)
		if err != nil {
			log.Print(err)
		}
	case "oci", "digitalocean", "linode", "helm":
		log.Printf("Provider %s not yet implemented", result.Provider)
		return myPermission, nil
	case "azurerm", "azuread":
		myPermission.AZURE, err = GetAZUREPermissions(result)
		if err != nil {
			log.Print(err)
		}
	case "gcp", "google":
		myPermission.GCP, err = GetGCPPermissions(result)
		if err != nil {
			log.Print(err)
		}
	case "provider", "random", "main", "ip", "http", "test", "local",
		"archive", "tls", "template", "null", "time", "external":
		return myPermission, nil
	default:
		if result.Provider != "" && !(result.TypeName == "module") {
			log.Printf("Provider %s was not found", result.Provider)
		}
	}

	return myPermission, err
}

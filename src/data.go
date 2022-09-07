package pike

import (
	"io"
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
	src, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	parser := hclparse.NewParser()
	parsedFile, fileDiags := parser.ParseHCL(src, file)

	if fileDiags != nil {
		return nil, fileDiags
	}

	temp := parsedFile.Body.(*hclsyntax.Body)

	for _, block := range temp.Blocks {
		var resource ResourceV2
		resource.TypeName = block.Type

		ignore := []string{"terraform", "output", "provider", "variable", "locals", "template"}

		if stringInSlice(resource.TypeName, ignore) {
			continue
		}

		if strings.Contains(resource.TypeName, "module") {
			modulePath := GetModulePath(block)
			_, err := os.Stat(modulePath)
			if err != nil {
				//probably a directory or path
			} else {
				name := block.Labels[0]
				outPath := filepath.Join(dirName, "/.local_modules/", name)
				modulePath = filepath.Join(dirName, "/", modulePath)

				os.Remove(outPath)
				err := os.MkdirAll(outPath, 0750)
				files, _ := os.ReadDir(modulePath)
				for _, src := range files {
					sourceFile := filepath.Join(modulePath, src.Name())
					dst := outPath + "/" + src.Name()
					sourceFileStat, err := os.Stat(sourceFile)
					if err != nil {
						return nil, err
					}

					if !sourceFileStat.Mode().IsRegular() {
						//not copying directories
						continue
					}
					source, err := os.Open(sourceFile)

					if err != nil {
						return nil, err
					}
					defer source.Close()

					destination, err := os.Create(dst)
					if err != nil {
						return nil, err
					}
					defer destination.Close()
					_, err = io.Copy(destination, source)
					if err != nil {
						return nil, err
					}
				}
				if err != nil {
					return nil, err
				}

				//now process these extras
				ExtraFiles, _ := GetTF(outPath)
				for _, file := range ExtraFiles {
					resource, _ := GetResources(file, dirName)
					Resources = append(Resources, resource...)
				}
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

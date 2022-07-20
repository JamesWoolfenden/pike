package pike

import (
	"io/fs"
	"io/ioutil"
	"log"
	"path/filepath"
)

// Scan looks for resources in a given directory
func Scan(dirname string) error {

	files, err2 := GetTF(dirname)
	if err2 != nil {
		return err2
	}

	var results []template

	for _, file := range files {

		resources := GetResources(file, dirname)

		for _, resource := range resources {
			hcltype := GetHCLType(resource)
			provider := GetProvider(resource.name)
			result := template{resource, provider, hcltype}
			results = append(results, result)
		}
	}
	var PermissionBag []interface{}

	for _, result := range results {
		PermissionBag = append(PermissionBag, GetPermission(result)...)
	}

	for _, action := range PermissionBag {
		log.Print(action)
	}
	return nil
}

// GetTF return tf files in a directory
func GetTF(dirname string) ([]fs.FileInfo, error) {
	rawfiles, err := ioutil.ReadDir(dirname)
	var files []fs.FileInfo
	for _, file := range rawfiles {
		if file.IsDir() {
			continue
		}
		fileExtension := filepath.Ext(file.Name())

		if fileExtension != ".tf" {
			continue
		}
		files = append(files, file)
	}

	if err != nil {
		return nil, err
	}
	return files, nil
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// GetHCLType gets the template type
func GetHCLType(hcl Resource) string {
	return hcl.code.Keys[0].Token.Text
}

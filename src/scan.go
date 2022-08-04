package pike

import (
	"io/ioutil"
	"log"
	"path/filepath"
)

// Scan looks for resources in a given directory
func Scan(dirname string) error {
	fullPath, err := filepath.Abs(dirname)

	if err != nil {
		return err
	}

	files, err2 := GetTF(fullPath)
	if err2 != nil {
		return err2
	}

	var results []template

	for _, file := range files {

		resources, err := GetResources(file)

		if err != nil {
			//parse the other files
			log.Print(err)
		}

		for _, resource := range resources {
			hcltype := GetHCLType(resource)
			provider := GetProvider(resource.name)
			result := template{resource, provider, hcltype}
			results = append(results, result)
		}
	}
	var PermissionBag Sorted

	for _, result := range results {
		newPerms, err := GetPermission(result)

		if err != nil {
			return err
		}

		PermissionBag.AWS = append(PermissionBag.AWS, newPerms.AWS...)
		PermissionBag.GCP = append(PermissionBag.GCP, newPerms.AWS...)
	}

	err = GetPolicy(PermissionBag)

	return err
}

// GetTF return tf files in a directory
func GetTF(dirname string) ([]string, error) {
	rawFiles, err := ioutil.ReadDir(dirname)
	var files []string
	for _, file := range rawFiles {
		if file.IsDir() {

			if file.Name() == ".terraform" || file.Name() == ".git" {
				continue
			}
			newdirName := dirname + "/" + file.Name()
			moreFiles, err := GetTF(newdirName)
			if err == nil {
				if moreFiles != nil {
					files = append(files, moreFiles...)
				}
			}
		}

		fileExtension := filepath.Ext(file.Name())

		if fileExtension != ".tf" {
			continue
		}
		files = append(files, dirname+"/"+file.Name())
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

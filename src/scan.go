package pike

import (
	"io/ioutil"
	"log"
)

// Scan looks for resources in a given directory
func Scan(dirname string) error {
	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		return err
	}

	var results []template
	ignores := []string{".terraform", ".terraform.info.hcl", "terraform.tfstate", "terraform.tfstate.backup"}

	for _, file := range files {
		if stringInSlice(file.Name(), ignores) {
			continue
		}
		resources := GetResources(file, dirname)

		for _, resource := range resources {
			myAPI := GetAPI(resource.name)
			provider := GetProvider(resource.name)
			result := template{resource, myAPI, provider}
			results = append(results, result)
		}
	}
	var PermissionBag []string

	for _, result := range results {
		if result.API != "" {
			PermissionBag = append(PermissionBag, GetPermission(result)...)
		}
		// this now contains all the information we need for each resource/data/provider in our code
	}

	log.Print(PermissionBag)
	return nil
}

// Permission object probably delete this
type Permission struct {
}

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

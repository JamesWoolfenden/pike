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

	for _, file := range files {

		resources := GetResources(file, dirname)

		for _, resource := range resources {
			myAPI := GetAPI(resource.name)
			provider := GetProvider(resource.name)
			result := template{resource, myAPI, provider}
			results = append(results, result)
		}
	}

	for _, result := range results {
		var PermissionBag []Permission
		// this now contains all the information we need for each resource/data/provider in our code
		PermissionBag = append(PermissionBag, GetPermission(result))
		log.Print(PermissionBag)
	}

	return nil
}

// Permission object probably delete this
type Permission struct {
}

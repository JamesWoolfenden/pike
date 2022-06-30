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

		resources, filename, code := GetResources(file, dirname)
		for _, resource := range resources {
			myAPI := GetAPI(resource)
			provider := GetProvider(resource)
			result := template{resource, myAPI, filename, provider, code}
			results = append(results, result)
		}
	}

	log.Print(results)
	return nil
}

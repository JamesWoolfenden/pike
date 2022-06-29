package pike

import (
	"io/ioutil"
	"log"
)

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
			result := template{resource, myAPI, filename, code}
			results = append(results, result)
		}
	}

	log.Print(results)
	return nil
}

type template struct {
	resource string
	api      string
	filename string
	code     string
}

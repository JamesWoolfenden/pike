package main

import (
	"github.com/jameswoolfenden/pike/src"
	"io/ioutil"
	"log"
)

func main() {
	dirname := "./terraform/"
	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	var results []template

	for _, file := range files {

		resources, filename, code := pike.GetResources(file, dirname)
		for _, resource := range resources {
			myAPI := pike.GetAPI(resource)
			result := template{resource, myAPI, filename, code}
			results = append(results, result)
		}
	}

	log.Print(results)
}

type template struct {
	resource string
	api      string
	filename string
	code     string
}

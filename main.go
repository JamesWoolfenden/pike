package main

import (
	"io/fs"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	dirname := "./terraform/"
	files, err := ioutil.ReadDir(dirname)

	if err != nil {
		log.Fatal(err)
	}

	var results []template

	for _, file := range files {

		resources, filename, code := getResources(file, dirname)
		for _, resource := range resources {
			myAPI := getAPI(resource)
			result := template{resource, myAPI, filename, code}
			results = append(results, result)
		}
	}

	log.Print(results)
}

func getAPI(resource string) string {
	m := map[string]string{
		"aws_instance":  "ec2",
		"aws_s3_bucket": "s3"}

	api := m[resource]
	if api == "" {
		log.Printf("Resource:%s does not map", resource)
		return ""
	}

	return api
}

func getResources(file fs.FileInfo, dirname string) ([]string, string, string) {
	var result []string

	src, _ := ioutil.ReadFile(dirname + file.Name())
	resources := strings.Split(string(src), "}")

	for _, resource := range resources {
		if len(resource) > 1 {
			s := strings.Split(resource, "\"")
			verb, template := strings.TrimSpace(s[0]), strings.TrimSpace(s[1])
			if verb == "resource" {
				result = append(result, template)
			}
		}
	}

	return result, dirname + file.Name(), string(src)
}

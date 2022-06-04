package main

import (
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	//parser := hclparse.NewParser()
	//results, parseDiags := parser.ParseHCLFile()
	src, _ := ioutil.ReadFile("./terraform/aws_s3_bucket.pail.tf")
	resources := strings.Split(string(src), "}")

	for _, resource := range resources {
		if len(resource) > 1 {
			mine := (strings.Split(resource, "\""))[1]
			log.Print(mine)
		}
	}

}

package pike

import (
	_ "embed"
	"encoding/json"
	"log"
	"reflect"
	"strings"

	"github.com/hashicorp/hcl/hcl/ast"
)

// Mapping structure for permissions
type Mapping struct {
	Apply      []string   `json:"apply"`
	Destroy    []string   `json:"destroy"`
	Plan       []string   `json:"plan"`
	Attributes Attributes `json:"attributes"`
}

// Attributes Object
//goland:noinspection GoSnakeCaseUsage
type Attributes struct {
	Tag                 []string `json:"tag"`
	Object_lock_enabled []string `json:"object_lock_enabled"`
}

func (m Mapping) GetPermissions(file []byte) []Mapping {
	var mappings []Mapping
	_ = json.Unmarshal(file, &mappings)

	return mappings
}

//go:embed s3_bucket.json
var s3 []byte

// get permissions for AWS resources
func GetAWSPermissions(result template) []string {
	switch result.Resource.name {
	case "aws_s3_bucket":
		var newer Mapping
		var permissions []string
		mappings := newer.GetPermissions(s3)
		myAttributes := GetAttributes(result)

		matches := findMatches(mappings, myAttributes)

		for _, match := range matches {
			permissions = GetAttributesValue(mappings, match)
		}

		moreperms := GetBaseValue(mappings)
		permissions = append(permissions, moreperms...)
		return permissions
	case "aws_instance":

	default:
		log.Fatalf("resource %s not found", result.Resource.name)
	}

	return nil
}

func findMatches(mappings []Mapping, myAttributes []string) []string {
	e := reflect.ValueOf(&mappings[0].Attributes).Elem()

	var Fields []string
	for i := 0; i < e.NumField(); i++ {
		Fields = append(Fields, strings.ToLower(e.Type().Field(i).Name))
	}

	var matches []string
	for _, field := range Fields {
		if contains(myAttributes, field) {
			matches = append(matches, field)
		}
	}
	return matches
}

// GetAttributes gets the name of the important attributes for this resource
func GetAttributes(result template) []string {
	temp := result.Resource.code.Val.(*ast.ObjectType)
	attributes := temp.List.Items
	var myAttributes []string
	for _, item := range attributes {
		mytemp := item.Keys
		myAttributes = append(myAttributes, mytemp[0].Token.Text)
	}
	return myAttributes
}

// GetAttributesValue gets the values of attributes
func GetAttributesValue(result []Mapping, somestring string) []string {

	e := reflect.ValueOf(result[0].Attributes)
	//e := reflect.ValueOf(&mappings[0].Attributes).Elem()
	//test1 := e.Type().Field(3)
	//log.Print(test1)
	for i := 0; i < e.NumField(); i++ {
		if somestring == strings.ToLower(e.Type().Field(i).Name) {
			return e.Field(i).Interface().([]string)
		}
	}

	return nil
}

// GetBasevalue gets the base permissions - plan, apply, destroy
func GetBaseValue(result []Mapping) []string {

	e := reflect.ValueOf(result[0])

	var myparams []string
	for i := 0; i < 3; i++ {
		//myresult := e.Type().Field(i)
		myparams = append(myparams, e.Field(i).Interface().([]string)...)
	}

	return myparams
}

func contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

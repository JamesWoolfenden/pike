package pike

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
	"log"
)

func Compare(directory string, arn string) error {

	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := iam.NewFromConfig(cfg)

	Version := GetVersion(client, arn)
	Policy, _ := GetPolicyVersion(client, arn, Version)

	iacPolicy, _ := MakePolicy(directory, "json")
	Sorted, _ := SortActions(iacPolicy)

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)
	_, err = CompareIAMPolicy(Policy, string(Sorted))

	if err != nil {
		log.Fatal(err)
	}

	return nil
}

func CompareIAMPolicy(Policy string, OldPolicy string) (bool, error) {

	differ := diff.New()
	d, err := differ.Compare([]byte(Policy), []byte(OldPolicy))

	if err != nil {
		return false, err
	}

	if d.Modified() {
		var aJson map[string]interface{}
		json.Unmarshal([]byte(Policy), &aJson)

		config := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		formatter := formatter.NewAsciiFormatter(aJson, config)
		diffString, err := formatter.Format(d)
		if err != nil {
			// No error can occur
		}
		fmt.Print(diffString)
		return true, nil
	}

	return false, nil
}

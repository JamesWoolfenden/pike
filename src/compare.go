package pike

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// Compare IAC codebase to AWS policy
func Compare(directory string, arn string, init bool) (bool, error) {
	var theSame bool
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := iam.NewFromConfig(cfg)

	Version := GetVersion(client, arn)
	Policy, _ := GetPolicyVersion(client, arn, Version)

	iacPolicy, err := MakePolicy(directory, "", init)

	if err != nil {
		return theSame, err
	}

	Sorted, err := SortActions(iacPolicy.AWS.JSONOut)

	if err != nil {
		return theSame, err
	}

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)
	theSame, err = CompareIAMPolicy(Policy, Sorted)

	return theSame, err
}

// CompareIAMPolicy takes to IAm policies and compares
func CompareIAMPolicy(Policy string, OldPolicy string) (bool, error) {

	differ := diff.New()
	d, err := differ.Compare([]byte(Policy), []byte(OldPolicy))

	if err != nil {
		return false, err
	}

	if d.Modified() {
		var aJSON map[string]interface{}
		err = json.Unmarshal([]byte(Policy), &aJSON)

		if err != nil {
			return false, err
		}

		myConfig := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		myFormatter := formatter.NewAsciiFormatter(aJSON, myConfig)
		diffString, err := myFormatter.Format(d)

		if err != nil {
			return false, err
		}

		fmt.Print(diffString)

		return false, nil
	}

	return true, nil
}

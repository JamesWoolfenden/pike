package pike

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// Compare IAC codebase to AWS policy.
func Compare(directory string, arn string, init bool) (bool, error) {
	var theSame bool
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal().Err(err)
	}

	client := iam.NewFromConfig(cfg)

	Version, err := GetVersion(client, arn)
	if err != nil {
		return theSame, err
	}

	Policy, _ := GetPolicyVersion(client, arn, *Version)

	iacPolicy, err := MakePolicy(directory, nil, init, false)
	if err != nil {
		return theSame, err
	}

	Sorted, err := SortActions(iacPolicy.AWS.JSONOut)
	if err != nil {
		return theSame, err
	}

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)

	return CompareIAMPolicy(*Policy, *Sorted)
}

// CompareIAMPolicy takes two IAM policies and compares.
func CompareIAMPolicy(policy string, oldPolicy string) (bool, error) {
	differ := diff.New()
	compare, err := differ.Compare([]byte(policy), []byte(oldPolicy))

	if err != nil {
		return false, err
	}

	if compare.Modified() {
		var aJSON map[string]interface{}
		err = json.Unmarshal([]byte(policy), &aJSON)

		if err != nil {
			return false, err
		}

		myConfig := formatter.AsciiFormatterConfig{
			ShowArrayIndex: true,
			Coloring:       true,
		}

		myFormatter := formatter.NewAsciiFormatter(aJSON, myConfig)
		diffString, err := myFormatter.Format(compare)
		if err != nil {
			return false, err
		}

		fmt.Print(diffString)

		return false, nil
	}

	return true, nil
}

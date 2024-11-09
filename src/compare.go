package pike

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// Compare IAC codebase to AWS policy.
func Compare(directory string, arn string, init bool) (bool, error) {

	if directory == "" {
		return false, errors.New("directory cannot be empty")
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false, fmt.Errorf("directory does not exist: %s", directory)
	}

	if arn == "" {
		return false, errors.New("ARN cannot be empty")
	}

	if !strings.HasPrefix(arn, "arn:") {
		return false, errors.New("invalid ARN format")
	}

	var polciesMatch bool
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return false, fmt.Errorf("failed to load AWS config: %w", err)
	}

	client := iam.NewFromConfig(cfg)

	version, err := GetVersion(client, arn)
	if err != nil {
		return polciesMatch, err
	}

	policy, _ := GetPolicyVersion(client, arn, *version)

	iacPolicy, err := MakePolicy(directory, nil, init, false)
	if err != nil {
		return polciesMatch, err
	}

	sorted, err := SortActions(iacPolicy.AWS.JSONOut)
	if err != nil {
		return polciesMatch, err
	}

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)

	return CompareIAMPolicy(*policy, *sorted)
}

// CompareIAMPolicy takes two IAM policies and compares.
func CompareIAMPolicy(policy string, oldPolicy string) (bool, error) {
	differ := diff.New()
	compare, err := differ.Compare([]byte(policy), []byte(oldPolicy))

	if err != nil {
		return false, err
	}

	if compare.Modified() {
		return showDifferences(policy, compare)
	}

	return true, nil
}

func showDifferences(policy string, compare diff.Diff) (bool, error) {
	var aJSON map[string]interface{}
	err := json.Unmarshal([]byte(policy), &aJSON)

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

package pike

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	diff "github.com/yudai/gojsondiff"
	"github.com/yudai/gojsondiff/formatter"
)

// Compare IAC codebase to AWS policy.
func Compare(directory string, arn string, init bool) (bool, error) {

	valid, err := inputValidationCompare(directory, arn)
	if err != nil {
		return valid, &inputValidationError{err}
	}

	// Load the Shared AWS Configuration (~/.aws/config)
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)

	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)

	if err != nil {
		return false, &awsConfigError{err}
	}

	client := iam.NewFromConfig(cfg)

	version, err := GetVersion(client, arn)

	if err != nil {
		return false, &getVersionError{err}
	}

	policy, err := GetPolicyVersion(client, arn, *version)

	if err != nil {
		return false, &getPolicyVersionError{err}
	}

	iacPolicy, err := MakePolicy(directory, nil, init, false)

	if err != nil {
		return false, &getIAMVersionError{err}
	}

	sorted, err := SortActions(iacPolicy.AWS.JSONOut)

	if err != nil {
		return false, &sortActionsError{iacPolicy.AWS.JSONOut}
	}

	// iam versus iac
	fmt.Printf("IAM Policy %s versus Local %s \n", arn, directory)

	return CompareIAMPolicy(*policy, *sorted)
}

func inputValidationCompare(directory string, arn string) (bool, error) {
	if directory == "" {
		return false, &emptyDirectoryError{}
	}

	if _, err := os.Stat(directory); os.IsNotExist(err) {
		return false, &directoryNotFoundError{directory}
	}

	if arn == "" {
		return false, &arnEmptyError{}
	}

	if !strings.HasPrefix(arn, "arn:") {
		return false, &invalidARNError{arn}
	}

	return false, nil
}

// CompareIAMPolicy takes two IAM policies and compares.
func CompareIAMPolicy(policy string, oldPolicy string) (bool, error) {
	differ := diff.New()
	compare, err := differ.Compare([]byte(policy), []byte(oldPolicy))

	if err != nil {
		return false, err
	}

	if compare.Modified() {
		return ShowDifferences(policy, compare)
	}

	return true, nil
}

type formatterError struct {
	err error
}

func (m *formatterError) Error() string {
	return fmt.Sprintf("formatter failed: %v", m.err)
}

func ShowDifferences(policy string, compare diff.Diff) (bool, error) {
	var aJSON map[string]interface{}
	err := json.Unmarshal([]byte(policy), &aJSON)

	if err != nil {
		return false, &marshallPolicyError{err}
	}

	myConfig := formatter.AsciiFormatterConfig{
		ShowArrayIndex: true,
		Coloring:       true,
	}

	myFormatter := formatter.NewAsciiFormatter(aJSON, myConfig)
	diffString, err := myFormatter.Format(compare)

	if err != nil {
		return false, &formatterError{err}
	}

	fmt.Print(diffString)

	return false, nil
}

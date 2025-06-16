package pike

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"reflect"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/rs/zerolog/log"
)

const pollIntervalSeconds int = 5

// Watch looks at IAM policy for new revisions.
func Watch(arn string, wait int) error {
	if arn == "" {
		return &arnEmptyError{}
	}

	if wait <= 0 {
		return fmt.Errorf("wait time must be positive, got %d", wait)
	}

	if err := verifyAWSARN(arn); err != nil {
		return fmt.Errorf("invalid ARN format: %s", arn)
	}

	// Load the Shared AWS Configuration (~/.aws/config)
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)

	defer cancel()

	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return &awsConfigError{err}
	}

	client := iam.NewFromConfig(cfg)

	Version, err := getVersion(client, arn)
	if err != nil {
		return &getVersionError{err}
	}

	log.Info().Msgf("Waiting for change on policy Version %s", *Version)

	delay, err := waitForPolicyChange(client, arn, *Version, wait, pollIntervalSeconds) // Added default pollInterval of 10
	if err != nil {
		return &waitForPolicyChangeError{err}
	}

	log.Info().Msgf("Policy updated after %d", delay)

	return nil
}

// waitForPolicyChange looks at IAM policy change.
func waitForPolicyChange(client *iam.Client, arn string, version string, wait, pollInterval int) (int, error) {
	for item := 1; item < wait; item++ {
		time.Sleep(time.Duration(pollInterval))

		NewVersion, err := getVersion(client, arn)
		if err != nil {
			continue
		}

		if *NewVersion != version {
			return item, nil
		}

		log.Print("Not equal")
	}

	return wait, &waitExpiredError{}
}

type waitExpiredError struct{}

func (e *waitExpiredError) Error() string {
	return "wait expired with no change"
}

// getVersion gets the version of the IAM policy.
func getVersion(client *iam.Client, policyArn string) (*string, error) {
	output, err := client.GetPolicy(context.Background(), &iam.GetPolicyInput{PolicyArn: aws.String(policyArn)})

	if err != nil {
		return nil, &getVersionError{err}
	}

	return output.Policy.DefaultVersionId, nil
}

type urlEscapeError struct {
	err error
}

func (e *urlEscapeError) Error() string {
	return fmt.Sprintf("failed to unescape url: %v", e.err)
}

// getPolicyVersion Obtains the versioned IAM policy.
func getPolicyVersion(client *iam.Client, policyArn string, version string) (*string, error) {
	output, err := client.GetPolicyVersion(
		context.Background(),
		&iam.GetPolicyVersionInput{
			PolicyArn: aws.String(policyArn),
			VersionId: &version,
		})
	if err != nil {
		return nil, &getVersionError{err}
	}

	Policy, err := url.QueryUnescape(*(output.PolicyVersion.Document))
	if err != nil {
		return nil, &urlEscapeError{err}
	}

	fixed, err := sortActions(Policy)
	if err != nil {
		return nil, &sortActionsError{Policy}
	}

	return fixed, err
}

type castToListOfInterfaceError struct{}

func (e *castToListOfInterfaceError) Error() string {
	return "failed to convert to list of interfaces"
}

// sortActions sorts the actions list of an IAM policy.
func sortActions(myPolicy string) (*string, error) {
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(myPolicy), &raw)

	if err != nil {
		return nil, &unmarshallJSONError{err, myPolicy}
	}

	Statements, ok := raw["Statement"].([]interface{})

	if !ok {
		return nil, &castToListOfInterfaceError{}
	}

	var NewStatements []interface{}

	for _, block := range Statements {
		blocked, ok := block.(map[string]interface{})
		if !ok {
			log.Info().Msgf("assertion failed")
		}

		Actions := blocked["Action"]

		switch v := Actions.(type) {
		case string:
			// handle string case
		case []interface{}:
			theActions := sortInterfaceStrings(v)

			if theActions != nil {
				blocked["Action"] = theActions
			}
		default:
			log.Print(reflect.TypeOf(v).Kind())
		}

		NewStatements = append(NewStatements, block)
	}

	if NewStatements != nil {
		raw["Statement"] = NewStatements
	}

	fixed, err := json.Marshal(raw)

	if err != nil {
		return nil, &marshallPolicyError{err}
	}

	result := string(fixed)

	return &result, nil
}

func sortInterfaceStrings(actions interface{}) []string {
	temp, ok := actions.([]interface{})

	if !ok {
		log.Info().Msgf("failed to assert list for actions")

		return nil
	}

	myActions := make([]string, len(temp))

	for index, action := range temp {
		myAction, ok := action.(string)
		if !ok {
			log.Info().Msgf("failed to convert to string %s", action)

			continue
		}

		myActions[index] = myAction
	}

	sort.Strings(myActions)

	return myActions
}

type getVersionError struct {
	err error
}

func (e *getVersionError) Error() string {
	return fmt.Sprintf("failed to get version %v", e.err)
}

type waitForPolicyChangeError struct {
	err error
}

func (e *waitForPolicyChangeError) Error() string {
	return fmt.Sprintf("failed to wait for policy change %v", e.err)
}

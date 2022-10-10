package pike

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
	"reflect"
	"sort"
	"time"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
)

// Watch looks at IAM policy for new revisions
func Watch(arn string, wait int) error {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		return err
	}

	client := iam.NewFromConfig(cfg)

	Version, err := GetVersion(client, arn)

	if err != nil {
		return err
	}

	log.Printf("Waiting for change on policy Version %s", *Version)

	delay, err := WaitForPolicyChange(client, "arn:aws:iam::680235478471:policy/basic", *Version, wait)

	if err != nil {
		return err
	}

	log.Printf("Policy updated after %d", delay)
	return nil
}

// WaitForPolicyChange looks at IAM policy change
func WaitForPolicyChange(client *iam.Client, arn string, Version string, Wait int) (int, error) {

	for i := 1; i < Wait; i++ {
		time.Sleep(time.Duration(5))
		NewVersion, err := GetVersion(client, arn)
		if err != nil {
			continue
		}
		if NewVersion == &Version {
			return i, nil
		}
		log.Print("Not equal")
	}
	return Wait, errors.New("wait expired with no change")
}

// GetVersion gets the version of the IAM policy
func GetVersion(client *iam.Client, PolicyArn string) (*string, error) {

	output, err := client.GetPolicy(context.TODO(), &iam.GetPolicyInput{PolicyArn: aws.String(PolicyArn)})

	if err != nil {
		return nil, err
	}

	return output.Policy.DefaultVersionId, nil
}

// GetPolicyVersion Obtains the versioned IAM policy
func GetPolicyVersion(client *iam.Client, PolicyArn string, Version string) (*string, error) {
	output, err := client.GetPolicyVersion(context.TODO(), &iam.GetPolicyVersionInput{PolicyArn: aws.String(PolicyArn), VersionId: &Version})

	if err != nil {
		return nil, err
	}

	Policy, err := url.QueryUnescape(*(output.PolicyVersion.Document))
	if err != nil {
		return nil, err
	}

	fixed, err := SortActions(Policy)

	if err != nil {
		return nil, err
	}

	return fixed, err
}

// SortActions sorts the actions list of an IAM policy
func SortActions(myPolicy string) (*string, error) {
	//var raw Policy
	var raw map[string]interface{}
	err := json.Unmarshal([]byte(myPolicy), &raw)
	if err != nil {
		return nil, err
	}

	//var blocks []interface{}
	Statements := raw["Statement"].([]interface{})
	var NewStatements []interface{}
	for _, block := range Statements {

		blocked := block.(map[string]interface{})
		Actions := blocked["Action"]
		myType := reflect.TypeOf(Actions)
		switch myType.Kind() {
		case reflect.String:
			//do nothing
		case reflect.Slice:
			blocked["Action"] = sortInterfaceStrings(Actions)
		default:
			log.Print(myType.Kind())
		}
		NewStatements = append(NewStatements, block)
	}

	raw["Statement"] = NewStatements

	fixed, err := json.Marshal(raw)
	result := string(fixed)
	return &result, err
}

func sortInterfaceStrings(Actions interface{}) []string {
	temp := Actions.([]interface{})
	var myActions []string
	for _, action := range temp {
		myAction := action.(string)
		myActions = append(myActions, myAction)
	}
	sort.Strings(myActions)
	return myActions
}

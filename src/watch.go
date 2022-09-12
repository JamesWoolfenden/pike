package pike

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"net/url"
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

	Version := GetVersion(client, arn)

	log.Printf("Waiting for change on policy Version %s", Version)

	delay, err := WaitForPolicyChange(client, "arn:aws:iam::680235478471:policy/basic", Version, wait)

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
		NewVersion := GetVersion(client, arn)
		if NewVersion != Version {
			return i, nil
		}
	}
	return Wait, errors.New("wait expired with no change")
}

// GetVersion gets the version of the IAM policy
func GetVersion(client *iam.Client, PolicyArn string) string {

	output, err := client.GetPolicy(context.TODO(), &iam.GetPolicyInput{PolicyArn: aws.String(PolicyArn)})

	if err != nil {
		log.Fatal(err)
	}
	return *(output.Policy.DefaultVersionId)
}

// GetPolicyVersion Obtains the versioned IAM policy
func GetPolicyVersion(client *iam.Client, PolicyArn string, Version string) (string, error) {
	output, err := client.GetPolicyVersion(context.TODO(), &iam.GetPolicyVersionInput{PolicyArn: aws.String(PolicyArn), VersionId: &Version})

	if err != nil {
		return "", err
	}

	Policy, err := url.QueryUnescape(*(output.PolicyVersion.Document))
	if err != nil {
		return "", err
	}

	fixed, err := SortActions(Policy)

	if err != nil {
		return "", err
	}

	return fixed, err
}

// SortActions sorts the actions list of an IAM policy
func SortActions(myPolicy string) (string, error) {
	var raw Policy
	err := json.Unmarshal([]byte(myPolicy), &raw)
	if err != nil {
		return "", err
	}

	var blocks []Statement
	for _, block := range raw.Statements {
		sort.Strings(block.Action)
		blocks = append(blocks, block)
	}

	raw.Statements = blocks

	fixed, err := json.Marshal(raw)
	return string(fixed), err
}

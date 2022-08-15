package pike

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"log"
	"time"
)

func Watch(arn string, wait int) error {
	// Load the Shared AWS Configuration (~/.aws/config)
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	client := iam.NewFromConfig(cfg)

	Version := GetVersion(client, arn)

	log.Printf("Waiting for change on policy Version %s", Version)

	err, delay := WaitForPolicyChange(client, "arn:aws:iam::680235478471:policy/basic", Version, wait)

	if err != nil {
		return err
	}
	log.Printf("Policy updated after %d", delay)
	return nil
}

func WaitForPolicyChange(client *iam.Client, arn string, Version string, Wait int) (error, int) {

	for i := 1; i < Wait; i++ {
		time.Sleep(time.Duration(5))
		NewVersion := GetVersion(client, arn)
		if NewVersion != Version {
			return nil, i
		}
	}
	return errors.New("Wait expired with no change"), Wait
}

func GetVersion(client *iam.Client, PolicyArn string) string {

	output, err := client.GetPolicy(context.TODO(), &iam.GetPolicyInput{PolicyArn: aws.String(PolicyArn)})

	if err != nil {
		log.Fatal(err)
	}
	return *(output.Policy.DefaultVersionId)
}

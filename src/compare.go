package pike

import (
	"context"
	"errors"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/iam"
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
	log.Print(Version)
	iacPolicy, _ := MakePolicy(directory, "json")
	log.Print(iacPolicy)
	return errors.New("Not Implemented fully")
}

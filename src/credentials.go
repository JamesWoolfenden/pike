package pike

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
	"github.com/aws/aws-sdk-go-v2/service/sts/types"
)

const waitForConsistency = 900

type emptyRegionError struct{}

func (m emptyRegionError) Error() string {
	return "region cannot be empty"
}

type iamRoleEmptyError struct{}

func (m iamRoleEmptyError) Error() string {
	return "iamRole cannot be empty"
}

func getAWSCredentials(iamRole string, region string) (*sts.AssumeRoleOutput, error) {
	if iamRole == "" {
		return nil, &iamRoleEmptyError{}
	}

	if region == "" {
		return nil, &emptyRegionError{}
	}

	ctx := context.Background()
	cfg, err := config.LoadDefaultConfig(ctx, config.WithRegion(region))
	if err != nil {
		return nil, fmt.Errorf("failed to load AWS config: %w", err)
	}

	svc := sts.NewFromConfig(cfg)
	duration := int32(waitForConsistency)

	input := &sts.AssumeRoleInput{
		ExternalId:      aws.String("123ABC"),
		Policy:          nil,
		RoleArn:         aws.String(iamRole),
		RoleSessionName: aws.String("testAssumeRoleSession"),
		DurationSeconds: &duration,
		Tags:            []types.Tag{},
		TransitiveTagKeys: []string{
			"Project",
			"Cost-Center",
		},
	}

	result, err := svc.AssumeRole(ctx, input)
	if err != nil {
		var mpde *types.MalformedPolicyDocumentException
		var pptl *types.PackedPolicyTooLargeException
		var rde *types.RegionDisabledException
		var ete *types.ExpiredTokenException

		switch {
		case errors.As(err, &mpde):
			fmt.Println("MalformedPolicyDocumentException:", err.Error())
		case errors.As(err, &pptl):
			fmt.Println("PackedPolicyTooLargeException:", err.Error())
		case errors.As(err, &rde):
			fmt.Println("RegionDisabledException:", err.Error())
		case errors.As(err, &ete):
			fmt.Println("ExpiredTokenException:", err.Error())
		default:
			fmt.Println(err.Error())
		}

		return nil, err
	}

	return result, nil
}

type getAWSCredentialsError struct {
	err error
}

func (e getAWSCredentialsError) Error() string {
	return fmt.Sprintf("failed to get AWS credentials: %v", e.err)
}

func setAWSAuth(iamRole string, region string) error {
	credentials, err := getAWSCredentials(iamRole, region)
	if err != nil {
		return &getAWSCredentialsError{err}
	}

	_ = os.Setenv("AWS_ACCESS_KEY_ID", *credentials.Credentials.AccessKeyId)
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", *credentials.Credentials.SecretAccessKey)
	_ = os.Setenv("AWS_SESSION_TOKEN", *credentials.Credentials.SessionToken)

	return nil
}

func unSetAWSAuth() {
	_ = os.Setenv("AWS_ACCESS_KEY_ID", "")
	_ = os.Setenv("AWS_SECRET_ACCESS_KEY", "")
	_ = os.Setenv("AWS_SESSION_TOKEN", "")
}

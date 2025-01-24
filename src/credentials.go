package pike

import (
	"errors"
	"fmt"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sts"
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

	config := aws.NewConfig()

	config.Region = &region
	mySession := session.Must(session.NewSession(config))
	svc := sts.New(mySession)
	duration := int64(waitForConsistency)

	input := &sts.AssumeRoleInput{
		ExternalId:      aws.String("123ABC"),
		Policy:          nil,
		RoleArn:         aws.String(iamRole),
		RoleSessionName: aws.String("testAssumeRoleSession"),
		DurationSeconds: &duration,
		Tags:            []*sts.Tag{},
		TransitiveTagKeys: []*string{
			aws.String("Project"),
			aws.String("Cost-Center"),
		},
	}

	result, err := svc.AssumeRole(input)
	if err != nil {
		var aerr awserr.Error
		if errors.As(err, &aerr) {
			switch aerr.Code() {
			case sts.ErrCodeMalformedPolicyDocumentException:
				fmt.Println(sts.ErrCodeMalformedPolicyDocumentException, aerr.Error())
			case sts.ErrCodePackedPolicyTooLargeException:
				fmt.Println(sts.ErrCodePackedPolicyTooLargeException, aerr.Error())
			case sts.ErrCodeRegionDisabledException:
				fmt.Println(sts.ErrCodeRegionDisabledException, aerr.Error())
			case sts.ErrCodeExpiredTokenException:
				fmt.Println(sts.ErrCodeExpiredTokenException, aerr.Error())
			default:
				fmt.Println(aerr.Error())
			}
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

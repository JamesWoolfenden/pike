package pike

import (
	"context"
	"fmt"
	"strings"

	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/sts"
)

// AWS Partition constants
const (
	PartitionStandard = "aws"
	PartitionChina    = "aws-cn"
	PartitionGovCloud = "aws-us-gov"
)

// ExtractPartitionFromARN extracts the AWS partition from an ARN.
// ARN format: arn:partition:service:region:account-id:resource-type/resource-id
// Examples:
//
//	arn:aws:iam::123456789012:role/MyRole -> "aws"
//	arn:aws-cn:iam::123456789012:role/MyRole -> "aws-cn"
//	arn:aws-us-gov:iam::123456789012:role/MyRole -> "aws-us-gov"
func ExtractPartitionFromARN(arn string) (string, error) {
	if arn == "" {
		return "", fmt.Errorf("ARN cannot be empty")
	}

	// ARN format: arn:partition:service:region:account-id:resource-type/resource-id
	parts := strings.Split(arn, ":")
	if len(parts) < 2 {
		return "", fmt.Errorf("invalid ARN format: %s", arn)
	}

	// parts[0] should be "arn"
	if parts[0] != "arn" {
		return "", fmt.Errorf("ARN must start with 'arn': %s", arn)
	}

	// parts[1] is the partition
	partition := parts[1]
	if partition == "" {
		return "", fmt.Errorf("partition cannot be empty in ARN: %s", arn)
	}

	return partition, nil
}

// IsNonCommercialPartition checks if the partition is non-commercial
func IsNonCommercialPartition(partition string) bool {
	return partition != PartitionStandard && partition != ""
}

// GetCallerPartition looks up the current caller identity's ARN via STS and
// extracts its AWS partition, so callers can detect non-commercial partitions
// (aws-cn, aws-us-gov) without hardcoding a region-to-partition mapping.
func GetCallerPartition(ctx context.Context) (string, error) {
	cfg, err := config.LoadDefaultConfig(ctx)
	if err != nil {
		return "", fmt.Errorf("failed to load AWS config: %w", err)
	}

	svc := sts.NewFromConfig(cfg)

	result, err := svc.GetCallerIdentity(ctx, &sts.GetCallerIdentityInput{})
	if err != nil {
		return "", fmt.Errorf("failed to get caller identity: %w", err)
	}

	return ExtractPartitionFromARN(*result.Arn)
}

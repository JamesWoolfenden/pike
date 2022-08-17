# Pike

![alt text](pike.jfif "Pike")

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/jameswoolfenden/pike/graphs/commit-activity)
[![Build Status](https://github.com/JamesWoolfenden/pike/workflows/CI/badge.svg?branch=master)](https://github.com/JamesWoolfenden/pike)
[![Latest Release](https://img.shields.io/github/release/JamesWoolfenden/pike.svg)](https://github.com/JamesWoolfenden/pike/releases/latest)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/JamesWoolfenden/pike.svg?label=latest)](https://github.com/JamesWoolfenden/pike/releases/latest)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D0.14.0-blue.svg)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![checkov](https://img.shields.io/badge/checkov-verified-brightgreen)](https://www.checkov.io/)
[![Github All Releases](https://img.shields.io/github/downloads/jameswoolfenden/pike/total.svg)](https://github.com/JamesWoolfenden/pike/releases)

Pike is a tool, to determine the minimum permissions required to run a TF/IAC run:

Still very much under active development, I intend it to:

- run in ci - limit external outbound connections
- run on a path or a file (path is done)
- determines permission drift (Compare verb now exists)
- least privilege enabler
- policy creator (This is done)
- test policy against environment (This is done see compare)

Pike currently supports Terraform and can support multiple providers,
but I am currently focused on AWS.
Feel free to submit PR or Issue, and then I'll take a look.

## Install

Download the latest binary here:

<https://github.com/JamesWoolfenden/pike/releases>

Install from code:

- Clone repo
- Run `go install`

## Usage

### Scan

To scan a directory of Terraform file:

```shell
./pike -d .\terraform\ scan
{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": [
            "ec2:MonitorInstances",
            "ec2:UnmonitorInstances",
            "ec2:DescribeInstances",
            "ec2:DescribeTags",
            "ec2:DescribeInstanceAttribute",
            "ec2:DescribeVolumes",
            "ec2:DescribeInstanceTypes",
            "ec2:RunInstances",
            "ec2:DescribeInstanceCreditSpecifications",
            "ec2:StopInstances",
            "ec2:StartInstances",
            "ec2:ModifyInstanceAttribute",
            "ec2:TerminateInstances",
            "s3:GetBucketObjectLockConfiguration",
            "s3:PutBucketObjectLockConfiguration",
            "s3:PutObjectLegalHold",
            "s3:PutObjectRetention",
            "s3:PutObject",
            "s3:DeleteBucket",
            "s3:CreateBucket",
            "s3:GetLifecycleConfiguration",
            "s3:GetBucketTagging",
            "s3:GetBucketWebsite",
            "s3:GetBucketLogging",
            "s3:ListBucket",
            "s3:GetAccelerateConfiguration",
            "s3:GetBucketVersioning",
            "s3:GetBucketAcl",
            "s3:GetBucketPolicy",
            "s3:GetReplicationConfiguration",
            "s3:GetObjectAcl",
            "s3:GetObject",
            "s3:GetEncryptionConfiguration",
            "s3:GetBucketRequestPayment",
            "s3:GetBucketCORS",
            "ec2:AuthorizeSecurityGroupIngress",
            "ec2:AuthorizeSecurityGroupEgress",
            "ec2:CreateSecurityGroup",
            "ec2:DescribeSecurityGroups",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DeleteSecurityGroup",
            "ec2:RevokeSecurityGroupEgress"
        ],
        "Resource": "*"
    }
}
```

You can also generate the policy as Terraform instead:

```bash
$pike -o terraform -d ../modules/aws/terraform-aws-activemq  scan
resource "aws_iam_policy" "terraformXVlBzgba" {
  name        = "terraformXVlBzgba"
  path        = "/"
  description = "Add Description"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "ec2:AuthorizeSecurityGroupEgress",
                "ec2:AuthorizeSecurityGroupIngress",
                "ec2:CreateNetworkInterface",
                "ec2:CreateNetworkInterfacePermission",
                "ec2:CreateSecurityGroup",
                "ec2:CreateTags",
                "ec2:DeleteNetworkInterface",
                "ec2:DeleteNetworkInterfacePermission",
                "ec2:DeleteSecurityGroup",
                "ec2:DeleteTags",
                "ec2:DescribeAccountAttributes",
                "ec2:DescribeInternetGateways",
                "ec2:DescribeNetworkInterfaces",
                "ec2:DescribeSecurityGroups",
                "ec2:DescribeSubnets",
                "ec2:DescribeVpcs",
                "ec2:DetachNetworkInterface",
                "ec2:RevokeSecurityGroupEgress",
                "ec2:RevokeSecurityGroupIngress"
            ],
            "Resource": "*"
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "kms:CreateKey",
                "kms:DescribeKey",
                "kms:EnableKeyRotation",
                "kms:GetKeyPolicy",
                "kms:GetKeyRotationStatus",
                "kms:ListResourceTags",
                "kms:ScheduleKeyDeletion",
                "kms:TagResource",
                "kms:UntagResource"
            ],
            "Resource": "*"
        },
        {
            "Sid": "VisualEditor2",
            "Effect": "Allow",
            "Action": [
                "mq:CreateBroker",
                "mq:CreateConfiguration",
                "mq:CreateTags",
                "mq:CreateUser",
                "mq:DeleteBroker",
                "mq:DeleteTags",
                "mq:DeleteUser",
                "mq:DescribeBroker",
                "mq:DescribeConfiguration",
                "mq:DescribeConfigurationRevision",
                "mq:DescribeUser",
                "mq:RebootBroker",
                "mq:UpdateBroker",
                "mq:UpdateConfiguration",
                "mq:UpdateUser"
            ],
            "Resource": "*"
        }
    ]
})
}
```

### Readme

Pike can now be used to update a projects README.md file:

./pike -o terraform -d ..\modules\aws\terraform-aws-activemq\ readme

This looks in the readme for the deliminators:

```html
<!-- BEGINNING OF PRE-COMMIT-PIKE DOCS HOOK -->
<!-- END OF PRE-COMMIT-PIKE DOCS HOOK -->
```

and replaces is either with json or Terraform like so:

```markdown
This is the policy required to build this project:

The Policy required is

{
    "Version": "2012-10-17",
    "Statement": {
        "Effect": "Allow",
        "Action": [
            "mq:CreateTags",
            "mq:DeleteTags",
            "ec2:DescribeInternetGateways",
            "ec2:DescribeAccountAttributes",
            "ec2:DescribeVpcs",
            "ec2:DescribeSubnets",
            "ec2:DescribeSecurityGroups",
            "ec2:CreateNetworkInterface",
            "ec2:CreateNetworkInterfacePermission",
            "ec2:DeleteNetworkInterfacePermission",
            "ec2:DetachNetworkInterface",
            "ec2:DeleteNetworkInterface",
            "mq:CreateBroker",
            "mq:DescribeBroker",
            "mq:DescribeUser",
            "mq:UpdateBroker",
            "mq:DeleteBroker",
            "mq:CreateConfiguration",
            "mq:UpdateConfiguration",
            "mq:DescribeConfiguration",
            "mq:DescribeConfigurationRevision",
            "mq:RebootBroker",
            "ec2:CreateTags",
            "ec2:DeleteTags",
            "ec2:CreateSecurityGroup",
            "ec2:DescribeNetworkInterfaces",
            "ec2:DeleteSecurityGroup",
            "ec2:RevokeSecurityGroupEgress",
            "kms:TagResource",
            "kms:UntagResource",
            "kms:EnableKeyRotation",
            "kms:CreateKey",
            "kms:DescribeKey",
            "kms:GetKeyPolicy",
            "kms:GetKeyRotationStatus",
            "kms:ListResourceTags",
            "kms:ScheduleKeyDeletion"
        ],
        "Resource": "*"
    }
}
```

You can see an example here <https://github.com/jamesWoolfenden/terraform-aws-activemq#policy>.

## Compare

Want to check your deployed IAM policy against your IAC requirement?

>$./pike -d ../modules/aws/terraform-aws-appsync -a arn:aws:iam::680235478471:policy/basic compare

```markdown
IAM Policy arn:aws:iam::680235478471:policy/basic versus Local ../modules/aws/terraform-aws-appsync
 {
   "Statement": [
     0: {
       "Action": [
-        0: "kinesisvideo:CreateStream"
+        0: "firehose:CreateDeliveryStream"
+        0: "firehose:CreateDeliveryStream"
+        1: "firehose:DeleteDeliveryStream"
+        2: "firehose:DescribeDeliveryStream"
+        3: "firehose:ListTagsForDeliveryStream"
+        4: "iam:AttachRolePolicy"
+        5: "iam:CreateRole"
+        6: "iam:DeleteRole"
+        7: "iam:DetachRolePolicy"
+        8: "iam:GetRole"
+        9: "iam:ListAttachedRolePolicies"
+        10: "iam:ListInstanceProfilesForRole"
+        11: "iam:ListRolePolicies"
+        12: "iam:PassRole"
+        13: "iam:TagRole"
+        14: "kms:CreateKey"
+        15: "kms:DescribeKey"
+        16: "kms:EnableKeyRotation"
+        17: "kms:GetKeyPolicy"
+        18: "kms:GetKeyRotationStatus"
+        19: "kms:ListResourceTags"
+        20: "kms:ScheduleKeyDeletion"
+        21: "logs:AssociateKmsKey"
+        22: "logs:CreateLogGroup"
+        23: "logs:DeleteLogGroup"
+        24: "logs:DeleteRetentionPolicy"
+        25: "logs:DescribeLogGroups"
+        26: "logs:DisassociateKmsKey"
+        27: "logs:ListTagsLogGroup"
+        28: "logs:PutRetentionPolicy"
+        29: "s3:CreateBucket"
+        30: "s3:DeleteBucket"
+        31: "s3:GetAccelerateConfiguration"
+        32: "s3:GetBucketAcl"
+        33: "s3:GetBucketCORS"
+        34: "s3:GetBucketLogging"
+        35: "s3:GetBucketObjectLockConfiguration"
+        36: "s3:GetBucketPolicy"
+        37: "s3:GetBucketPublicAccessBlock"
+        38: "s3:GetBucketRequestPayment"
+        39: "s3:GetBucketTagging"
+        40: "s3:GetBucketVersioning"
+        41: "s3:GetBucketWebsite"
+        42: "s3:GetEncryptionConfiguration"
+        43: "s3:GetLifecycleConfiguration"
+        44: "s3:GetObject"
+        45: "s3:GetObjectAcl"
+        46: "s3:GetReplicationConfiguration"
+        47: "s3:ListAllMyBuckets"
+        48: "s3:ListBucket"
+        49: "s3:PutBucketAcl"
+        50: "s3:PutBucketPublicAccessBlock"
+        51: "s3:PutEncryptionConfiguration"
+        52: "wafv2:CreateWebACL"
+        53: "wafv2:DeleteWebACL"
+        54: "wafv2:GetWebACL"
       ],
       "Effect": "Allow",
       "Resource": "*",
-      "Sid": ""
+      "Sid": "VisualEditor0"
     }
   ],
   "Version": "2012-10-17"
 }

```

## Help

```bash
./pike -h
NAME:
   pike - Generate IAM policy from your IAC code

USAGE:
   pike [global options] command [command options] [arguments...]

VERSION:
   v0.1.26

AUTHOR:
   James Woolfenden <support@bridgecrew.io>

COMMANDS:
   compare, c  policy comparison of deployed versus IAC
   readme, r   Looks in dir for a README.md and updates it with the Policy required to build the code
   scan, s     scan a directory for IAM code
   version, v  Outputs the application version
   watch, w    Waits for policy update
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --arn value, -a value        Policy identifier e.g. arn (default: "arn:aws:iam::680235478471:policy/basic") [%ARN%]
   --config FILE, -c FILE       Load configuration from FILE
   --directory value, -d value  Directory to scan (defaults to .) (default: ".")
   --help, -h                   show help (default: false)
   --output json, -o json       Output types e.g. json terraform [%OUTPUT%]
   --version, -v                print the version (default: false)
   --wait value, -W value       Time to wait for policy change (in tenths of seconds) (default: 100) [%WAIT%]

```

## Building

```go
go build
```

or

```Make
Make build
```

## Extending

Determine and Create IAM mapping file,
working out the permissions required for your resource:
e.g. *aws_security_group.json*

```json
[
  {
    "apply": [
      "ec2:CreateSecurityGroup",
      "ec2:DescribeSecurityGroups",
      "ec2:DescribeAccountAttributes",
      "ec2:DescribeNetworkInterfaces",
      "ec2:DeleteSecurityGroup",
      "ec2:RevokeSecurityGroupEgress"
    ],
    "attributes": {
      "ingress": [
        "ec2:AuthorizeSecurityGroupIngress",
        "ec2:AuthorizeSecurityGroupEgress"
      ],
      "tags": [
        "ec2:CreateTags",
        "ec2:DeleteTags"
      ]
    },
    "destroy": [
      "ec2:DeleteSecurityGroup"
    ],
    "modify": [],
    "plan": []
  }
]

```

### Add Import mapping file

Update **files.go** with:

```go
//go:embed aws_security_group.json
var securityGroup []byte
```

### Add to provider Scan

Once you have added the json import above you just need to update the lookup table,
so we can read it and get the permissions:

```go
func GetAWSResourcePermissions(result template) []interface{} {
    TFLookup := map[string]interface{}{
        "aws_s3_bucket":            awsS3Bucket,
        "aws_s3_bucket_acl":        awsS3BucketACL,
+         "aws_security_group":       awsSecurityGroup,

```

Also add an example Terraform file into the folder terraform/backups.

## Related Tools

<https://github.com/iann0036/iamlive>

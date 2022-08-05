# Pike

![alt text](pike.jfif "Pike")

[![Maintenance](https://img.shields.io/badge/Maintained%3F-yes-green.svg)](https://GitHub.com/jameswoolfenden/pike/graphs/commit-activity)
[![Build Status](https://github.com/JamesWoolfenden/pike/workflows/CI/badge.svg?branch=master)](https://github.com/JamesWoolfenden/pike)
[![Latest Release](https://img.shields.io/github/release/JamesWoolfenden/pike.svg)](https://github.com/JamesWoolfenden/pike/releases/latest)
[![GitHub tag (latest SemVer)](https://img.shields.io/github/tag/JamesWoolfenden/pike.svg?label=latest)](https://github.com/JamesWoolfenden/pike/releases/latest)
![Terraform Version](https://img.shields.io/badge/tf-%3E%3D0.14.0-blue.svg)
[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit&logoColor=white)](https://github.com/pre-commit/pre-commit)
[![checkov](https://img.shields.io/badge/checkov-verified-brightgreen)](https://www.checkov.io/)
[![Github All Releases](https://img.shields.io/github/downloads/jameswoolfenden/pike/total.svg)]()

Pike is a tool, to determine the minimum permissions required to run a TF/IAC run:
Still very much under development, I intend it to: 

- run in ci - limit external outbound connections
- run on a path or a file
- determines permission drift
- least privilege enabler
- policy creator
- test policy against environment

It currently supports Terraform and can support multiple providers but focus is only on AWS at present. 

## Current issues:

- The current HCL parser needs to be replaced with one that ignores variables. We dont care about them yet.

## Usage

Scan a directory of terraform file

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

```bash
$./pike -h
NAME:
   pike - Generate IAM policy from your IAC code

USAGE:
   pike [global options] command [command options] [arguments...]

COMMANDS:
   scan, s  scan
   help, h  Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --config FILE, -c FILE       Load configuration from FILE
   --directory value, -D value  Directory to scan
   --help, -h                   show help (default: false)


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

```go
func GetAWSResourcePermissions(result template) []interface{} {
    myAttributes := GetAttributes(result)
    var Permissions []interface{}
    switch result.Resource.name {
    case "aws_s3_bucket":
       Permissions = GetPermissionMap(s3, myAttributes)
    case "aws_instance":
       Permissions = GetPermissionMap(ec2raw, myAttributes)
+   case "aws_security_group":
+      Permissions = GetPermissionMap(securityGroup, myAttributes)
```

## Related Tools

<https://github.com/iann0036/iamlive>

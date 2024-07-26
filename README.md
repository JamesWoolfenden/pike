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
[![codecov](https://codecov.io/gh/JamesWoolfenden/pike/graph/badge.svg?token=S5SW3BHIQQ)](https://codecov.io/gh/JamesWoolfenden/pike)
[![OpenSSF Best Practices](https://www.bestpractices.dev/projects/7032/badge)](https://www.bestpractices.dev/projects/7032)

Pike is a tool to determine the minimum permissions required to run a TF/IAC run:

Pike currently supports Terraform and supports multiple providers (AWS, GCP, AZURE),
Azure is the newest with AWS having the most supported resources
<https://github.com/JamesWoolfenden/pike/tree/master/src/mapping>.
Feel free to submit PR or Issue if you find an issue or even better add new resources, and then I'll take a look at
merging it ASAP.

**CAVEAT** The outputs of this tool are your first step, if you have AWS, you can now generate resources partially, there are no conditions and even partial resources are wildcarded (for now).
(for AWS)
**best practice** would go further (and I am working on it as well), you will need to modify these permissions to the minimum required in your environment by adding these
restrictions, you can also deploy using short-lived credentials (using this tool or Vault) (in AWS so far), generating short-lived credentials for your build
and then remotely (REMOTE) supply and invoke your builds (INVOKE).

Ideally I would like to do this for you, but these policies are currently determined statically (QUICKER), and unrecorded intentions can be impossible to infer.

## Table of Contents

<!--toc:start-->
- [Pike](#pike)
  - [Table of Contents](#table-of-contents)
  - [Install](#install)
    - [MacOS](#macos)
    - [Windows](#windows)
    - [Docker](#docker)
  - [Usage](#usage)
    - [Scan](#scan)
    - [Output](#output)
    - [Make](#make)
    - [Invoke](#invoke)
    - [Inspect](#inspect)
    - [Apply](#apply)
    - [Remote](#remote)
    - [Readme](#readme)
    - [Pull](#pull)
  - [Compare](#compare)
  - [Help](#help)
  - [Building](#building)
  - [Extending](#extending)
    - [Add Import mapping file](#add-import-mapping-file)
    - [Add to provider Scan](#add-to-provider-scan)
  - [Related Tools](#related-tools)
<!--toc:end-->

## Install

Download the latest binary here:

<https://github.com/JamesWoolfenden/pike/releases>

Install from code:

- Clone repo
- Run `go install`

Install remotely:

```shell
go install  github.com/jameswoolfenden/pike@latest
```

### MacOS

```shell
brew tap jameswoolfenden/homebrew-tap
brew install jameswoolfenden/tap/pike
```

### Windows

I'm now using Scoop to distribute releases, it's much quicker to update and easier to manage than previous methods,
you can install scoop from <https://scoop.sh/>.

Add my scoop bucket:

```shell
scoop bucket add iac https://github.com/JamesWoolfenden/scoop.git
```

Then you can install a tool:

```bash
scoop install pike
```

### Docker

```shell
docker pull jameswoolfenden/pike
docker run --tty --volume /local/path/to/tf:/tf jameswoolfenden/pike scan -d /tf
```

<https://hub.docker.com/repository/docker/jameswoolfenden/pike>

## Usage

### Scan

To scan a directory containing Terraform files:

```shell
./pike scan -d .\terraform\
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
$pike scan -o terraform -d ../modules/aws/terraform-aws-activemq
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

And I am working on further enhancements to policy generation, if you have AWS auth installed:

```hcl
e:\pike scan -d . -i -e
9:13AM DBG terraform init at E:\Code\modules\aws\terraform-aws-activemq
9:13AM DBG downloaded ip
resource "aws_iam_policy" "terraform_pike" {
  name_prefix = "terraform_pike"
  path        = "/"
  description = "Pike Autogenerated policy from IAC"

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
            "Resource": [
                "arn:aws:ec2:eu-west-2:680235478471:*"
            ]
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "kms:CreateGrant"
            ],
            "Resource": [
                "arn:aws:kms:eu-west-2:680235478471:*"
            ]
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
            "Resource": [
                "arn:aws:mq:eu-west-2:680235478471:*"
            ]
        }
    ]
})
}
```

I'm not finished with this yet as I am also working on bringing in resource names into the policies.

### Output

If you select the -w flag, pike will write out the role/policy required to build your project into the .pike folder:

```bash
$pike scan -w -i -d .
2022/09/17 13:50:51 terraform init at .
2022/09/17 13:50:51 downloaded ip
```

The .pike folder will contain:

``` shell
aws_iam_role.terraform_pike.tf
pike.generated_policy.tf
```

Which you can deploy using terraform to create the role/policy to build your IAC project.

### Make

You can now deploy the policy you need directly (AWS only so far):

```bash
$pike make -d ../modules/aws/terraform-aws-apigateway/

2022/09/18 08:53:41 terraform init at ..\modules\aws\terraform-aws-apigateway\
2022/09/18 08:53:41 modules not found at ..\modules\aws\terraform-aws-apigateway\
2022/09/18 08:53:49 aws role create/updated arn:aws:iam::680235478471:role/terraform_pike_20220918071439382800000002
 arn:aws:iam::680235478471:role/terraform_pike_20220918071439382800000002
```

This new verb returns the ARN of the role created, and you can find the Terraform used in your .pike folder.

### Invoke

Invoke is currently for triggering GitHub actions, if supplied with the workflow (defaults to main.yaml), repository and
branch (defaults to main) flags, it will trigger the dispatch event.

You'll need to include the dispatch event in your workflow:

```yaml
on:
  workflow_dispatch:
  push:
    branches:
      - master
```

To authenticate the GitHub API you will need to set you GitHub Personal Access Token, as the environment variable
*GITHUB_TOKEN*

To Invoke a workflow, it is then:

```shell
pike invoke -workflow master.yml -branch master -repository JamesWoolfenden/terraform-aws-s3
```

I created Invoke to be used in tandem with the new remote command which supplies temporary credentials to a workflow.

**Note The GitHub API is rate limited, usually 5000 calls per hour.

```shell
pike make -d ./module/aws/terraform-aws-s3/example/examplea
```

### Apply

Apply is an extension to make and will apply the policy and role and use that role to create your infrastructure:

```shell
pike apply -d ./module/aws/terraform-aws-s3/example/examplea -region eu-west-2
```

It is intended for testing and developing the permissions for Pike itself

### Remote

Remote uses the core code of make and apply, to write temporary AWS credentials(only so far) into your workflow.

```shell
pike remote -d ./module/aws/terraform-aws-s3/example/examplea -region eu-west-2 -repository terraform-aws-s3
```

### Readme

Pike can now be used to update a projects README.md file:

./pike readme -o terraform -d ..\modules\aws\terraform-aws-activemq\

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

>$./pike compare -d ../modules/aws/terraform-aws-appsync -a arn:aws:iam::680235478471:policy/basic

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

## Pull

Pull adds the ability to work with Git repositories (thanks to **go-git**),
to output the required permissions in json or Terraform:

```bash
./pike  pull
NAME:
   pike pull - Clones remote repo and scans it using pike

USAGE:
   pike pull [command options] [arguments...]

OPTIONS:
   --directory value, -d value        Directory to scan (defaults to .) (default: ".")
   --destination value, --dest value  Where to clone repository (default: ".destination")
   --output json, -o json             Policy Output types e.g. json terraform (default: "terraform") [%OUTPUT%]
   --repository value, -r value       Repository url
   --init, -i                         Run Terraform init to download modules (default: false)
   --write, -w                        Write the policy output to a file at .pike (default: false)
   --help, -h                         show help

```

Like so:

```hcl
$ ./pike.exe pull -r https://github.com/JamesWoolfenden/terraform-aws-codebuild -i -d .
10:31PM INF .destination was not empty, removing
10:31PM INF git clone https://github.com/JamesWoolfenden/terraform-aws-codebuild .destination --recursive
10:31PM DBG terraform init at E:\Code\pike\.destination
10:31PM DBG modules not found at .destination
resource "aws_iam_policy" "terraform_pike" {
  name_prefix = "terraform_pike"
  path        = "/"
  description = "Pike Autogenerated policy from IAC"

  policy = jsonencode({
    "Version": "2012-10-17",
    "Statement": [
        {
            "Sid": "VisualEditor0",
            "Effect": "Allow",
            "Action": [
                "codebuild:BatchGetProjects",
                "codebuild:CreateProject",
                "codebuild:DeleteProject",
                "codebuild:UpdateProject"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "VisualEditor1",
            "Effect": "Allow",
            "Action": [
                "events:DeleteRule",
                "events:DescribeRule",
                "events:ListTagsForResource",
                "events:ListTargetsByRule",
                "events:PutRule",
                "events:PutTargets",
                "events:RemoveTargets"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "VisualEditor2",
            "Effect": "Allow",
            "Action": [
                "iam:AttachRolePolicy",
                "iam:CreatePolicy",
                "iam:CreateRole",
                "iam:DeletePolicy",
                "iam:DeleteRole",
                "iam:DeleteRolePolicy",
                "iam:DetachRolePolicy",
                "iam:GetPolicy",
                "iam:GetPolicyVersion",
                "iam:GetRole",
                "iam:GetRolePolicy",
                "iam:ListAttachedRolePolicies",
                "iam:ListInstanceProfilesForRole",
                "iam:ListPolicyVersions",
                "iam:ListRolePolicies",
                "iam:PassRole",
                "iam:PutRolePolicy",
                "iam:TagRole"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "VisualEditor3",
            "Effect": "Allow",
            "Action": [
                "kms:Decrypt"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "VisualEditor4",
            "Effect": "Allow",
            "Action": [
                "s3:CreateBucket",
                "s3:DeleteBucket",
                "s3:GetAccelerateConfiguration",
                "s3:GetBucketAcl",
                "s3:GetBucketCORS",
                "s3:GetBucketLogging",
                "s3:GetBucketObjectLockConfiguration",
                "s3:GetBucketPolicy",
                "s3:GetBucketPublicAccessBlock",
                "s3:GetBucketRequestPayment",
                "s3:GetBucketTagging",
                "s3:GetBucketVersioning",
                "s3:GetBucketWebsite",
                "s3:GetEncryptionConfiguration",
                "s3:GetLifecycleConfiguration",
                "s3:GetObject",
                "s3:GetObjectAcl",
                "s3:GetReplicationConfiguration",
                "s3:ListAllMyBuckets",
                "s3:ListBucket",
                "s3:PutBucketAcl",
                "s3:PutBucketLogging",
                "s3:PutBucketPublicAccessBlock",
                "s3:PutBucketVersioning",
                "s3:PutEncryptionConfiguration",
                "s3:PutLifecycleConfiguration"
            ],
            "Resource": [
                "*"
            ]
        },
        {
            "Sid": "VisualEditor5",
            "Effect": "Allow",
            "Action": [
                "ssm:AddTagsToResource",
                "ssm:DeleteParameter",
                "ssm:DescribeParameters",
                "ssm:GetParameter",
                "ssm:GetParameters",
                "ssm:ListTagsForResource",
                "ssm:PutParameter"
            ],
            "Resource": [
                "*"
            ]
        }
    ]
})
}

```

## Help

```bash
./pike -h
NAME:
   pike - Generate IAM policy from your IAC code

USAGE:
   pike [global options] command [command options]

VERSION:
   9.9.9

AUTHOR:
   James Woolfenden <james.woolfenden@gmail.com>

COMMANDS:
   apply, a    Create a policy and use it to instantiate the IAC
   compare, c  policy comparison of deployed versus IAC
   inspect, x  policy comparison of environment versus IAC
   invoke, i   Triggers a gitHub action specified with the workflow flag
   make, m     make the policy/role required for this IAC to deploy
   parse, p    Triggers a gitHub action specified with the workflow flag
   pull, l     Clones remote repo and scans it using pike
   readme, r   Looks in dir for a README.md and updates it with the Policy required to build the code
   remote, o   Create/Update the Policy and set credentials/secret for Github Action
   scan, s     scan a directory for IAM code
   version, v  Outputs the application version
   watch, w    Waits for policy update
   help, h     Shows a list of commands or help for one command

GLOBAL OPTIONS:
   --help, -h     show help
   --version, -v  print the version

```

## Building

```shell
go build
```

or

```Make
Make build
```

## Inspect

This new feature is in *beta*, and is not yet fully supported and currently only for AWS.
When Pike is run with inspect, it will scan your code and output a policy that is required to deploy the code, as normal,
but it will also detect the running IAM credentials.
It will then report on the overlap between the running credentials and the minimum policy.

This works with AWS IAM user, group and role/assumed role credentials.

```bash
./pike inspect -d terraform/aws
The following are over-permissive:
s3:*
s3-object-lambda:*
*
account:GetAccountInformation
aws-portal:*Billing
aws-portal:*PaymentMethods
aws-portal:*Usage
billing:GetBillingData
billing:GetBillingDetails
billing:GetBillingNotifications
billing:GetBillingPreferences

```

This currently uses a different AWS profile to run the scan - presently hardcoded to "basic",
which only has the following permissions:

```json
statement {
    effect = "Allow"
    actions = [
      "iam:ListUserPolicies",
      "iam:ListAttachedUserPolicies",
      "iam:ListRolePolicies",
      "iam:ListAttachedRolePolicies",
      "iam:ListGroupPolicies",
      "iam:ListAttachedGroupPolicies",
      "iam:GetPolicy",
      "iam:GetPolicyVersion",
      "iam:GetUserPolicy",
      "iam:GetRolePolicy",
      "iam:GetGroupPolicy",
      "iam:ListGroupsForUser"
    ]
    resources = ["*"]
  }
```

Expect this all to change and be configurable SOON.

## Extending

Determine and Create IAM mapping file ("./src/mapping"),
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

### How

Datasources are the easiest to start with, I have a script (resource.ps1 - add pwsh with **brew install --cask powershell**)
that creates a blank mapping file and tf
resource, but you've seen the example json file - make one without any entries.
You also need to create a minimal resource/datasource, that you are trying to figure out the permissions for, and place it in the correct dir
e.g../terraform/aws, I have a script for making a profile for the profile in the role directory.
You can then tf using the empty role against the resource/datasource with no permissions.
The debug output from the tf run will help you figure out the permissions you need to add to your basic role.
You then update your "basic" role.

Issues?
The providers don't always tell you want you need to add,
you will need to check the IAM docs and the online IAM policymakers.
Not all resource are as easy as others, anything that make/scripts CF internally.
Some roles require *Passrole* and *CreateLinkedRole* but won't say so. Trail and error

#### What about "attributes" ?

Some cloud providers require extra permissions depending on the attributes you add, this is how this is handled.
Build out your tf resources to cover all reasonable scenarios.

#### Eventual consistency

Some cloud providers follow this model which means your test IAM role will take time after you change it to be
changed, how long? This seems to vary on time of day and the resource. Whilst other providers like
Azure just take a long time for the TF to change.

### Add Import mapping file

Update **files.go** with:

```txt
//go:embed aws_security_group.json
var securityGroup []byte
```

### Add to provider Scan

Once you have added the json import above you just need to update the lookup table,
so we can read it and get the permissions:

```txt
func GetAWSResourcePermissions(result template) []interface{} {
    TFLookup := map[string]interface{}{
        "aws_s3_bucket":            awsS3Bucket,
        "aws_s3_bucket_acl":        awsS3BucketACL,
+         "aws_security_group":       awsSecurityGroup,

```

Also add an example Terraform file into the folder **terraform/<cloud>/backups** this helps test that all your
new code is picked up pby pike.

## Related Tools

<https://github.com/iann0036/iamlive>

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=jameswoolfenden/pike&type=Date)](https://star-history.com/#jameswoolfenden/pike&Date)

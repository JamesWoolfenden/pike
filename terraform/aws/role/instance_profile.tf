#resource "aws_iam_role" "iam_emr_profile_role" {
#  name               = "iam_emr_profile_role"
#  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json
#}
#
#resource "aws_iam_instance_profile" "emr_profile" {
#  name = "emr_profile"
#  role = aws_iam_role.iam_emr_profile_role.name
#}
#
#data "aws_iam_policy_document" "iam_emr_profile_policy" {
#  statement {
#    effect = "Allow"
#
#    actions = [
#      "cloudwatch:*",
#      "dynamodb:*",
#      "ec2:Describe*",
#      "elasticmapreduce:Describe*",
#      "elasticmapreduce:ListBootstrapActions",
#      "elasticmapreduce:ListClusters",
#      "elasticmapreduce:ListInstanceGroups",
#      "elasticmapreduce:ListInstances",
#      "elasticmapreduce:ListSteps",
#      "kinesis:CreateStream",
#      "kinesis:DeleteStream",
#      "kinesis:DescribeStream",
#      "kinesis:GetRecords",
#      "kinesis:GetShardIterator",
#      "kinesis:MergeShards",
#      "kinesis:PutRecord",
#      "kinesis:SplitShard",
#      "rds:Describe*",
#      "s3:*",
#      "sdb:*",
#      "sns:*",
#      "sqs:*",
#    ]
#
#    resources = ["*"]
#  }
#}
#
#resource "aws_iam_role_policy" "iam_emr_profile_policy" {
#  name   = "iam_emr_profile_policy"
#  role   = aws_iam_role.iam_emr_profile_role.id
#  policy = data.aws_iam_policy_document.iam_emr_profile_policy.json
#}
#
#data "aws_iam_policy_document" "ec2_assume_role" {
#  statement {
#    effect = "Allow"
#
#    principals {
#      type        = "Service"
#      identifiers = ["ec2.amazonaws.com"]
#    }
#
#    actions = ["sts:AssumeRole"]
#  }
#}
#
#output "instance_profile" {
#  value = aws_iam_instance_profile.emr_profile
#}

resource "aws_emr_managed_scaling_policy" "pike" {
  cluster_id = aws_emr_cluster.pike.id
  compute_limits {
    unit_type                       = "Instances"
    minimum_capacity_units          = 3
    maximum_capacity_units          = 10
    maximum_ondemand_capacity_units = 10
    maximum_core_capacity_units     = 10
  }

}

resource "aws_emr_cluster" "pike" {
  name          = "pike"
  release_label = "emr-5.30.0"
  service_role  = "arn:aws:iam::680235478471:role/emr_service"
  master_instance_group {
    instance_type = "m5.xlarge"
  }

  core_instance_group {
    instance_count = 1
    instance_type  = "m5.xlarge"
  }
  ec2_attributes {
    #    subnet_id                         = aws_subnet.main.id
    #    emr_managed_master_security_group = aws_security_group.allow_access.id
    #    emr_managed_slave_security_group  = aws_security_group.allow_access.id
    instance_profile = aws_iam_instance_profile.emr_profile.arn
  }
}

resource "aws_iam_instance_profile" "emr_profile" {
  name = "emr_profile"
  role = aws_iam_role.iam_emr_profile_role.name
}

data "aws_iam_policy_document" "iam_emr_profile_policy" {
  statement {
    effect = "Allow"

    actions = [
      "cloudwatch:*",
      "dynamodb:*",
      "ec2:Describe*",
      "elasticmapreduce:Describe*",
      "elasticmapreduce:ListBootstrapActions",
      "elasticmapreduce:ListClusters",
      "elasticmapreduce:ListInstanceGroups",
      "elasticmapreduce:ListInstances",
      "elasticmapreduce:ListSteps",
      "kinesis:CreateStream",
      "kinesis:DeleteStream",
      "kinesis:DescribeStream",
      "kinesis:GetRecords",
      "kinesis:GetShardIterator",
      "kinesis:MergeShards",
      "kinesis:PutRecord",
      "kinesis:SplitShard",
      "rds:Describe*",
      "s3:*",
      "sdb:*",
      "sns:*",
      "sqs:*",
    ]

    resources = ["*"]
  }
}

resource "aws_iam_role_policy" "iam_emr_profile_policy" {
  name   = "iam_emr_profile_policy"
  role   = aws_iam_role.iam_emr_profile_role.id
  policy = data.aws_iam_policy_document.iam_emr_profile_policy.json
}

data "aws_iam_policy_document" "ec2_assume_role" {
  statement {
    effect = "Allow"

    principals {
      type        = "Service"
      identifiers = ["ec2.amazonaws.com"]
    }

    actions = ["sts:AssumeRole"]
  }
}

resource "aws_iam_role" "iam_emr_profile_role" {
  name               = "iam_emr_profile_role"
  assume_role_policy = data.aws_iam_policy_document.ec2_assume_role.json
}

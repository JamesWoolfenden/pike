resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [

          "config:DeleteConfigRule",
          "config:DescribeComplianceByConfigRule",
          "config:DescribeConfigRules",
          "config:ListTagsForResource",
          "config:PutConfigRule",
          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeImages",
          "ec2:DescribeInstanceAttribute",
          "ec2:DescribeInstanceCreditSpecifications",
          "ec2:DescribeInstanceTypes",
          "ec2:DescribeInstances",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeTags",
          "ec2:DescribeVolumes",
          "ec2:ModifyInstanceAttribute",
          "ec2:RunInstances",
          "ec2:StartInstances",
          "ec2:StopInstances",
          "ec2:TerminateInstances",
          "logs:CreateLogGroup",
          "logs:DeleteLogGroup",
          "logs:DescribeLogGroups",
          "logs:ListTagsLogGroup",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",

          # aws_cloudwatch_log_delivery_source
          "logs:PutDeliverySource",
          "logs:DeleteDeliverySource",

          # aws_cloudwatch_log_group
          "logs:ListTagsForResource",

          # aws_cloudwatch_log_index_policy
          "logs:PutIndexPolicy",
          "logs:DeleteIndexPolicy",
          "logs:DescribeIndexPolicies",

          # aws_cloudwatch_log_delivery_destination
          "logs:PutDeliveryDestination",
          "logs:GetDeliveryDestination",
          "logs:DescribeDeliveryDestinations",
          "logs:DeleteDeliveryDestination",

          # aws_cloudwatch_log_anomaly_detector
          "logs:CreateLogAnomalyDetector",
          "logs:GetLogAnomalyDetector",
          "logs:DeleteLogAnomalyDetector",
          "logs:UpdateLogAnomalyDetector",

          # aws_cloudwatch_log_delivery_destination_policy
          "logs:PutDeliveryDestinationPolicy",
          "logs:GetDeliveryDestinationPolicy",
          "logs:DeleteDeliveryDestinationPolicy"
        ],
        "Resource" : [
          "*"
        ]
      }
    ]
  })
  tags = {
    pike      = "permissions"
    createdby = "JamesWoolfenden"
  }
}

resource "aws_iam_role_policy_attachment" "basic" {
  role       = aws_iam_role.basic.name
  policy_arn = aws_iam_policy.basic.arn
}

resource "aws_iam_user_policy_attachment" "basic" {
  # checkov:skip=CKV_AWS_40: By design
  user       = "basic"
  policy_arn = aws_iam_policy.basic.arn
}

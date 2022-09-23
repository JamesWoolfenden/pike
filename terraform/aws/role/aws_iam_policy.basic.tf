resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:DescribeAccountAttributes",

          "batch:CreateComputeEnvironment",
          "batch:DeleteComputeEnvironment",
          "batch:UpdateComputeEnvironment",
          "batch:DescribeComputeEnvironments",
          "batch:TagResource",
          "batch:UntagResource",
          // "ec2:*",
          "iam:PassRole",
          "ec2:DescribeSubnets",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeKeyPairs",
          "ec2:DescribeVpcs",
          "ec2:DescribeImages",
          "ec2:DescribeLaunchTemplates",
          "ec2:DescribeLaunchTemplateVersions",

          "ecs:DescribeClusters",
          "ecs:Describe*",
          "ecs:List*",

          "batch:CreateJobQueue",
          "batch:UpdateJobQueue",
          "batch:DeleteJobQueue",
          "batch:DescribeJobQueues",
          "batch:UpdateJobQueue",
          "batch:TagResource",
          "batch:UntagResource",
          "batch:DescribeSchedulingPolicies",

          "batch:TagResource",
          "batch:UntagResource",
          "batch:CreateSchedulingPolicy",
          "batch:DeleteSchedulingPolicy",
          "batch:UpdateSchedulingPolicy",


          "batch:TagResource",
          "batch:UntagResource",
          "batch:DeregisterJobDefinition",
          "batch:DescribeJobDefinitions",
          "batch:RegisterJobDefinition",

        ]
        "Resource" : "*"
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

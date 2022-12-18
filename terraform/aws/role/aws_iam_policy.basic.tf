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
          "SNS:ListTopics",

          "resource-groups:CreateGroup",
          "resource-groups:GetGroup",
          "resource-groups:GetGroupQuery",
          "resource-groups:GetTags",
          "resource-groups:DeleteGroup",
          "resource-groups:Untag",
          "resource-groups:Tag",
          "resource:UpdateGroup",

          "applicationinsights:CreateApplication",
          "applicationinsights:TagResource",
          "applicationinsights:UntagResource",
          "iam:CreateServiceLinkedRole",
          "logs:DescribeLogGroups",
          "applicationinsights:DescribeApplication",
          "applicationinsights:ListTagsForResource",
          "applicationinsights:DeleteApplication",

          "applicationinsights:UpdateApplication"
        ],
        "Resource" : "*",
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

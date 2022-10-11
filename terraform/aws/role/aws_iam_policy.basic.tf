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
          "eks:CreateCluster",
          "eks:DeleteCluster",
          "eks:TagResource",
          "eks:UntagResource",
          "eks:ListTagsForResource",
          "iam:PassRole",
          "eks:DescribeCluster",
          "eks:UpdateClusterConfig",

          "eks:CreateAddon",
          "eks:DescribeAddon",
          "eks:DescribeAddonVersions",
          "eks:UpdateAddon",
          "eks:DeleteAddon",
          "eks:TagResource",
          "eks:UntagResource",

          "ec2:DescribeAccountAttributes",
          "eks:CreateNodegroup",
          "eks:DeleteNodegroup",
          "eks:UpdateNodegroupConfig",
          "eks:TagResource",
          "eks:UntagResource",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "ec2:DescribeSubnets",
          "iam:CreateServiceLinkedRole",
          "eks:DescribeNodegroup"
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

resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "codeartifact:CreateDomain",
          "codeartifact:TagResource",
          "codeartifact:UntagResource",
          "codeartifact:DescribeDomain",
          "codeartifact:ListTagsForResource",
          "codeartifact:DeleteDomain",
          "codeartifact:PutDomainPermissionsPolicy",
          "codeartifact:GetDomainPermissionsPolicy",
          "codeartifact:DeleteDomainPermissionsPolicy",
          "codeartifact:CreateRepository",
          "codeartifact:DescribeRepository",
          "codeartifact:DeleteRepository",
          "codeartifact:UpdateRepository",
          "codeartifact:PutRepositoryPermissionsPolicy",
          "codeartifact:GetRepositoryPermissionsPolicy",
          "codeartifact:DeleteRepositoryPermissionsPolicy"
        ],
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

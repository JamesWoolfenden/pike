resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "mediaconvert:UpdateQueue",
          "mediaconvert:DeleteQueue",
          "mediaconvert:ListTagsForResource",
          "mediaconvert:DescribeEndpoints",
          "mediaconvert:UntagResource",
          "mediaconvert:CreateQueue",
          "ec2:DescribeAccountAttributes",
          "mediaconvert:TagResource",
          "mediaconvert:GetQueue",
          "iam:ListRole",
          "iam:PassRole",
          "s3:ListBucket",
          "s3:GetBucketLocation",
          "s3:ListAllMyBuckets"
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

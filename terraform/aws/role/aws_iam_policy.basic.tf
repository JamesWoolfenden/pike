resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_cloudfront_function
          "cloudfront:CreateFunction",
          "cloudfront:PublishFunction",
          "cloudfront:DescribeFunction",
          "cloudfront:GetFunction",
          "cloudfront:DeleteFunction",
          "cloudfront:UpdateFunction",

          //aws_cognito_identity_pool
          "cognito-identity:CreateIdentityPool",
          "cognito-identity:DescribeIdentityPool",
          "cognito-identity:TagResource",
          "cognito-identity:UntagResource",
          "cognito-identity:DeleteIdentityPool",
          "cognito-identity:UpdateIdentityPool",
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

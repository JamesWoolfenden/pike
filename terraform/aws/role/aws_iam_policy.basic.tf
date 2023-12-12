resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_timestreamwrite_database
          "timestream:DescribeEndpoints",
          "timestream:CreateDatabase",
          "timestream:TagResource",
          "timestream:UntagResource",
          "timestream:DescribeDatabase",
          "timestream:ListTagsForResource",
          "timestream:DeleteDatabase",
          "timestream:UpdateDatabase",
          "kms:CreateGrant",
          "kms:DescribeKey",
          "kms:ListKeys",
          #          "kms:EnableKey"
          #          "kms:Encrypt",
          #          "kms:Decrypt"


          "timestream:ListTagsForResource",
          "timestream:TagResource",
          "timestream:UntagResource",
          "timestream:CreateTable",
          "timestream:DeleteTable",
          "timestream:UpdateTable",
          "timestream:DescribeTable",

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

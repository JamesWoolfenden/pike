resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_db_proxy
          "rds:CreateDBProxy",
          "rds:DeleteDBProxy",
          "rds:DescribeDBProxies",
          "rds:ModifyDBProxy",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",



          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeSubnets",
          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "rds:CreateDBInstance",
          "rds:DeleteDBInstance",
          "rds:DescribeDBInstances",
          "rds:ListTagsForResource",
          "rds:ModifyDBInstance",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "secretsmanager:CreateSecret",
          "secretsmanager:DeleteSecret",
          "secretsmanager:DescribeSecret",
          "secretsmanager:GetResourcePolicy",
          "secretsmanager:TagResource",
          "secretsmanager:UntagResource"
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

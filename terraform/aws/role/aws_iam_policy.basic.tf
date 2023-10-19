resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_organizations_policy
          "organizations:DescribePolicy",
          "organizations:CreatePolicy",
          "organizations:DeletePolicy",
          "organizations:UpdatePolicy",
          "organizations:UntagResource",
          "organizations:TagResource",
          //aws_secretsmanager_secret_policy
          "secretsmanager:DescribeSecret",
          "secretsmanager:PutResourcePolicy",
          "secretsmanager:DeleteResourcePolicy",

          //aws_secretsmanager_secret
          "secretsmanager:GetResourcePolicy",
          //aws_secretsmanager_secret_rotation
          "secretsmanager:RotateSecret",
          "secretsmanager:CancelRotateSecret",
          "lambda:InvokeFunction",
          //
          "states:DescribeStateMachineAlias",
          "states:CreateStateMachineAlias",
          "states:DeleteStateMachineAlias",
          "states:UpdateStateMachineAlias"
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

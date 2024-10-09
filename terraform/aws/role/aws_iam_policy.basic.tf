resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          //aws_iam_user_policies_exclusive
          "iam:ListUserPolicies",
          "iam:PutUserPolicy",

          //aws_iam_role_policies_exclusive
          "iam:ListRolePolicies",
          "iam:PutRolePolicy",

          //aws_iam_group_policies_exclusive
          "iam:ListGroupPolicies",
          "iam:PutGroupPolicy",

          //aws_sagemaker_human_task_ui
          "sagemaker:CreateHumanTaskUi",
          "sagemaker:DescribeHumanTaskUi",
          "sagemaker:ListTags",
          "sagemaker:DeleteHumanTaskUi",
          "sagemaker:AddTags",
          "sagemaker:DeleteTags",

          //aws_memorydb_user
          "memorydb:CreateUser",
          "memorydb:DescribeUsers",
          "memorydb:ListTags",
          "memorydb:DeleteUser",
          "memorydb:TagResource",
          "memorydb:UntagResource",

          //aws_m2_environment
          "m2:CreateEnvironment",
          "m2:GetEnvironment",
          "m2:DeleteEnvironment",
          "m2:UpdateEnvironment",

          //aws_m2_deployment
          "m2:GetDeployment",
          "m2:CreateDeployment",

          //aws_m2_application
          "m2:GetApplication",
          "m2:CreateApplication",
          "m2:DeleteApplication",
          "m2:UpdateApplication",

          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor1",
        "Effect" : "Allow",
        "Action" : [
          "iam:CreateUser",
          "iam:DeleteUser",
          "iam:GetUser",
          "iam:ListGroupsForUser",
          "iam:CreateRole",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:DeleteRole",
          "iam:CreatePolicy",
          "iam:GetPolicy",
          "iam:GetPolicyVersion",
          "iam:ListPolicyVersions",
          "iam:DeletePolicy"
        ],
        "Resource" : [
          "*"
        ]
      },
      {
        "Sid" : "VisualEditor2",
        "Effect" : "Allow",
        "Action" : [
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject"
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

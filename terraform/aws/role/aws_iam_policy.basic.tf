resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "datazone:ListEnvironmentBlueprints",


          //aws_datazone_domain
          "datazone:CreateDomain",
          "datazone:GetDomain",
          "datazone:ListTagsForResource",
          "datazone:DeleteDomain",
          "datazone:UpdateDomain",

          "datazone:TagResource",
          "datazone:UntagResource",

          //aws_datazone_environment_blueprint_configuration
          "datazone:PutEnvironmentBlueprintConfiguration",
          "datazone:GetEnvironmentBlueprintConfiguration",
          "datazone:DeleteEnvironmentBlueprintConfiguration",

          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:GetRole",
          "iam:GetRolePolicy",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "iam:PutRolePolicy",
          "iam:DeleteRolePolicy",

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

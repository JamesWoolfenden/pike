resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "iam:EnableOrganizationsRootCredentialsManagement",
          "iam:DisableOrganizationsRootCredentialsManagement",

          "dynamodb:DeleteItem",
          "dynamodb:DescribeTable",
          "dynamodb:GetItem",
          "dynamodb:PutItem",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcs",
          "ec2:GetManagedPrefixListEntries",
          "grafana:CreateWorkspace",
          "grafana:DeleteWorkspace",
          "grafana:DescribeWorkspace",
          "grafana:DescribeWorkspaceAuthentication",
          "grafana:DescribeWorkspaceConfiguration",
          "grafana:UpdateWorkspace",
          "grafana:UpdateWorkspaceAuthentication",
          "grafana:UpdateWorkspaceConfiguration",
          "iam:CreateRole",
          "iam:CreateServiceLinkedRole",
          "iam:DeleteRole",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",
          "iam:PassRole",
          "organizations:DescribeOrganization",
          "s3:DeleteObject",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutObject",
          "sso:CreateManagedApplicationInstance",
          "sso:DeleteManagedApplicationInstance",
          "sso:DescribeRegisteredRegions",
          "sso:GetApplicationInstance",
          "sso:GetSharedSsoConfiguration",
          "sso:ListApplicationInstances",

          "grafana:UpdatePermissions"

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

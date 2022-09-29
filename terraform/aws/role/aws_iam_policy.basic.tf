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
          "ds:CreateDirectory",
          "ds:DeleteDirectory",
          "ds:DescribeDirectories",
          "ds:ListTagsForResource",
          "ec2:AuthorizeSecurityGroupEgress",
          "ec2:AuthorizeSecurityGroupIngress",
          "ec2:CreateNetworkInterface",
          "ec2:CreateSecurityGroup",
          "ec2:DeleteNetworkInterface",
          "ec2:DeleteSecurityGroup",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcs",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:RevokeSecurityGroupIngress",
          "ds:AddTagsToResource",
          "ds:RemoveTagsFromResource",
          "ec2:CreateTags",
          "ec2:DeleteTags",

          "ec2:DeleteSecurityGroup",
          "ds:DeleteDirectory",
          "ec2:DeleteNetworkInterface",
          "ec2:DeleteSecurityGroup",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:RevokeSecurityGroupIngress",
          "workspaces:DescribeWorkspaceBundles",

          "workspaces:RegisterWorkspaceDirectory",
          "workspaces:CreateTags",
          "workspaces:DeleteTags",

          "kms:DescribeKey",
          "kms:ListAliases",
          "kms:ListKeys",

          "workspaces:CreateWorkspaces",
          "workspaces:CreateWorkspaceImage",

          "workspaces:DescribeTags",
          "workspaces:DescribeWorkspaceBundles",
          "workspaces:DescribeWorkspaceDirectories",
          "workspaces:DescribeWorkspaces",
          "workspaces:DescribeWorkspacesConnectionStatus",

          "workspaces:ModifyWorkspaceProperties",
          "workspaces:ModifySamlProperties",

          "workspaces:RebootWorkspaces",
          "workspaces:RebuildWorkspaces",
          "workspaces:StartWorkspaces",
          "workspaces:StopWorkspaces",
          "workspaces:TerminateWorkspaces",
          "workspaces:DeregisterWorkspaceDirectory",


          "ec2:CreateVpc",
          "ec2:CreateSubnet",
          "ec2:CreateInternetGateway",
          "ec2:CreateRouteTable",
          "ec2:CreateRoute",
          "ec2:DescribeInternetGateways",
          "ec2:DescribeRouteTables",
          "ec2:DescribeAvailabilityZones",
          "ec2:AttachInternetGateway",
          "ec2:AssociateRouteTable",

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

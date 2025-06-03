resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "ec2:AssociateSubnetCidrBlock",
          "ec2:CreateSecurityGroup",
          "ec2:CreateSubnet",
          "ec2:CreateVPC",
          "ec2:DeleteSecurityGroup",
          "ec2:DeleteSubnet",
          "ec2:DeleteVPC",
          "ec2:DescribeAccountAttributes",
          "ec2:DescribeNetworkAcls",
          "ec2:DescribeNetworkInterfaces",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcAttribute",
          "ec2:DescribeVpcs",
          "ec2:DisassociateSubnetCidrBlock",
          "ec2:ModifySubnetAttribute",
          "ec2:ModifyVpcAttribute",
          "ec2:ModifyVpcTenancy",
          "ec2:RevokeSecurityGroupEgress",
          "ec2:DescribeAvailabilityZones",
          "kms:CreateKey",
          "kms:DescribeKey",
          "kms:GetKeyPolicy",
          "kms:GetKeyRotationStatus",
          "kms:ListResourceTags",
          "kms:ScheduleKeyDeletion",
          "kms:PutKeyPolicy",


          //aws_quicksight_account_settings
          "quicksight:UpdateAccountSettings",

          //aws_workspacesweb_browser_settings
          "workspaces-web:CreateBrowserSettings",
          "workspaces-web:GetBrowserSettings",
          "workspaces-web:ListTagsForResource",
          "workspaces-web:DeleteBrowserSettings",
          "workspaces-web:TagResource",
          "workspaces-web:UntagResource",

          //aws_workspacesweb_network_settings
          "workspaces-web:CreateNetworkSettings",
          "iam:CreateServiceLinkedRole",
          "workspaces-web:GetNetworkSettings",
          "workspaces-web:DeleteNetworkSettings",

          //aws_workspacesweb_user_settings
          "workspaces-web:CreateUserSettings",
          "workspaces-web:TagResource",
          "workspaces-web:UntagResource",
          "workspaces-web:GetUserSettings",
          "workspaces-web:DeleteUserSettings",
          "kms:CreateGrant",
          "kms:Decrypt",
          "kms:GenerateDataKey"

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

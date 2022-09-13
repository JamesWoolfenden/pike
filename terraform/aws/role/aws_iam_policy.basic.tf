resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "redshift:CreateTags",
          "redshift:DeleteTags",
          "redshift:DeleteCluster",
          //"redshift:CreateEndpointAccess",
          //"redshift:RevokeEndpointAccess"
          //"redshift:DescribeClusters"

          #          "redshift:CreateTags",
          #          "redshift:DeleteTags",
          "redshift:CreateCluster",
          "redshift:ModifyCluster",
          "redshift:DescribeClusters",
          "redshift:DescribeLoggingStatus",

          //logging
          "redshift:DisableLogging",
          "redshift:EnableLogging",

          "redshift:ModifyClusterIamRoles",
          #          "redshift:DescribeClusters"
          #          "redshift:CreateScheduledAction",
          //"redshift:PauseCluster",
          //iam_roles
          "iam:PassRole",
          #          "redshift:DescribeScheduledActions",
          #          "redshift:DeleteScheduledAction",
          #          "redshift:ModifyScheduledAction"
          //"redshift:CreateClusterUser"

          "ec2:DescribeAccountAttributes",
          "ec2:DescribeAddresses",
          "ec2:DescribeAvailabilityZones",
          "ec2:DescribeSecurityGroups",
          "ec2:DescribeSubnets",
          "ec2:DescribeVpcs",
          "ec2:DescribeInternetGateways",



          "redshift:CreateEndpointAccess"
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

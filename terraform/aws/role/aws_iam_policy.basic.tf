resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          "aps:CreateWorkspace",
          "aps:DeleteWorkspace",
          "aps:DescribeWorkspace",
          "aps:ListTagsForResource",
          "aps:UpdateWorkspaceAlias",
          "s3express:CreateBucket",
          "s3express:DeleteBucket",
          "s3:CreateAccessPoint",
          "s3:DeleteAccessPoint",
          "s3:DeleteAccessPointPolicy",
          "s3:DeleteObject",
          "s3:GetAccessPoint",
          "s3:GetAccessPointPolicy",
          "s3:GetObject",
          "s3:ListBucket",
          "s3:PutAccessPointPolicy",
          "s3:PutAccessPointPublicAccessBlock",
          "s3:PutObject",
          "ec2:DescribeAvailabilityZones",



          //aws_lightsail_bucket
          "lightsail:CreateBucket",
          "lightsail:DeleteBucket",
          "lightsail:GetBuckets",
          "lightsail:GetInstance",
          "lightsail:SetResourceAccessForBucket",
          "lightsail:UpdateBucket",
          "lightsail:UpdateBucketBundle",


          //aws_prometheus_workspace_configuration
          "aps:UpdateWorkspaceConfiguration",
          "aps:DescribeWorkspaceConfiguration",

          //aws_prometheus_workspace
          "aps:DescribeLoggingConfiguration",



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

resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_dms_replication_task
          "dms:CreateReplicationTask",
          //aws_dms_s3_endpoint
          "dms:CreateEndpoint",
          "dms:DeleteEndpoint",
          "rds:DescribeDBProxies",
          "dms:ModifyEndpoint",
          "iam:PassRole",
          "rds:RemoveTagsFromResource",
          "oam:UnTagResource",


          "ec2:DescribeCapacityBlockOfferings",
          "acm:ListCertificates",
          "oam:GetLink",
          "iam:AttachRolePolicy",
          "ec2:DescribeSubnets",
          "oam:UpdateLink",
          "oam:DeleteLink",
          "iam:DetachRolePolicy",

          //aws_oam_sink_policy
          "iam:GetRolePolicy",
          "iam:PutRolePolicy",
          "rds:DeleteDBInstance",
          "kms:DescribeKey",
          "rds:ListTagsForResource",
          "rds:ModifyDBInstance",
          "sns:CreateTopic",
          "sns:DeleteTopic",
          "secretsmanager:DescribeSecret",
          "sns:GetTopicAttributes",
          "sns:ListTagsForResource",
          "sns:SetTopicAttributes"
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

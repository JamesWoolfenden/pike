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
          "rds:CreateDBCluster",
          "rds:DescribeDBClusters",
          "rds:ListTagsForResource",
          "rds:DeleteDBCluster",

          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",

          "rds:AddRoleToDBCluster",
          "iam:PassRole",

          "rds:ModifyDBCluster",

          "ec2:DescribeAccountAttributes",
          "rds:CreateDBInstance",
          "rds:DescribeDBInstances",
          "rds:DeleteDBInstance",
          "rds:ModifyDBInstance",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",

          "ec2:DescribeAccountAttributes",
          "rds:CreateDBClusterEndpoint",
          "rds:DeleteDBClusterEndpoint",
          "rds:ModifyDBClusterEndpoint",
          "rds:DescribeDBClusterEndpoints",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",

          #snapshot
          "ec2:DescribeAccountAttributes",
          "rds:CreateDBClusterSnapshot",
          "rds:DeleteDBClusterSnapshot",
          "rds:DescribeDBClusterSnapshots",

          "ec2:DescribeAccountAttributes",
          "rds:CreateEventSubscription",
          "rds:DescribeEventSubscriptions",
          "rds:DeleteEventSubscription",
          "rds:ModifyEventSubscription",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",
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

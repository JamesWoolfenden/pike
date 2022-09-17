resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "rds:ListTagsForResource",
          "rds:CreateDBParameterGroup",
          "rds:DescribeDBParameterGroups",
          "rds:DescribeDBParameters",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",
          "rds:DeleteDBParameterGroup",
          "rds:ModifyDBParameterGroup",

          "rds:CreateDBClusterParameterGroup",
          "rds:ModifyDBClusterParameterGroup",
          "rds:DescribeDBClusterParameterGroups",
          "rds:DescribeDBClusterParameters",
          "rds:DeleteDBClusterParameterGroup",
          "rds:ListTagsForResource",
          "rds:AddTagsToResource",
          "rds:RemoveTagsFromResource",

          "elasticache:CreateCacheParameterGroup",
          "elasticache:AddTagsToResource",
          "elasticache:ModifyCacheParameterGroup",
          "elasticache:DescribeCacheParameterGroups",
          "elasticache:DescribeCacheParameters",
          "elasticache:ListTagsForResource",
          "elasticache:DeleteCacheParameterGroup",
          "elasticache:RemoveTagsFromResource",

          "dax:CreateParameterGroup",
          "dax:UpdateParameterGroup",
          "dax:DescribeParameterGroups",
          "dax:DescribeParameters",
          "dax:DeleteParameterGroup"
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

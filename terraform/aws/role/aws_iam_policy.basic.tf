resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_wafregional_web_acl
          "waf-regional:ListWebACLs",
          //aws_wafregional_rule
          "waf-regional:ListRules",
          //aws_wafregional_rate_based_rule
          "waf-regional:ListRateBasedRules",
          //aws_wafregional_ipset
          "waf-regional:ListIPSets",
          //aws_waf_web_acl
          "waf:ListWebACLs",
          //aws_waf_rule
          "waf:ListRules",
          //aws_waf_rate_based_rule
          "waf:ListRateBasedRules",
          //aws_waf_ipset
          "waf:ListIPSets",
          //aws_dms_replication_task
          "dms:DescribeReplicationTasks",
          //aws_dms_replication_subnet_group
          "dms:DescribeReplicationSubnetGroups",
          //aws_dms_replication_instance
          "dms:DescribeReplicationInstances",
          //aws_dms_endpoint
          "dms:DescribeEndpoints"
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

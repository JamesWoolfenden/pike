resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //

          //aws_ce_anomaly_monitor
          "ce:CreateAnomalyMonitor",
          "ce:TagResource",
          "ce:UntagResource",
          "ce:GetAnomalyMonitors",
          "ce:ListTagsForResource",
          "ce:DeleteAnomalyMonitor",
          "ce:UpdateAnomalyMonitor",

          //aws_ce_cost_allocation_tag
          "ce:ListCostAllocationTags",
          "ce:UpdateCostAllocationTagsStatus",

          //aws_ce_cost_category
          "ce:CreateCostCategoryDefinition",
          "ce:DescribeCostCategoryDefinition",
          "ce:DeleteCostCategoryDefinition",
          "ce:UpdateCostCategoryDefinition",

          //aws_ce_anomaly_subscription
          "ce:CreateAnomalySubscription",
          "ce:GetAnomalySubscriptions",
          "ce:DeleteAnomalySubscription",
          "ce:UpdateAnomalySubscription",
          "ce:TagResource",
          "ce:UntagResource",

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

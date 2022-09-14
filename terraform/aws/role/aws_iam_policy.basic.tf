resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "inspector:CreateAssessmentTemplate",
          "inspector:DescribeAssessmentTemplates",
          "inspector:ListEventSubscriptions",
          "inspector:ListTagsForResource",

          "inspector:SetTagsForResource",

          "inspector:DeleteAssessmentTemplate",

          "inspector:SubscribeToEvent",
          "inspector:UnsubscribeFromEvent",

          "inspector:CreateAssessmentTarget",
          "inspector:DescribeAssessmentTargets",
          "inspector:DeleteAssessmentTarget",

          "inspector:CreateResourceGroup",
          "inspector:DescribeResourceGroups",

          "inspector:ListRulesPackages"
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

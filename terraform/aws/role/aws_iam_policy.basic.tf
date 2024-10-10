resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "VisualEditor0",
        "Effect" : "Allow",
        "Action" : [
          //aws_quicksight_account_subscription
          "quicksight:DescribeAccountSubscription",
          "quicksight:CreateAccountSubscription",
          "quicksight:DeleteAccountSubscription",

          //aws_quicksight_folder_membership
          "quicksight:CreateFolderMembership",
          "quicksight:DeleteFolderMembership",

          //aws_quicksight_iam_policy_assignment
          "quicksight:DescribeIAMPolicyAssignment",
          "quicksight:CreateIAMPolicyAssignment",
          "quicksight:DeleteIAMPolicyAssignment",
          "quicksight:UpdateIAMPolicyAssignment",

          //aws_quicksight_ingestion
          "quicksight:DescribeIngestion",
          "quicksight:CreateIngestion",
          "quicksight:CancelIngestion",

          //aws_quicksight_template_alias
          "quicksight:DescribeTemplateAlias",
          "quicksight:CreateTemplateAlias",
          "quicksight:DeleteTemplateAlias",
          "quicksight:UpdateTemplateAlias",

          //aws_quicksight_vpc_connection
          "quicksight:DescribeVPCConnection",
          "quicksight:CreateVPCConnection",
          "quicksight:DeleteVPCConnection",
          "quicksight:UpdateVPCConnection"
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

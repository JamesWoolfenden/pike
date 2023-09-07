resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_ssm_patch_baseline
          "ssm:DescribePatchBaselines",
          //aws_ssm_parameters_by_path
          "ssm:GetParametersByPath",
          //aws_ssm_maintenance_windows
          "ssm:DescribeMaintenanceWindows",
          //aws_ssm_instances
          "ssm:DescribeInstanceInformation",
          //aws_sfn_activity
          "states:ListActivities",
          //aws_quicksight_user
          "quicksight:DescribeUser",
          //aws_quicksight_theme
          "quicksight:DescribeTheme",
          //aws_quicksight_group
          "quicksight:DescribeGroup",

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

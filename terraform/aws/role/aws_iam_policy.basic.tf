resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [

          //aws_codepipeline_webhook
          "codepipeline:DeleteWebhook",
          "codepipeline:PutWebhook",
          "codepipeline:ListWebhooks",

          //aws_codepipeline_custom_action_type
          "codepipeline:CreateCustomActionType",
          "codepipeline:DeleteCustomActionType",
          "codepipeline:ListActionTypes",
          "codepipeline:ListTagsForResource",

          //aws_codegurureviewer_repository_association
          "codeguru-reviewer:AssociateRepository",
          "codeguru-reviewer:DisassociateRepository",
          "codeguru-reviewer:DescribeRepositoryAssociation",
          //aws_codeguruprofiler_profiling_group
          "codeguru-profiler:UpdateProfilingGroup",
          "codeguru-profiler:DeleteProfilingGroup",
          "codeguru-profiler:CreateProfilingGroup",
          "codeguru-profiler:DescribeProfilingGroup",
          //aws_codecatalyst_source_repository

          //aws_codecatalyst_project
          //aws_codecatalyst_dev_environment




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

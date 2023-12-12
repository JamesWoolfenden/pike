resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_codebuild_source_credential
          "codebuild:ImportSourceCredentials",
          "codebuild:ListSourceCredentials",
          "codebuild:DeleteSourceCredentials",


          //aws_codebuild_project
          "codebuild:CreateProject",
          "iam:PassRole",
          "codebuild:BatchGetProjects",
          "codebuild:DeleteProject",

          //aws_codebuild_webhook
          "codebuild:CreateWebhook",
          "codebuild:DeleteWebhook",
          "codebuild:UpdateWebhook",

          //aws_codecommit_repository
          "codecommit:CreateRepository",
          "codecommit:GetRepository",
          "codecommit:ListTagsForResource",
          "codecommit:DeleteRepository",
          "codecommit:ListTagsForResource",

          //aws_codecommit_approval_rule_template
          "codecommit:CreateApprovalRuleTemplate",
          "codecommit:GetApprovalRuleTemplate",
          "codecommit:DeleteApprovalRuleTemplate",
          "codecommit:DisassociateApprovalRuleTemplateFromRepository",

          //aws_kms_key_policy
          "kms:PutKeyPolicy",
          "kms:DescribeKey",
          "kms:GetKeyPolicy",
          "kms:GetKeyRotationStatus",
          "kms:ListResourceTags",

          //aws_ebs_encryption_by_default
          "ec2:EnableEbsEncryptionByDefault",
          "ec2:GetEbsEncryptionByDefault",
          "ec2:DisableEbsEncryptionByDefault",

          //aws_ebs_default_kms_key
          "ec2:ModifyEbsDefaultKmsKeyId",
          "ec2:GetEbsEncryptionByDefault",
          "ec2:GetEbsDefaultKmsKeyId",
          "ec2:ResetEbsDefaultKmsKeyId",


          //aws_codecommit_trigger
          "codecommit:PutRepositoryTriggers",
          "codecommit:GetRepositoryTriggers",
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

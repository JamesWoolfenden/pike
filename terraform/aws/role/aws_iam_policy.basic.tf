resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "s3:CreateBucket",
          "s3:DeleteBucket",
          "s3:GetAccelerateConfiguration",
          "s3:GetBucketAcl",
          "s3:GetBucketCORS",
          "s3:GetBucketLogging",
          "s3:GetBucketObjectLockConfiguration",
          "s3:GetBucketPolicy",
          "s3:GetBucketRequestPayment",
          "s3:GetBucketTagging",
          "s3:GetBucketVersioning",
          "s3:GetBucketWebsite",
          "s3:GetEncryptionConfiguration",
          "s3:GetLifecycleConfiguration",
          "s3:GetObject",
          "s3:GetObjectAcl",
          "s3:GetReplicationConfiguration",
          "s3:ListBucket",
          "iam:CreateRole",
          "iam:DeleteRole",
          "iam:GetRole",
          "iam:ListAttachedRolePolicies",
          "iam:ListInstanceProfilesForRole",
          "iam:ListRolePolicies",


          //aws_auditmanager_assessment
          "auditmanager:GetAssessment",
          "auditmanager:CreateAssessment",
          "auditmanager:DeleteAssessment",
          "auditmanager:UpdateAssessment",

          //aws_auditmanager_control
          "auditmanager:GetControl",
          "auditmanager:CreateControl",
          "auditmanager:DeleteControl",
          "auditmanager:UpdateControl",
          "auditmanager:TagResource",
          "auditmanager:UntagResource",

          //aws_auditmanager_framework_share
          "auditmanager:DeleteAssessmentFrameworkShare",
          "auditmanager:StartAssessmentFrameworkShare",
          "auditmanager:UpdateAssessmentFrameworkShare",
          "auditmanager:ListAssessmentFrameworkShareRequests",

          //aws_auditmanager_organization_admin_account_registration
          "auditmanager:RegisterOrganizationAdminAccount",
          "auditmanager:DeregisterOrganizationAdminAccount",
          "auditmanager:GetOrganizationAdminAccount",

          //aws_auditmanager_account_registration
          "auditmanager:RegisterAccount",
          "auditmanager:DeregisterAccount",
          "iam:CreateServiceLinkedRole",
          "events:PutRule",
          "events:PutTargets",
          "auditmanager:GetAccountStatus",


          //aws_auditmanager_framework
          "auditmanager:DeleteAssessmentFramework",
          "auditmanager:CreateAssessmentFramework",
          "auditmanager:GetAssessmentFramework",
          "auditmanager:UpdateAssessmentFramework",

          //aws_auditmanager_assessment_delegation
          "auditmanager:GetDelegations",
          "auditmanager:BatchDeleteDelegationByAssessment"
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

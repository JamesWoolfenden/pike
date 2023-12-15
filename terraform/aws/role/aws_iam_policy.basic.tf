resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          "kms:DescribeKey",
          "kms:GetKeyPolicy",
          "kms:GetKeyRotationStatus",

          //aws_ec2_availability_zone_group
          "ec2:DescribeAvailabilityZones",

          //aws_eip_association
          "ec2:AssociateAddress",
          "ec2:DisassociateAddress",

          //aws_elastic_beanstalk_application
          "elasticbeanstalk:CreateApplication",

          //aws_iam_security_token_service_preferences
          "iam:SetSecurityTokenServicePreferences",
          "iam:GetAccountSummary",

          //aws_iam_signing_certificate
          "iam:UploadSigningCertificate",
          "iam:DeleteSigningCertificate",
          "iam:UpdateSigningCertificate",
          "iam:ListSigningCertificates",

          //aws_iam_virtual_mfa_device
          "iam:CreateVirtualMFADevice",
          "iam:ListVirtualMFADevices",
          "iam:ListMFADeviceTags",
          "iam:DeleteVirtualMFADevice",

          //aws_imagebuilder_component
          "imagebuilder:CreateComponent",
          "imagebuilder:TagResource",
          "imagebuilder:UnagResource",
          "imagebuilder:GetComponent",
          "imagebuilder:DeleteComponent",

          //aws_inspector2_delegated_admin_account
          "inspector2:EnableDelegatedAdminAccount",
          "organizations:ListDelegatedAdministrators",
          "organizations:EnableAwsServiceAccess",
          "inspector2:ListDelegatedAdminAccounts",
          "inspector2:DisableDelegatedAdminAccount",

          //aws_inspector2_enabler
          "inspector2:Enable",
          "iam:CreateServiceLinkedRole",
          "inspector2:BatchGetAccountStatus",
          "inspector2:Disable",

          //aws_inspector2_member_association
          "inspector2:AssociateMember",
          "inspector2:DisassociateMember",

          //aws_inspector2_organization_configuration
          "inspector2:UpdateOrganizationConfiguration",
          "inspector2:DescribeOrganizationConfiguration",

          //aws_internetmonitor_monitor
          "internetmonitor:TagResource",
          "internetmonitor:CreateMonitor",
          "internetmonitor:GetMonitor",
          "internetmonitor:UpdateMonitor",
          "internetmonitor:DeleteMonitor",

          //aws_kms_ciphertext
          "kms:Encrypt",

          //aws_kms_external_key
          "kms:CreateKey",
          "iam:CreateServiceLinkedRole",
          "kms:GetParametersForImport",
          "kms:GetKeyPolicy",
          "kms:ListResourceTags",
          "kms:ScheduleKeyDeletion",
          "kms:ImportKeyMaterial",

          //aws_kms_replica_key
          "kms:ReplicateKey"
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

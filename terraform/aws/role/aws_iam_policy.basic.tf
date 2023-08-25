resource "aws_iam_policy" "basic" {
  name = "basic"
  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Sid" : "0",
        "Effect" : "Allow",
        "Action" : [
          //aws_acmpca_certificate_authority
          "acm-pca:DescribeCertificateAuthority",
          "acm-pca:GetCertificateAuthorityCertificate",
          "acm-pca:GetCertificateAuthorityCsr",
          "acm-pca:ListTags",

          //aws_acmpca_certificate
          "acm-pca:IssueCertificate",
          "acm-pca:GetCertificate",
          "acm-pca:RevokeCertificate",
          //aws_acmpca_certificate_authority_certificate
          "acm-pca:ImportCertificateAuthorityCertificate",
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

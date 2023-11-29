data "aws_iam_policy_document" "example" {
  statement {
    sid    = "search_only"
    effect = "Allow"

    principals {
      type        = "*"
      identifiers = ["*"]
    }

    actions = [
      "cloudsearch:search",
      "cloudsearch:document",
    ]

    condition {
      test     = "IpAddress"
      variable = "aws:SourceIp"
      values   = ["192.0.2.0/32"]
    }
  }
}



resource "aws_cloudsearch_domain_service_access_policy" "example" {
  provider      = aws.central
  domain_name   = aws_cloudsearch_domain.pike.id
  access_policy = data.aws_iam_policy_document.example.json
}

data "aws_iam_policy_document" "example" {
  statement {
    effect    = "Allow"
    actions   = ["*"]
    resources = ["*"]
  }
}

resource "aws_organizations_policy" "example" {
  name    = "example"
  content = data.aws_iam_policy_document.example.json
}

resource "aws_glacier_vault_lock" "pike" {
  complete_lock = false
  policy        = data.aws_iam_policy_document.example.json
  vault_name    = aws_glacier_vault.pike.name
}

data "aws_iam_policy_document" "example" {
  statement {
    actions   = ["glacier:DeleteArchive"]
    effect    = "Deny"
    resources = [aws_glacier_vault.pike.arn]
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::680235478471:root"]
    }
    condition {
      test     = "NumericLessThanEquals"
      variable = "glacier:ArchiveAgeinDays"
      values   = ["365"]
    }
  }
}

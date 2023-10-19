

data "aws_secretsmanager_secret" "pike" {
  name = "pike"
}

data "aws_iam_policy_document" "example" {
  statement {
    sid    = "EnableAnotherAWSAccountToReadTheSecret"
    effect = "Allow"

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::680235478471:root"]
    }

    actions   = ["secretsmanager:GetSecretValue"]
    resources = ["*"]
  }
}

resource "aws_secretsmanager_secret_policy" "example" {
  secret_arn = data.aws_secretsmanager_secret.pike.arn
  policy     = data.aws_iam_policy_document.example.json
}

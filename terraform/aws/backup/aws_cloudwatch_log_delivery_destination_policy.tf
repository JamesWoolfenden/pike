resource "aws_cloudwatch_log_delivery_destination_policy" "example" {
  delivery_destination_name   = aws_cloudwatch_log_delivery_destination.example.name
  delivery_destination_policy = data.aws_iam_policy_document.example.json
}


data "aws_iam_policy_document" "example" {
  statement {
    sid    = "1"
    effect = "Allow"
    actions = [
      "logs:CreateDelivery",
    ]

    resources = [
      "*",
    ]
    principals {
      identifiers = ["AWS"]
      type        = "arn:aws:iam::680235478471:root"
    }
  }
}

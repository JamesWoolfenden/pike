resource "aws_cloudwatch_log_destination_policy" "pike" {
  access_policy    = data.aws_iam_policy_document.test_destination_policy.json
  destination_name = aws_cloudwatch_log_destination.pike.name
}

data "aws_iam_policy_document" "test_destination_policy" {
  statement {
    effect = "Allow"

    principals {
      type = "AWS"

      identifiers = [
        "680235478471",
      ]
    }

    actions = [
      "logs:PutSubscriptionFilter",
    ]

    resources = [
      aws_cloudwatch_log_destination.pike.arn,
    ]
  }
}

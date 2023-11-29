data "aws_iam_policy_document" "test" {
  statement {
    sid    = "AllowAccountToPutEvents"
    effect = "Allow"
    actions = [
      "events:PutEvents",
    ]
    resources = [
      "arn:aws:events:us-east-1:680235478471:event-bus/pike"
    ]

    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::680235478471:root"]
    }
  }
  statement {
    actions = ["events:PutEvents"]

    condition {
      test     = "StringEquals"
      values   = ["o-5qepxnywkl"]
      variable = "aws:PrincipalOrgID"
    }

    resources = ["arn:aws:events:us-east-1:680235478471:event-bus/pike"]
    sid       = "AllowAllAccountsFromOrganizationToPutEvents"

  }

  statement {
    actions = [
      "events:PutRule",
      "events:PutTargets",
      "events:DeleteRule",
      "events:RemoveTargets",
      "events:DisableRule",
      "events:EnableRule",
      "events:TagResource",
      "events:UntagResource",
      "events:DescribeRule",
      "events:ListTargetsByRule",
      "events:ListTagsForResource",
    ]

    condition {
      test     = "StringEqualsIfExists"
      values   = ["680235478471"]
      variable = "events:creatorAccount"
    }
    sid       = "AllowAccountToManageRulesTheyCreated"
    resources = ["arn:aws:events:us-east-1:680235478471:rule/pike"]
    principals {
      identifiers = ["arn:aws:iam::680235478471:root"]
      type        = "AWS"
    }
  }
}

resource "aws_cloudwatch_event_bus_policy" "test" {
  provider       = aws.central
  policy         = data.aws_iam_policy_document.test.json
  event_bus_name = aws_cloudwatch_event_bus.pike.name
}

resource "aws_cloudwatch_event_bus" "pike" {
  provider = aws.central
  name     = "pike"
  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

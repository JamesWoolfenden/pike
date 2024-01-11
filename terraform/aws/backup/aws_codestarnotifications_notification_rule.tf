resource "aws_codecommit_repository" "code" {
  repository_name = "example-code-repo"
}

resource "aws_sns_topic" "notif" {
  name = "notification"
}

data "aws_iam_policy_document" "notif_access" {
  statement {
    actions = ["sns:Publish"]

    principals {
      type        = "Service"
      identifiers = ["codestar-notifications.amazonaws.com"]
    }

    resources = [aws_sns_topic.notif.arn]
  }
}

resource "aws_sns_topic_policy" "default" {
  arn    = aws_sns_topic.notif.arn
  policy = data.aws_iam_policy_document.notif_access.json
}

resource "aws_codestarnotifications_notification_rule" "commits" {
  detail_type    = "BASIC"
  event_type_ids = ["codecommit-repository-comments-on-commits"]

  name     = "example-code-repo-commits"
  resource = aws_codecommit_repository.code.arn

  target {
    address = aws_sns_topic.notif.arn
  }

  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

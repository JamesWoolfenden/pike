resource "aws_dms_event_subscription" "pike" {
  enabled          = true
  event_categories = ["creation", "failure"]
  name             = "my-favorite-event-subscription"
  sns_topic_arn    = aws_sns_topic.pike.arn
  source_ids       = [aws_dms_replication_task.pike.replication_task_id]
  source_type      = "replication-task"

  tags = {
    Name = "example"
  }
}

resource "aws_sns_topic" "pike" {}

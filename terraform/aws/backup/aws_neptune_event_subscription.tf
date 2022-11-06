resource "aws_neptune_event_subscription" "pike" {
  name          = "neptune-event-sub"
  sns_topic_arn = "arn:aws:sns:eu-west-2:680235478471:pike"

  source_type = "db-instance"
  source_ids  = [aws_neptune_cluster_instance.pike.id]

  event_categories = [
    "maintenance",
    "availability",
    "creation",
    "backup",
    "restoration",
    "recovery",
    "deletion",
    "failover",
    "failure",
    "notification",
    "configuration change",
    "read replica",
  ]

  tags = {
    pike = "permissions"
  }
}

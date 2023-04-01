resource "aws_db_event_subscription" "default" {
  name      = "rds-event-sub"
  sns_topic = "topic"

  source_type = "db-instance"
  source_ids  = ["id"]

  event_categories = [
    "availability",
    "deletion",
    "failover",
    "failure",
    "low storage",
    "maintenance",
    "notification",
    "read replica",
    "recovery",
    "restoration",
  ]
}

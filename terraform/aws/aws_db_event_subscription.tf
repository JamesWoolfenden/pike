#resource "aws_db_event_subscription" "pike" {
#    name      = "rds-event-sub"
#    sns_topic = aws_sns_topic.default.arn
#
#    source_type = "db-instance"
#    source_ids  = [aws_db_instance.pike.id]
#
#    event_categories = [
#      "availability",
#      "deletion",
#      "failover",
#      "failure",
#      "low storage",
#      "maintenance",
#      "notification",
#      "read replica",
#      "recovery",
#      "restoration",
#    ]
#  }

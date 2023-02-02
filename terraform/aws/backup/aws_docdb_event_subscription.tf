resource "aws_docdb_event_subscription" "pike" {
  name             = "example"
  enabled          = true
  event_categories = ["creation", "failure"]
  source_type      = "db-cluster"
  source_ids       = [aws_docdb_cluster.pike.id]
  sns_topic_arn    = aws_sns_topic.example.arn
}

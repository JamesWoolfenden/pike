resource "aws_redshift_event_subscription" "pike" {
  name          = "pike"
  sns_topic_arn = "arn:aws:sns:eu-west-2:680235478471:pike"

  source_type = "cluster"
  source_ids  = ["redshift-cluster-1"]

  severity = "INFO"

  event_categories = [
    // "configuration",
    "management",
    "monitoring",
    "security",
  ]

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}

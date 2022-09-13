resource "aws_redshift_scheduled_action" "pike" {
  name     = "tf-redshift-scheduled-action"
  schedule = "cron(00 22 * * ? *)"
  iam_role = "arn:aws:iam::680235478471:role/redshiftScheduler"

  target_action {
    pause_cluster {
      cluster_identifier = "redshift-cluster-1"
    }
  }

}

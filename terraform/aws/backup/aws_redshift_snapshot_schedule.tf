resource "aws_redshift_snapshot_schedule" "pike" {
  identifier = "pike-snapshot-schedule"
  definitions = [
    "rate(14 hours)",
  ]
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}

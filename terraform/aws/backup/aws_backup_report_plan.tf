resource "aws_backup_report_plan" "pike" {
  name        = "pike"
  description = "example description"

  report_delivery_channel {
    formats = [
      "CSV",
      "JSON"
    ]
    s3_bucket_name = "testbucketineu-west2"
  }

  report_setting {
    report_template = "RESTORE_JOB_REPORT"
  }

  tags = {
    pike = "permissions"
  }
}

resource "aws_s3control_storage_lens_configuration" "pike" {
  config_id = "example-1"

  storage_lens_configuration {
    enabled = true

    account_level {
      activity_metrics {
        enabled = true
      }

      bucket_level {
        activity_metrics {
          enabled = true
        }
      }
    }
  }
}

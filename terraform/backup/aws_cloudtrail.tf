resource "aws_cloudtrail" "pike" {
  name                       = "pike"
  s3_bucket_name             = "trails-test-jgw"
  cloud_watch_logs_group_arn = "arn:aws:logs:eu-west-2:680235478471:log-group:pike:*"
  cloud_watch_logs_role_arn  = "arn:aws:iam::680235478471:role/cloudwatch"

  event_selector {
    exclude_management_event_sources = ["kms.amazonaws.com", "rdsdata.amazonaws.com"]
    include_management_events        = true
  }
  event_selector {
    read_write_type           = "All"
    include_management_events = true

    data_resource {
      type = "AWS::S3::Object"

      # Make sure to append a trailing '/' to your ARN if you want
      # to monitor all objects in a bucket.
      values = ["arn:aws:s3:::testbucketineu-west2/"]
    }
  }
  #  advanced_event_selector {
  #    name = "Log Delete* events for one S3 bucket"
  #
  #    field_selector {
  #      field  = "eventCategory"
  #      equals = ["Data"]
  #    }
  #
  #    field_selector {
  #      field       = "eventName"
  #      starts_with = ["Delete"]
  #    }
  #
  #    field_selector {
  #      field = "resources.ARN"
  #
  #      #The trailing slash is intentional; do not exclude it.
  #      equals = [
  #        "arn:aws:s3:::testbucketineu-west2/"
  #      ]
  #    }
  #
  #    field_selector {
  #      field  = "readOnly"
  #      equals = ["false"]
  #    }
  #
  #    field_selector {
  #      field  = "resources.type"
  #      equals = ["AWS::S3::Object"]
  #    }
  #  }
  kms_key_id                 = "arn:aws:kms:eu-west-2:680235478471:key/554dbedc-3cf9-4aec-b621-9c4387bed449"
  enable_log_file_validation = true

  #  insight_selector {
  #    insight_type = "ApiErrorRateInsight"
  #  }

  sns_topic_name = "arn:aws:sns:eu-west-2:680235478471:aws-cloudtrail-logs-680235478471-6fab4a93"

  tags = {
    pike    = "permissions"
    another = "one"
  }
}

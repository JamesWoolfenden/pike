data "aws_cloudwatch_log_data_protection_policy_document" "pike" {
  name = "pike"

  statement {
    sid = "Audit"

    data_identifiers = [
      "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
      "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
    ]

    operation {
      audit {
        findings_destination {
          cloudwatch_logs {
            log_group = "pike"
          }

          s3 {
            bucket = "pike"
          }
        }
      }
    }
  }

  statement {
    sid = "Deidentify"

    data_identifiers = [
      "arn:aws:dataprotection::aws:data-identifier/EmailAddress",
      "arn:aws:dataprotection::aws:data-identifier/DriversLicense-US",
    ]

    operation {
      deidentify {
        mask_config {}
      }
    }
  }
}

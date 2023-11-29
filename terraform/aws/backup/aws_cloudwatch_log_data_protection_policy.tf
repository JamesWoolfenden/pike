resource "aws_cloudwatch_log_group" "example" {
  provider = aws.central
  name     = "example-jgw-test"
}

resource "aws_s3_bucket" "example" {
  provider = aws.central
  bucket   = "example-jgw-test2023"
}

resource "aws_cloudwatch_log_data_protection_policy" "pike" {
  provider       = aws.central
  log_group_name = aws_cloudwatch_log_group.example.name

  policy_document = jsonencode({
    Name    = "Example2"
    Version = "2021-06-01"

    Statement = [
      {
        Sid            = "Audit"
        DataIdentifier = ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"]
        Operation = {
          Audit = {
            FindingsDestination = {
              S3 = {
                Bucket = aws_s3_bucket.example.bucket
              }
            }
          }
        }
      },
      {
        Sid            = "Redact"
        DataIdentifier = ["arn:aws:dataprotection::aws:data-identifier/EmailAddress"]
        Operation = {
          Deidentify = {
            MaskConfig = {}
          }
        }
      }
    ]
  })
}

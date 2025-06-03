resource "aws_kms_key" "example" {
  description             = "KMS key for WorkSpaces Web Browser Settings"
  deletion_window_in_days = 7
}

resource "aws_kms_key_policy" "example" {
  key_id = aws_kms_key.example.key_id

  policy = jsonencode({
    "Version" : "2012-10-17",
    "Statement" : [
      {
        "Effect" : "Allow",
        "Action" : [
          "kms:*"
        ],
        Principal = {
          AWS = "*"
        }
        "Resource" : "*"
      },
      {
        "Effect" : "Allow",
        "Action" : [
          "workspaces:*"
        ],
        Principal = {
          AWS = "*"
        }
        "Resource" : "*"
      }
    ]
  })
}

resource "aws_workspacesweb_browser_settings" "pike" {
  browser_policy = jsonencode({
    chromePolicies = {
      DefaultDownloadDirectory = {
        value = "/home/as2-streaming-user/MyFiles/TemporaryFiles1"
      }
    }
  })
  customer_managed_key = aws_kms_key.example.arn
  additional_encryption_context = {
    Environment = "Development"
  }
  tags = {
    Name = "example-browser-settings"
  }
}

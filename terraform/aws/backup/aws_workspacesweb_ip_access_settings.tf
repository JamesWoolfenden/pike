resource "aws_kms_key" "example" {
  description             = "KMS key for WorkSpaces Web IP Access Settings"
  deletion_window_in_days = 7

  policy = jsonencode({
    Id = "example"
    Statement = [
      {
        Action = "kms:*"
        Effect = "Allow"
        Principal = {
          AWS = "*"
        }

        Resource = "*"
        Sid      = "Enable IAM User Permissions"
      },
    ]
    Version = "2012-10-17"
  })
}

resource "aws_workspacesweb_ip_access_settings" "example" {
  display_name         = "example"
  description          = "Example IP access settings"
  customer_managed_key = aws_kms_key.example.arn
  additional_encryption_context = {
    Environment = "Production"
  }
  ip_rule {
    ip_range    = "10.0.0.0/16"
    description = "Main office"
  }
  ip_rule {
    ip_range    = "192.168.0.0/24"
    description = "Branch office"
  }
  tags = {
    Name = "example-ip-access-settings"
  }
}

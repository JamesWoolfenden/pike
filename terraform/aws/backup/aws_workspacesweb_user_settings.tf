resource "aws_workspacesweb_user_settings" "example" {
  copy_allowed                       = "Enabled"
  download_allowed                   = "Enabled"
  paste_allowed                      = "Enabled"
  print_allowed                      = "Enabled"
  upload_allowed                     = "Enabled"
  deep_link_allowed                  = "Enabled"
  disconnect_timeout_in_minutes      = 30
  idle_disconnect_timeout_in_minutes = 15
  customer_managed_key               = aws_kms_key.example.arn

  additional_encryption_context = {
    Environment = "Production"
  }

  toolbar_configuration {
    toolbar_type           = "Docked"
    visual_mode            = "Dark"
    hidden_toolbar_items   = ["Webcam", "Microphone"]
    max_display_resolution = "size1920X1080"
  }

  cookie_synchronization_configuration {
    allowlist {
      domain = "example.com"
      path   = "/path"
    }
    blocklist {
      domain = "blocked.com"
    }
  }

  tags = {
    Name = "example-user-settings"
  }
}

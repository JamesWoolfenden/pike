
resource "aws_workspaces_workspace" "pike" {
  user_name                      = "jim1972"
  bundle_id                      = data.aws_workspaces_bundle.value_windows_10.id
  directory_id                   = aws_workspaces_directory.pike.id
  root_volume_encryption_enabled = false
  user_volume_encryption_enabled = true
  volume_encryption_key          = "arn:aws:kms:eu-west-1:680235478471:key/cb37def8-1aab-4b9f-95f8-461da92eab84"

  workspace_properties {
    compute_type_name = "STANDARD"
    //user_volume_size_gib                      = 0
    root_volume_size_gib                      = 80
    running_mode                              = "AUTO_STOP"
    running_mode_auto_stop_timeout_in_minutes = 60
  }
}

data "aws_workspaces_bundle" "value_windows_10" {
  bundle_id = "wsb-clj85qzj1"
}

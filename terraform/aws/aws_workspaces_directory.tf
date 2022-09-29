resource "aws_workspaces_directory" "pike" {
  tags = {
    pike = "permissions"
  }
  self_service_permissions {
    change_compute_type  = false
    increase_volume_size = false
    rebuild_workspace    = false
    restart_workspace    = true
    switch_running_mode  = false
  }

  workspace_access_properties {
    device_type_android    = "ALLOW"
    device_type_chromeos   = "ALLOW"
    device_type_ios        = "ALLOW"
    device_type_linux      = "DENY"
    device_type_osx        = "ALLOW"
    device_type_web        = "DENY"
    device_type_windows    = "ALLOW"
    device_type_zeroclient = "ALLOW"
  }
  directory_id = aws_directory_service_directory.pike2.id
}

resource "aws_directory_service_directory" "pike2" {
  name     = "corp.pike.com"
  password = "PikeI$Permission2"
  vpc_settings {
    vpc_id     = "vpc-032c4e15f24d3628e"
    subnet_ids = ["subnet-04338b6369d8288a5", "subnet-0d99e0f4e4f1ff84f"]
  }



  tags = {
    pike   = "permissions"
    delete = "me"
  }
}

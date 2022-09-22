#resource "aws_workspaces_directory" "pike" {}
#
#resource "aws_directory_service_directory" "pike2" {
#  name     = "corp.pike.com"
#  password = "PikeI$Permission2"
#  vpc_settings {
#    vpc_id     = "vpc-032c4e15f24d3628e"
#    subnet_ids = ["subnet-04338b6369d8288a5", "subnet-0d99e0f4e4f1ff84f"]
#  }
#
#  tags = {
#    pike="permissions"
#    delete="me"
#  }
#}
#
#data "aws_workspaces_bundle" "value_windows_10" {
#  bundle_id = "wsb-bh8rsxt14" # Value with Windows 10 (English)
#}

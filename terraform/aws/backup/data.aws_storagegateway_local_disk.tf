data "aws_storagegateway_local_disk" "pike" {
  gateway_arn = aws_storagegateway_gateway.example.arn
}

#resource "aws_storagegateway_gateway" "example" {
#  gateway_ip_address = "1.2.3.4"
#  gateway_name       = "example"
#  gateway_timezone   = "GMT"
#  gateway_type       = "FILE_FSX_SMB"
#  smb_active_directory_settings {
#    domain_name = "corp.example.com"
#    password    = "avoid-plaintext-passwords"
#    username    = "Admin"
#  }
#}

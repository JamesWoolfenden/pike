resource "aws_fsx_ontap_storage_virtual_machine" "pike" {
  file_system_id = aws_fsx_ontap_file_system.pike.id
  name           = "pike"
  tags = {
    pike = "permissions"
  }
}

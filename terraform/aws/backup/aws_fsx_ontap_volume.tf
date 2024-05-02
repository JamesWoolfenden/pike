resource "aws_fsx_ontap_volume" "pike" {
  name                       = "test"
  junction_path              = "/test"
  size_in_megabytes          = 1024
  storage_efficiency_enabled = true
  storage_virtual_machine_id = aws_fsx_ontap_storage_virtual_machine.pike.id
}

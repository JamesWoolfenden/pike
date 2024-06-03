resource "aws_storagegateway_stored_iscsi_volume" "pike" {
  gateway_arn            = aws_storagegateway_cache.pike.gateway_arn
  network_interface_id   = aws_instance.test.private_ip
  target_name            = "example"
  preserve_existing_data = false
  disk_id                = data.aws_storagegateway_local_disk.test.id
  tags = {
    pike = "permissions"
  }
}

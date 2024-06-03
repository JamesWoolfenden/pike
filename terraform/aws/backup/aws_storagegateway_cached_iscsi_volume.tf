resource "aws_storagegateway_cached_iscsi_volume" "pike" {
  gateway_arn          = aws_storagegateway_cache.pike.gateway_arn
  network_interface_id = aws_instance.test.private_ip
  target_name          = "example"
  volume_size_in_bytes = 5368709120 # 5 GB
}

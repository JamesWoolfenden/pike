resource "aws_storagegateway_working_storage" "pike" {
  disk_id     = data.aws_storagegateway_local_disk.test.id
  gateway_arn = aws_storagegateway_gateway.pike.arn
}

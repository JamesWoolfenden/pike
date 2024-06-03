resource "aws_storagegateway_upload_buffer" "pike" {
  disk_path   = data.aws_storagegateway_local_disk.test.disk_path
  gateway_arn = aws_storagegateway_gateway.pike.arn
}

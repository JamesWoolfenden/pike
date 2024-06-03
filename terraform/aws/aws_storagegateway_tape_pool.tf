resource "aws_storagegateway_tape_pool" "pike" {
  pool_name     = "example"
  storage_class = "GLACIER"
  tags = {
    pike = "permissions"
  }
}

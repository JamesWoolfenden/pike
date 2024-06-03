resource "aws_storagegateway_file_system_association" "pike" {
  gateway_arn           = aws_storagegateway_gateway.pike.arn
  location_arn          = aws_fsx_windows_file_system.pike.arn
  username              = "Admin"
  password              = "avoid-plaintext-passwords"
  audit_destination_arn = aws_s3_bucket.example.arn
  tags = {
    pike = "permission"
  }
}

resource "aws_fsx_windows_file_system" "pike" {
  subnet_ids          = [data.aws_subnets.example.ids[0]]
  throughput_capacity = 8
}

data "aws_subnets" "example" {

}

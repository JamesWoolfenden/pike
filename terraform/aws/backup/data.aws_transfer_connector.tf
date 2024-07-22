data "aws_transfer_connector" "pike" {
  id = "c-12345678901234567"
}

output "aws_transfer_connector" {
  value = data.aws_transfer_connector.pike
}

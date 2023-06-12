data "aws_kms_key" "example" {
  key_id = "34cdce9a-2322-427c-91bb-b572f435c032"
}
output "key" {
  value = data.aws_kms_key.example
}

data "aws_kms_public_key" "pike" {
  key_id = "03a3077b-1b63-4f42-98a1-20ea7a2fabba"
}

output "aws_kms_public_key" {
  value = data.aws_kms_public_key.pike
}

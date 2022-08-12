data "aws_kms_ciphertext" "example" {
  key_id    = "34cdce9a-2322-427c-91bb-b572f435c032"
  plaintext = "Jimbo"
}

output "aws_kms_ciphertext" {
  sensitive = true
  value     = data.aws_kms_ciphertext.example
}

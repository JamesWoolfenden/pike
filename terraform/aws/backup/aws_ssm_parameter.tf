resource "aws_ssm_parameter" "example" {
  name        = "foo"
  type        = "SecureString"
  value       = "bar"
  description = "language timothy!"
  key_id      = "alias/test2"
  tags = {
    environment = "production"
  }
}

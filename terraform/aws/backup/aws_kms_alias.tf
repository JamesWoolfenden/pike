resource "aws_kms_alias" "example" {
  name          = "alias/my-key-pike"
  target_key_id = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
}

data "aws_ebs_default_kms_key" "example" {}

output "kms_key" {
  value = data.aws_ebs_default_kms_key.example
}

data "aws_ebs_encryption_by_default" "pike" {}


output "ebsencryption" {
  value = data.aws_ebs_encryption_by_default.pike
}

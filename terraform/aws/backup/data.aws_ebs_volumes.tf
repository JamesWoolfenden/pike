data "aws_ebs_volumes" "pike" {}

output "volumes" {
  value = data.aws_ebs_volumes.pike
}

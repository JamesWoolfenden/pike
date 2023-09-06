data "aws_launch_template" "pike" {}

output "template" {
  value = data.aws_launch_template.pike
}

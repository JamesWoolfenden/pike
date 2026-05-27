data "aws_alb_target_group" "pike" {
}

output "aws_alb_target_group" {
  value = data.aws_alb_target_group.pike
}

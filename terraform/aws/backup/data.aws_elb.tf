data "aws_elb" "pike" {
  name = "pike"
}

output "aws_elb" {
  value = data.aws_elb.pike
}

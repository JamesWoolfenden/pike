data "aws_service" "pike" {
  service_id = "ec2"
}

output "service" {
  value = data.aws_service.pike
}

data "aws_mq_broker_instance_type_offerings" "pike" {}

output "offer" {
  value = data.aws_mq_broker_instance_type_offerings.pike
}
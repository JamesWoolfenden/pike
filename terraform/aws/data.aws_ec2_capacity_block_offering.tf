data "aws_ec2_capacity_block_offering" "pike" {
  capacity_duration_hours = 24
  end_date_range          = "2024-06-30T15:04:05Z"
  instance_count          = 1
  instance_type           = "p4d.24xlarge"
  start_date_range        = "2024-06-28T15:04:05Z"
}

output "aws_ec2_capacity_block_offering" {
  value = data.aws_ec2_capacity_block_offering.pike
}

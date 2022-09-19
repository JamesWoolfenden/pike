resource "aws_ec2_tag" "pike" {
  key         = "pike"
  resource_id = "sg-002ed1a53dc5fe0ad"
  value       = "permissions"
}

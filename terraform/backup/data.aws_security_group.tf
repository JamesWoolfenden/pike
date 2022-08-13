
data "aws_security_group" "example" {
  name = "launch-wizard-1"
  tags = {
    key = "value"
  }
}

output "group" {
  value = data.aws_security_group.example
}

data "aws_security_groups" "example" {
}
output "groups" {
  value = data.aws_security_groups.example
}

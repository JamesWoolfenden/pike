data "aws_ami_ids" "pike" {
  owners = ["680235478471"]
}

output "amis" {
  value = data.aws_ami_ids.pike
}

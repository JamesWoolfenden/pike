resource "aws_ami_launch_permission" "pike" {
  image_id   = aws_ami.pike.id
  account_id = "680235478471"
}

output "perms" {
  value = aws_ami_launch_permission.pike
}

resource "aws_efs_mount_target" "pike" {
  file_system_id = "fs-0898bc857b16b617a"
  subnet_id      = "subnet-03fdfb13a135366a7"
  ip_address     = "10.0.0.8"
  security_groups = [
  "sg-05b27cb61c9c46bd2", "sg-06db45d8099f7f549", "sg-0d2f425bb20ccbd04"]
}

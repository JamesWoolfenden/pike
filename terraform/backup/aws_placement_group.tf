resource "aws_placement_group" "web" {
  name     = "hunky"
  strategy = "cluster"
  tags = {
    pike = "permissions"
  }
}

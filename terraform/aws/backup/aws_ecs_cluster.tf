resource "aws_ecs_cluster" "pike" {
  name = "pike"
  tags = {
    pike = "permissions"
  }
}

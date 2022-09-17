resource "aws_neptune_subnet_group" "pike" {
  name       = "main"
  subnet_ids = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7"]

  tags = {
    Pike   = "Permissions"
    delete = "me"
  }
}

output "neptune_subnet_group" {
  value = aws_neptune_subnet_group.pike
}

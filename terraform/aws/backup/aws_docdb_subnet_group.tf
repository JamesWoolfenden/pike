resource "aws_docdb_subnet_group" "pike" {
  name       = "pike"
  subnet_ids = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7"]

  tags = {
    Pike   = "Permissions"
    delete = "me"
  }
}

output "docdb_subnet_group" {
  value = aws_docdb_subnet_group.pike
}

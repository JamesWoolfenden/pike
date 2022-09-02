resource "aws_db_subnet_group" "default" {
  name       = "main"
  subnet_ids = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7"]

  #  tags = {
  #    Pike = "Permissions"
  #  }
}

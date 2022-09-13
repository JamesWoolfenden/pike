resource "aws_redshift_subnet_group" "pike" {
  name       = "pike"
  subnet_ids = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7"]

  tags = {
    pike   = "permissions"
    delete = "me"
  }
}

resource "aws_memorydb_subnet_group" "pike" {
  name       = "main"
  subnet_ids = ["subnet-08d97e381dbc80d40", "subnet-03fdfb13a135366a7"] //,"subnet-05a6a6de2f4989d22"]

  tags = {
    Pike = "Permissions"
    //   delete = "me"
  }
}

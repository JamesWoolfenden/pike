resource "aws_dsql_cluster" "pike" {

  deletion_protection_enabled = false
  tags = {
    Name = "TestCluster"
  }
}

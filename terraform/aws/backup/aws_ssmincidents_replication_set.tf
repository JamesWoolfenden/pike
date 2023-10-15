resource "aws_ssmincidents_replication_set" "pike" {
  region {
    name = "eu-west-2"
  }

  tags = {
    pike = "permissions"
    more = "permissions"
  }
}

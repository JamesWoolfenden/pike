resource "aws_memorydb_snapshot" "pike" {
  cluster_name = aws_memorydb_cluster.pike.name
  name_prefix  = "pikey"
  kms_key_arn  = "arn:aws:kms:eu-west-2:680235478471:key/34cdce9a-2322-427c-91bb-b572f435c032"
  tags = {
    pike   = "permissions"
    delete = "me"
  }
}

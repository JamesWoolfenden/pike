resource "aws_docdb_global_cluster" "pike" {
  global_cluster_identifier = "pike"
  database_name             = "pike"
  deletion_protection       = false
  engine                    = "docdb"
  engine_version            = "4.0.0"
  storage_encrypted         = false
}

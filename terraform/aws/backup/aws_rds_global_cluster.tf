resource "aws_rds_global_cluster" "pike" {
  global_cluster_identifier = "global-test"
  engine                    = "aurora"
  engine_version            = "5.6.mysql_aurora.1.22.2"
  database_name             = "pike"
  deletion_protection       = false
  force_destroy             = true
  //source_db_cluster_identifier = ""
  storage_encrypted = true
}

output "globalcluster" {
  value = aws_rds_global_cluster.pike
}

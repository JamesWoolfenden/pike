resource "aws_glue_partition" "pike" {
  database_name    = "some-database"
  table_name       = "some-table"
  partition_values = ["some-value"]
}

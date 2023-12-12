resource "aws_timestreamwrite_table" "pike" {
  provider      = aws.central
  database_name = aws_timestreamwrite_database.pike.database_name
  table_name    = "example"

  retention_properties {
    magnetic_store_retention_period_in_days = 30
    memory_store_retention_period_in_hours  = 9
  }

  tags = {
    pike = "permissions"
    #    Name = "example-timestream-table"
  }
}

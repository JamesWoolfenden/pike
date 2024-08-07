resource "aws_cleanrooms_configured_table" "pike" {
  name            = "terraform-example-table"
  description     = "I made this table with terraform!"
  analysis_method = "DIRECT_QUERY"
  allowed_columns = [
    "column1",
    "column2",
    "column3",
  ]

  table_reference {
    database_name = "example_database"
    table_name    = "example_table"
  }

  tags = {
    Project = "Terraform"
  }
}

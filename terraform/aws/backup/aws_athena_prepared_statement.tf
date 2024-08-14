resource "aws_athena_prepared_statement" "pike" {
  name            = "tf_test"
  query_statement = "SELECT * FROM ${aws_athena_database.pike.name} WHERE x = ?"
  workgroup       = aws_athena_workgroup.pike.name
}

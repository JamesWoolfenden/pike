resource "aws_athena_named_query" "pike" {
  name      = "bar"
  workgroup = aws_athena_workgroup.pike.id
  database  = "pike2"
  query     = "SELECT * FROM pike2 limit 10;"
}

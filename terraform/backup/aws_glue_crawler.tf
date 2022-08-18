resource "aws_glue_crawler" "example" {
  database_name = aws_glue_catalog_database.example.name
  name          = "example"
  role          = aws_iam_role.example.arn

  dynamodb_target {
    path = "table-name"
  }

  tags = {
    pike = "permissions"
  }
}

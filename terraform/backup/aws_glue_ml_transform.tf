resource "aws_glue_ml_transform" "test" {
  name     = "example"
  role_arn = aws_iam_role.test.arn

  input_record_tables {
    database_name = aws_glue_catalog_table.test.database_name
    table_name    = aws_glue_catalog_table.test.name
  }

  parameters {
    transform_type = "FIND_MATCHES"

    find_matches_parameters {
      primary_key_column_name = "my_column_1"
    }
  }

  depends_on = [aws_iam_role_policy_attachment.test]
}

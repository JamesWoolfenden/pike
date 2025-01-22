resource "aws_s3tables_table_bucket_policy" "pike" {
  resource_policy  = data.aws_iam_policy_document.example.json
  table_bucket_arn = aws_s3tables_table_bucket.example.arn
}

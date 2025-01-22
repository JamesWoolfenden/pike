resource "aws_s3tables_namespace" "pike" {
  namespace        = "pike"
  table_bucket_arn = aws_s3tables_table_bucket.example.arn
}

resource "aws_s3tables_table_bucket" "example" {
  name = "example-bucket"
}

resource "aws_s3tables_table" "pike" {
  name             = "piketable"
  namespace        = aws_s3tables_namespace.pike.namespace
  table_bucket_arn = aws_s3tables_namespace.pike.table_bucket_arn
  format           = "ICEBERG"
}


resource "aws_s3tables_table_bucket" "pike" {
  name = "pike"
}

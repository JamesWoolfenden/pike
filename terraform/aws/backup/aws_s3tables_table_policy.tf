resource "aws_s3tables_table_policy" "pike" {
  resource_policy  = data.aws_iam_policy_document.example.json
  name             = aws_s3tables_table.pike.name
  namespace        = aws_s3tables_table.pike.namespace
  table_bucket_arn = aws_s3tables_table.pike.table_bucket_arn
}

data "aws_iam_policy_document" "example" {

  statement {
    effect = "Allow"
    principals {
      type        = "AWS"
      identifiers = ["arn:aws:iam::680235478471:role/datasteward"]
    }
    actions   = ["s3tables:PutTableBucketMaintenanceConfiguration"]
    resources = ["arn:aws:s3tables:eu-west-1:680235478471:bucket/*"]
  }

}

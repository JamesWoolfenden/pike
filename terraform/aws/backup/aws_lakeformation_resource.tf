data "aws_s3_bucket" "example" {
  bucket = "pike-680235478471"
}

resource "aws_lakeformation_resource" "example" {
  arn = data.aws_s3_bucket.example.arn
}

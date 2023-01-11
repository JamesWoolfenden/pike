resource "aws_s3_bucket" "pike" {
  bucket = "thisoneshouldbeunique-${random_string.pike.id}"
}


resource "aws_s3_bucket_versioning" "pike" {
  bucket = aws_s3_bucket.pike.id

  versioning_configuration {
    status = "Enabled"
  }
}


resource "random_string" "pike" {
  special = false
  length  = 6
  lower   = true
  upper   = false
}

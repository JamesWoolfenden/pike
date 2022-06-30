resource "aws_s3_bucket" "pail" {
  bucket = "pail"
}

resource "aws_s3_pail" "bucket" {
  bucket = "pail"
}

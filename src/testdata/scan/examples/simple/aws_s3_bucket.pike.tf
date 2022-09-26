resource "aws_s3_bucket" "pike" {
  bucket = "thisoneshouldbeunique-${random_string.pike.id}"
}

resource "random_string" "pike" {
  special = false
  length  = 6
  lower   = true
  upper   = false
}

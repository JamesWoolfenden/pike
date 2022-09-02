resource "aws_s3_bucket_object" "example" {
  key        = "Makefile"
  bucket     = "testbucketforlbjgw"
  source     = "Makefile"
  kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/8479ba6c-5008-43c5-bc2d-7151657d49af"
  tags = {
    pike = "permissions"
  }
}

resource "aws_s3_object" "example" {
  key        = "Bakefile"
  bucket     = "testbucketforlbjgw"
  source     = "Makefile"
  kms_key_id = "arn:aws:kms:eu-west-2:680235478471:key/8479ba6c-5008-43c5-bc2d-7151657d49af"
  tags = {
    pike = "permissions"
  }
}

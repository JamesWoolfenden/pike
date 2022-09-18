resource "aws_s3_bucket" "pail" {
  bucket              = "pail-${data.aws_caller_identity.current.account_id}"
  force_destroy       = true
  object_lock_enabled = true
  # tags={
  #   createdby="Terraformed"
  # }
}

resource "aws_s3_bucket" "tail" {
  bucket              = "tail-${data.aws_caller_identity.current.account_id}"
  force_destroy       = true
  object_lock_enabled = true
  tags = {
    createdby = "Terraformed"
  }
}

# resource "aws_s3_pail" "bucket" {
#   bucket = "pail"
# }

#data "aws_caller_identity" "current" {}

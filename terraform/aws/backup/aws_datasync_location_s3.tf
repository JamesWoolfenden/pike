resource "aws_datasync_location_s3" "destination_pike" {
  s3_bucket_arn = aws_s3_bucket.datasync.arn
  subdirectory  = "/example/prefix"

  s3_config {
    bucket_access_role_arn = aws_iam_role.datasync.arn
  }
}

resource "aws_iam_role" "datasync" {
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : { "AWS" : "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root" },
          "Action" : "sts:AssumeRole",
        }
      ]
    }
  )
}

data "aws_caller_identity" "current" {}


resource "aws_s3_bucket" "datasync" {}

resource "aws_storagegateway_smb_file_share" "pike" {
  authentication = "ActiveDirectory"
  gateway_arn    = aws_storagegateway_gateway.pike.arn
  location_arn   = aws_s3_bucket.example.arn
  role_arn       = aws_iam_role.example.arn
  tags = {
    pike = "permissions"
  }
}

resource "aws_s3_bucket" "example" {
  bucket = "totalguffjgw"
}

resource "aws_iam_role" "example" {
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

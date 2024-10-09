resource "aws_s3_bucket" "example" {
  bucket = "examples3controlpoint"
}

resource "aws_s3_access_point" "example" {
  bucket = aws_s3_bucket.example.id
  name   = "example"

  public_access_block_configuration {
    block_public_acls       = true
    block_public_policy     = false
    ignore_public_acls      = true
    restrict_public_buckets = false
  }

  lifecycle {
    ignore_changes = [policy]
  }
}

resource "aws_s3control_access_point_policy" "example" {
  access_point_arn = aws_s3_access_point.example.arn

  policy = jsonencode({
    Version = "2008-10-17"
    Statement = [
      {
        Effect = "Allow"
        Action = "s3:GetObjectTagging"
        Principal = {
          AWS = "*"
        }
        Resource = "${aws_s3_access_point.example.arn}/object/*"
      }
    ]
  })
}

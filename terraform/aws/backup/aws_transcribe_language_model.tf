data "aws_iam_policy_document" "example" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["transcribe.amazonaws.com"]
    }
  }
}

resource "aws_iam_role" "example" {
  name               = "example"
  assume_role_policy = data.aws_iam_policy_document.example.json
}

resource "aws_iam_role_policy" "test_policy" {
  name = "example"
  role = aws_iam_role.example.id

  policy = jsonencode({ Version = "2012-10-17"
    Statement = [
      {
        Action = [
          "s3:GetObject",
          "s3:ListBucket",
        ]
        Effect   = "Allow"
        Resource = ["*"]
      },
    ]
  })
}

resource "aws_s3_bucket" "example" {
  bucket        = "example-transcribe"
  force_destroy = true
}

resource "aws_s3_object" "object" {
  bucket = aws_s3_bucket.example.id
  key    = "transcribe/test1.txt"
  source = "test1.txt"
}

resource "aws_transcribe_language_model" "pike" {
  model_name      = "example"
  base_model_name = "NarrowBand"

  input_data_config {
    data_access_role_arn = aws_iam_role.example.arn
    s3_uri               = "s3://${aws_s3_bucket.example.id}/transcribe/"
  }

  language_code = "en-GB"

  tags = {
    ENVIRONMENT = "development"
    pike        = "permissions"
  }
}

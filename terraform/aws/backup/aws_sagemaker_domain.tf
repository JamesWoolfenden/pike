resource "aws_sagemaker_domain" "pike" {
  domain_name = "example"
  auth_mode   = "IAM"
  vpc_id      = "vpc-06074a092930bc809"
  subnet_ids  = [aws_subnet.example.id]

  default_user_settings {
    execution_role = aws_iam_role.example.arn
  }
}

resource "aws_iam_role" "example" {
  name               = "example"
  path               = "/"
  assume_role_policy = data.aws_iam_policy_document.domain.json
}

data "aws_iam_policy_document" "domain" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["sagemaker.amazonaws.com"]
    }
  }
}


resource "aws_subnet" "example" {
  vpc_id     = "vpc-06074a092930bc809"
  cidr_block = "10.0.1.0/24"
}

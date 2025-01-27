resource "aws_glue_dev_endpoint" "pike" {
  name     = "foo"
  role_arn = aws_iam_role.example.arn
}

resource "aws_iam_role" "example" {
  name               = "AWSGlueServiceRole-foo"
  assume_role_policy = data.aws_iam_policy_document.example.json
}

data "aws_iam_policy_document" "example" {
  statement {
    actions = ["sts:AssumeRole"]

    principals {
      type        = "Service"
      identifiers = ["glue.amazonaws.com"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "example-AWSGlueServiceRole" {
  policy_arn = "arn:aws:iam::aws:policy/service-role/AWSGlueServiceRole"
  role       = aws_iam_role.example.name
}

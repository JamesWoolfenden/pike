resource "aws_iam_policy" "wildcard_jsonencode" {
  name = "x"
  policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Effect   = "Allow"
      Action   = ["*"]
      Resource = "*"
    }]
  })
}

resource "aws_iam_role_policy" "service_wildcard" {
  name = "x"
  role = "r"
  policy = jsonencode({
    Statement = [{
      Sid      = "S3All"
      Effect   = "Allow"
      Action   = "s3:*"
      Resource = ["*"]
    }]
  })
}

resource "aws_iam_user_policy" "escalation_heredoc" {
  name   = "x"
  user   = "u"
  policy = <<EOF
{
  "Version": "2012-10-17",
  "Statement": [
    {"Effect": "Allow", "Action": "iam:PassRole", "Resource": "*"}
  ]
}
EOF
}

resource "aws_iam_group_policy" "not_action" {
  name  = "x"
  group = "g"
  policy = jsonencode({
    Statement = {
      Effect    = "Allow"
      NotAction = ["iam:*"]
      Resource  = "*"
    }
  })
}

resource "aws_iam_policy" "from_file" {
  name   = "x"
  policy = file("policy.json")
}

resource "aws_iam_policy" "interpolated" {
  name   = "x"
  policy = data.aws_iam_policy_document.doc.json
}

data "aws_iam_policy_document" "doc" {
  statement {
    sid       = "Admin"
    actions   = ["*"]
    resources = ["*"]
  }
  statement {
    actions   = ["iam:PutRolePolicy"]
    resources = ["arn:aws:iam::123456789012:role/app"]
    condition {
      test     = "StringEquals"
      variable = "aws:PrincipalTag/team"
      values   = ["platform"]
    }
  }
}

resource "aws_iam_role_policy_attachment" "admin" {
  role       = "r"
  policy_arn = "arn:aws:iam::aws:policy/AdministratorAccess"
}

resource "aws_iam_user_policy_attachment" "s3full" {
  user       = "u"
  policy_arn = "arn:aws:iam::aws:policy/AmazonS3FullAccess"
}

resource "aws_iam_role" "with_inline" {
  name = "r"
  assume_role_policy = jsonencode({
    Statement = [{ Effect = "Allow", Principal = { Service = "ec2.amazonaws.com" }, Action = "sts:AssumeRole" }]
  })
  inline_policy {
    name = "p"
    policy = jsonencode({
      Statement = [{ Effect = "Allow", Action = ["s3:PutObject"], Resource = "*" }]
    })
  }
}

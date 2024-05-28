resource "aws_datazone_domain" "pike" {
  name                  = "pike"
  domain_execution_role = aws_iam_role.domain_execution_role.arn

  tags = {
    pike = "permissions"
  }
}

resource "aws_iam_role" "domain_execution_role" {
  name = "my_domain_execution_role"
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = ["sts:AssumeRole", "sts:TagSession"]
        Effect = "Allow"
        Principal = {
          Service = "datazone.amazonaws.com"
        }
      },
      {
        Action = ["sts:AssumeRole", "sts:TagSession"]
        Effect = "Allow"
        Principal = {
          Service = "cloudformation.amazonaws.com"
        }
      },
    ]
  })

  inline_policy {
    name = "domain_execution_policy"
    policy = jsonencode({
      Version = "2012-10-17"
      Statement = [
        {
          # Consider scoping down
          Action = [
            "datazone:*",
            "ram:*",
            "sso:*",
            "kms:*",
          ]
          Effect   = "Allow"
          Resource = "*"
        },
      ]
    })
  }
}

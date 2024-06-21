resource "aws_db_proxy" "pike" {
  name                   = "example"
  debug_logging          = false
  engine_family          = "MYSQL"
  idle_client_timeout    = 1800
  require_tls            = true
  role_arn               = aws_iam_role.example.arn
  vpc_security_group_ids = ["sg-05b27cb61c9c46bd2"]
  vpc_subnet_ids         = data.aws_subnets.public.ids

  auth {
    auth_scheme = "SECRETS"
    description = "example"
    iam_auth    = "DISABLED"
    secret_arn  = aws_secretsmanager_secret.pike.arn
  }

  tags = {
    Name = "example"
    Key  = "value"
  }
}

resource "aws_iam_role" "example" {
  assume_role_policy = jsonencode(
    {
      "Version" : "2012-10-17",
      "Statement" : [
        {
          "Effect" : "Allow",
          "Principal" : { "AWS" : "arn:aws:iam::680235478471:root" },
          "Action" : "sts:AssumeRole",
        }
      ]
    }
  )
}

resource "aws_secretsmanager_secret" "pike" {
  description = "My secret"
  //kms_key_id  = "arn:aws:kms:eu-west-2:680235478471:key/00aba012-359a-4c01-aec7-2764e5bbbb8e"
  name = "pike"

  recovery_window_in_days = 0
  tags = {
    "pike" = "permissions"
  }
}

data "aws_subnets" "public" {
  filter {
    name   = "vpc-id"
    values = ["vpc-0c33dc8cd64f408c4"]
  }
}

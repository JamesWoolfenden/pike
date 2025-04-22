resource "aws_dynamodb_table" "example" {
  name           = "dynamodb-table-example"
  read_capacity  = 1
  write_capacity = 1
  hash_key       = "example"

  attribute {
    name = "example"
    type = "S"
  }

  point_in_time_recovery {
    enabled = true
  }
}

data "aws_subnets" "pike" {

}

resource "aws_redshiftserverless_namespace" "example" {
  namespace_name = "redshift-example"
}

resource "aws_redshiftserverless_workgroup" "example" {
  namespace_name      = aws_redshiftserverless_namespace.example.namespace_name
  workgroup_name      = "example-workgroup"
  base_capacity       = 8
  publicly_accessible = false

  subnet_ids = data.aws_subnets.pike.ids

  config_parameter {
    parameter_key   = "enable_case_sensitive_identifier"
    parameter_value = "true"
  }

  config_parameter {
    parameter_key   = "auto_mv"
    parameter_value = "true"
  }

  config_parameter {
    parameter_key   = "datestyle"
    parameter_value = "ISO, MDY"
  }

  config_parameter {
    parameter_key   = "enable_user_activity_logging"
    parameter_value = "true"
  }
  config_parameter {
    parameter_key   = "max_query_execution_time"
    parameter_value = "14400"
  }
  config_parameter {
    parameter_key   = "query_group"
    parameter_value = "default"
  }
  config_parameter {
    parameter_key   = "require_ssl"
    parameter_value = "true"
  }
  config_parameter {
    parameter_key   = "search_path"
    parameter_value = "$user, public"
  }
  config_parameter {
    parameter_key   = "use_fips_ssl"
    parameter_value = "false"
  }

}

resource "aws_redshift_integration" "pike" {
  integration_name = "pike"
  description      = "test"
  source_arn       = aws_dynamodb_table.example.arn
  target_arn       = aws_redshiftserverless_namespace.example.arn

  kms_key_id = aws_kms_key.example.key_id

  # additional_encryption_context = {
  #   "example" : "test",
  # }

  tags = {
    pike    = "permissions"
    another = "one"
  }

  depends_on = [
    aws_kms_key_policy.example
  ]
}


data "aws_caller_identity" "current" {}

resource "aws_kms_key" "example" {
  description             = "example-redshift2"
  deletion_window_in_days = 10
}

resource "aws_kms_key_policy" "example" {
  key_id = aws_kms_key.example.id

  policy = jsonencode({
    Id      = "key-consolepolicy-3"
    Version = "2012-10-17"
    Statement = [
      {
        Sid    = "Enable IAM User Permissions"
        Effect = "Allow"
        Principal = {
          AWS = "arn:aws:iam::${data.aws_caller_identity.current.account_id}:root"
        }
        Action   = "kms:*"
        Resource = "*"
      },
      {
        Sid    = "Allow use of the key"
        Effect = "Allow"
        Principal = {
          AWS = [
            "arn:aws:iam::680235478471:role/redshiftScheduler",
            "arn:aws:iam::680235478471:role/aws-service-role/redshift.amazonaws.com/AWSServiceRoleForRedshift",
          ]
        }
        Action = [
          "kms:Encrypt",
          "kms:Decrypt",
          "kms:ReEncrypt*",
          "kms:GenerateDataKey*",
          "kms:DescribeKey",
          "kms:CreateGrant",
        ]
        Resource = "*"
      },
      {
        Action = [
          "kms:CreateGrant",
          "kms:ListGrants",
          "kms:RevokeGrant",
        ]
        Condition = {
          Bool = {
            "kms:GrantIsForAWSResource" = "true"
          }
        }
        Effect = "Allow"
        Principal = {
          AWS = [
            "arn:aws:iam::680235478471:role/redshiftScheduler",
            "arn:aws:iam::680235478471:role/aws-service-role/redshift.amazonaws.com/AWSServiceRoleForRedshift",
          ]
        }
        Resource = "*"
        Sid      = "Allow attachment of persistent resources"
      },

    ]
  })
}

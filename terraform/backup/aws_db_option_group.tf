resource "aws_db_option_group" "example" {
  name                     = "option-group-pike"
  option_group_description = "Terraform Option Group"
  engine_name              = "sqlserver-ee"
  major_engine_version     = "11.00"

  option {
    db_security_group_memberships  = []
    option_name                    = "SQLSERVER_AUDIT"
    port                           = 0
    vpc_security_group_memberships = []

    option_settings {
      name  = "IAM_ROLE_ARN"
      value = "arn:aws:iam::680235478471:role/aws-service-role/rds.amazonaws.com/AWSServiceRoleForRDS"
    }
    option_settings {
      name  = "S3_BUCKET_ARN"
      value = "arn:aws:s3:::testbucketineu-west2/rds"
    }
  }

  option {
    option_name = "TDE"
    port        = 0
  }

  tags = {
    pike = "permissions"
  }
}

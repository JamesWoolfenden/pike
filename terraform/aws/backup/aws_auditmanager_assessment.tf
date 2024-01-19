resource "aws_auditmanager_assessment" "pike" {
  name = "example"

  assessment_reports_destination {
    destination      = "s3://${aws_s3_bucket.test.id}"
    destination_type = "S3"
  }

  framework_id = aws_auditmanager_framework.pike.id

  roles {
    role_arn  = aws_iam_role.example.arn
    role_type = "PROCESS_OWNER"
  }

  scope {
    aws_accounts {
      id = data.aws_caller_identity.current.account_id
    }
    aws_services {
      service_name = "S3"
    }
  }

  tags = {
    pike = "permissions"
  }
}


data "aws_caller_identity" "current" {}

resource "aws_s3_bucket" "test" {
  bucket = "reports-${data.aws_caller_identity.current.account_id}"
}

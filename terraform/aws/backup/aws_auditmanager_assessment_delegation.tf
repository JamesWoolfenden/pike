resource "aws_auditmanager_assessment_delegation" "pike" {
  assessment_id  = aws_auditmanager_assessment.pike.id
  role_arn       = aws_iam_role.example.arn
  role_type      = "RESOURCE_OWNER"
  control_set_id = "example"
}


resource "aws_iam_role" "example" {
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [
      {
        Action = "sts:AssumeRole"
        Effect = "Allow"
        Sid    = ""
        Principal = {
          Service = "ec2.amazonaws.com"
        }
      },
    ]
  })
}

data "aws_ssoadmin_principal_application_assignments" "pike" {
  instance_arn   = "arn:aws:ec2:us-east-1:123456789012:instance/fred"
  principal_type = "USER"
  principal_id   = "asdasd"

}

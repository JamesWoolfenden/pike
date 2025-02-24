resource "aws_qbusiness_application" "pike" {
  display_name                 = "pike"
  iam_service_role_arn         = aws_iam_role.example.arn
  identity_center_instance_arn = tolist(data.aws_ssoadmin_instances.example.arns)[0]

  attachments_configuration {
    attachments_control_mode = "ENABLED"
  }
}

resource "aws_auditmanager_framework_share" "pike" {
  destination_account = "680235478471"
  destination_region  = "us-east-1"
  framework_id        = aws_auditmanager_framework.pike.id
}

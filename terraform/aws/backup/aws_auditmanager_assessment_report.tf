resource "aws_auditmanager_assessment_report" "pike" {
  name          = "example"
  assessment_id = aws_auditmanager_assessment.pike.id
}

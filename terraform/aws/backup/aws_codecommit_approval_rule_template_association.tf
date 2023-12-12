resource "aws_codecommit_approval_rule_template_association" "pike" {
  approval_rule_template_name = aws_codecommit_approval_rule_template.pike.name
  repository_name             = aws_codecommit_repository.test.repository_name
}

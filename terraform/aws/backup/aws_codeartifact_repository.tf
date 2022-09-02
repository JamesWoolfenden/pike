resource "aws_codeartifact_repository" "pike" {
  repository  = "pike"
  domain      = aws_codeartifact_domain.pike.domain
  description = "for pike"
  tags = {
    pike = "permission"
  }
}

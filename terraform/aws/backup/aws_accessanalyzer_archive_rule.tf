resource "aws_accessanalyzer_archive_rule" "pike" {
  analyzer_name = aws_accessanalyzer_analyzer.pike.analyzer_name
  rule_name     = "example-rule"

  filter {
    criteria = "condition.aws:UserId"
    eq       = ["userid"]
  }

  filter {
    criteria = "error"
    exists   = true
  }

  filter {
    criteria = "isPublic"
    eq       = ["false"]
  }
}

resource "aws_accessanalyzer_analyzer" "pike" {
  analyzer_name = "pike"
}

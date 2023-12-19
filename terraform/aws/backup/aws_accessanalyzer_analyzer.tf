resource "aws_accessanalyzer_analyzer" "pike" {
  analyzer_name = "example"
  type          = "ORGANIZATION"
  tags = {
    pike = "permissions"
  }
}

resource "aws_glue_data_quality_ruleset" "pike" {
  name    = "example"
  ruleset = "Rules = [Completeness \"colA\" between 0.4 and 0.8]"
}

resource "aws_waf_regex_pattern_set" "pike" {
  name                  = "example"
  regex_pattern_strings = ["one", "two"]
}

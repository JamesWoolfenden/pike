data "aws_wafv2_regex_pattern_set" "pike" {
  name  = "pike"
  scope = "CLOUDFRONT"
}

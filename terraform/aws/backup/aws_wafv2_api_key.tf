resource "aws_wafv2_api_key" "pike" {
  scope         = "REGIONAL"
  token_domains = ["example.com"]
}

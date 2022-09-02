data "aws_wafv2_web_acl" "pike" {
  name  = "pike"
  scope = "CLOUDFRONT"
}

output "web_acl" {
  value = data.aws_wafv2_web_acl.pike
}

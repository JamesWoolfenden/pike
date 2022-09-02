data "aws_wafv2_ip_set" "pike" {
  name  = "pike"
  scope = "CLOUDFRONT"
}

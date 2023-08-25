module "waf2" {
  source      = "git::https://github.com/JamesWoolfenden/terraform-aws-waf2.git?ref=ca016d169646e640ea6d648c0def99f94d3f01a3"
  kms_key_arn = aws_kms_key.waf2.arn
}

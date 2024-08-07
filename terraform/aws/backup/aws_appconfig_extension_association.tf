resource "aws_appconfig_extension_association" "pike" {
  extension_arn = aws_appconfig_extension.pike.arn
  resource_arn  = aws_appconfig_application.pike.arn
}

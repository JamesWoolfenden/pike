data "aws_appconfig_application" "pike" {
  name = "pike"
}

output "aws_appconfig_application" {
  value = data.aws_appconfig_application.pike
}

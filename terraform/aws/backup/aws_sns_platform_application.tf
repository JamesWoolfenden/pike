resource "aws_sns_platform_application" "pike" {
  name                = "pike"
  platform            = "GCM"
  platform_credential = "<GCM API KEY>"
}

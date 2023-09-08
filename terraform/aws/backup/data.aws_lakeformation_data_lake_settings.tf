data "aws_lakeformation_data_lake_settings" "pike" {}
output "dl_settings" {
  value = data.aws_lakeformation_data_lake_settings.pike
}
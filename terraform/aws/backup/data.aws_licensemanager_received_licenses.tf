data "aws_licensemanager_received_licenses" "pike" {}

output "licences" {
  value = data.aws_licensemanager_received_licenses.pike
}

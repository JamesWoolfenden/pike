data "aws_licensemanager_grants" "pike" {}

output "grants" {
  value = data.aws_licensemanager_grants.pike
}

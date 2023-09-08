data "aws_licensemanager_received_license" "pike" {
  license_arn = "arn:aws:license-manager::680235478471:license:l-1534f5bfb27f41bc8ad622b95e20e39d"
}

output "licence" {
  value = data.aws_licensemanager_received_license.pike
}

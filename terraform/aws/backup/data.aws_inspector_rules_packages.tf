data "aws_inspector_rules_packages" "pike" {}

output "packages" {
  value = data.aws_inspector_rules_packages.pike
}

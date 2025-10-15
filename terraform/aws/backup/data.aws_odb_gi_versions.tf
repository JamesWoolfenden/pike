data "aws_odb_gi_versions" "pike" {
}

output "aws_odb_gi_versions" {
  value = data.aws_odb_gi_versions.pike
}

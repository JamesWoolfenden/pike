data "aws_synthetics_runtime_versions" "pike" {
}

output "aws_synthetics_runtime_versions" {
  value = data.aws_synthetics_runtime_versions.pike
}

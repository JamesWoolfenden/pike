data "aws_neptune_engine_version" "pike" {}

output "version" {
  value = data.aws_neptune_engine_version.pike
}

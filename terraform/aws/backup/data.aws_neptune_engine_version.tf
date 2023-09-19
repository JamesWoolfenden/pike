data "aws_neptune_engine_version" "pike" {}

output "engine_version" {
  value = data.aws_neptune_engine_version.pike
}

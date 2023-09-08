data "aws_resourcegroupstaggingapi_resources" "pike" {}

output "tags" {
  value = data.aws_resourcegroupstaggingapi_resources.pike
}

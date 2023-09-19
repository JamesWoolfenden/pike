data "aws_resourcegroupstaggingapi_resources" "pike" {}

output "api_tags" {
  value = data.aws_resourcegroupstaggingapi_resources.pike
}

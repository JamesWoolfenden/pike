data "aws_resourcegroupstaggingapi_required_tags" "pike" {
}

output "aws_resourcegroupstaggingapi_required_tags" {
  value = data.aws_resourcegroupstaggingapi_required_tags.pike
}

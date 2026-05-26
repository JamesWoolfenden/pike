data "aws_opensearchserverless_collection_groups" "pike" {
}

output "aws_opensearchserverless_collection_groups" {
  value = data.aws_opensearchserverless_collection_groups.pike
}

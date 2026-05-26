data "aws_opensearchserverless_collection_group" "pike" {
  name = "pike"
}

output "aws_opensearchserverless_collection_group" {
  value = data.aws_opensearchserverless_collection_group.pike
}

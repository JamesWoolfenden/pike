data "aws_devopsguru_resource_collection" "pike" {
  type = "AWS_SERVICE"
}

output "aws_devopsguru_resource_collection" {
  value = data.aws_devopsguru_resource_collection.pike
}

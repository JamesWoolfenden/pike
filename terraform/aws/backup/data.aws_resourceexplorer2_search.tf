data "aws_resourceexplorer2_search" "pike" {
  query_string = "region:eu-west-2"
}

output "aws_resourceexplorer2_search" {
  value = data.aws_resourceexplorer2_search.pike
}

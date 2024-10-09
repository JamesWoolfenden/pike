resource "aws_quicksight_group" "pike" {
  group_name = "pike"
  namespace  = aws_quicksight_namespace.pike.namespace
}

resource "aws_quicksight_group_membership" "pike" {
  group_name  = aws_quicksight_group.pike.group_name
  member_name = "jameswoolfenden"
}

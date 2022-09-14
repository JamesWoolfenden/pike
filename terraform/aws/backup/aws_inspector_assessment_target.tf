resource "aws_inspector_assessment_target" "pike" {
  name               = "assessment target"
  resource_group_arn = aws_inspector_resource_group.pike.arn
}

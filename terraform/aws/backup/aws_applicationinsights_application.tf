resource "aws_applicationinsights_application" "pike" {
  resource_group_name    = aws_resourcegroups_group.pike.name
  auto_config_enabled    = true
  auto_create            = true
  ops_center_enabled     = true
  ops_item_sns_topic_arn = data.aws_sns_topic.pike.arn
  tags = {
    pike = "permissions"
  }
}


data "aws_sns_topic" "pike" {
  name = "pike"
}

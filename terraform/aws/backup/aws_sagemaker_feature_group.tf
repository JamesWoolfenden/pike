resource "aws_sagemaker_feature_group" "pike" {
  feature_group_name             = "example"
  record_identifier_feature_name = "example"
  event_time_feature_name        = "example"
  role_arn                       = aws_iam_role.example.arn

  feature_definition {
    feature_name = "example"
    feature_type = "String"
  }

  online_store_config {
    enable_online_store = true
  }

  tags = {
    pike = "permissions"
  }
}

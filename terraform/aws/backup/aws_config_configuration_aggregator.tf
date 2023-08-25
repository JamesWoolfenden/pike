resource "aws_config_configuration_aggregator" "pike" {
  name = "example"

  account_aggregation_source {
    account_ids = ["123456789012"]
    regions     = ["eu-west-2"]
  }
}

data "aws_redshift_producer_data_shares" "pike" {
  producer_arn = "arn:aws:redshift:us-west-2:680235478471:datashare:3072dae5-022b-4d45-9cd3-01f010aae4b2/example_share"
}

output "aws_redshift_producer_data_shares" {
  value = data.aws_redshift_producer_data_shares.pike
}

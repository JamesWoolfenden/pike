data "aws_redshift_data_shares" "pike" {
}

output "aws_redshift_data_shares" {
  value = data.aws_redshift_data_shares.pike
}

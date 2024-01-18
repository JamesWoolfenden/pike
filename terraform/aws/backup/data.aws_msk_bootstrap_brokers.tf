data "aws_msk_bootstrap_brokers" "pike" {
  cluster_arn = "arn:aws:kafka:us-east-1:680235478471:cw-1234567890"
}

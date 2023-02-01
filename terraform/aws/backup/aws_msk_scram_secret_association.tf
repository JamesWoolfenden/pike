resource "aws_msk_scram_secret_association" "pike" {
  cluster_arn     = aws_msk_cluster.pike.arn
  secret_arn_list = ["arn:aws:secretsmanager:eu-west-2:680235478471:secret:AmazonMSK_shizzle-CBDwZQ"]
}

resource "aws_ram_principal_association" "pike" {
  principal          = "111111111111"
  resource_share_arn = aws_ram_resource_share.example.arn
}
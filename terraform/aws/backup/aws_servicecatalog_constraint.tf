resource "aws_servicecatalog_constraint" "pike" {
  description  = "Back off, man. I'm a scientist."
  portfolio_id = aws_servicecatalog_portfolio.pike.id
  product_id   = aws_servicecatalog_product.pike.id
  type         = "LAUNCH"

  parameters = jsonencode({
    "RoleArn" : "arn:aws:iam::680235478471:role/LaunchRole"
  })
}

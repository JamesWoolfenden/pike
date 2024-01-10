resource "aws_servicecatalog_portfolio_share" "pike" {
  principal_id = "680235478471"
  portfolio_id = aws_servicecatalog_portfolio.pike.id
  type         = "ACCOUNT"
}
resource "aws_servicecatalog_portfolio" "pike" {
  name          = "My App Portfolio"
  description   = "List of my organizations apps"
  provider_name = "Brett"
}

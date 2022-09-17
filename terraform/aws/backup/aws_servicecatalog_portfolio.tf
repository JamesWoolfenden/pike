resource "aws_servicecatalog_portfolio" "pike" {
  name          = "Pike App Portfolio"
  description   = "List of my organisations apps"
  provider_name = "Pike"
  tags = {
    pike = "permissions"
    //delete="me"
  }
}

output "portfolio" {
  value = aws_servicecatalog_portfolio.pike
}

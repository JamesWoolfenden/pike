data "aws_billing_views" "pike" {
}

output "aws_billing_views" {
  value = data.aws_billing_views.pike
}

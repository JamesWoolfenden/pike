data "aws_route_tables" "example" {}

output "route_tables" {
  value = data.aws_route_tables.example
}

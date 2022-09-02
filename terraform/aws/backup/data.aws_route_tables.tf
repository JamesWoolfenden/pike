data "aws_route_tables" "example" {}

output "routes" {
  value = data.aws_route_tables.example
}

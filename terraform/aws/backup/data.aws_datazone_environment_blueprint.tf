data "aws_datazone_environment_blueprint" "pike" {
  domain_id = aws_datazone_domain.pike.id
  managed   = true
  name      = "DefaultDataLake"
}

output "aws_datazone_environment_blueprint" {
  value = data.aws_datazone_environment_blueprint.pike
}

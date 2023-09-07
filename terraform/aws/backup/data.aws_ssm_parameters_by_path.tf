data "aws_ssm_parameters_by_path" "pike" {
  path="/pike"
}


output "path" {
  value = data.aws_ssm_parameters_by_path.pike
  sensitive = true
}
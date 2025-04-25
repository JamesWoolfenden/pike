data "aws_fis_experiment_templates" "pike" {
}

output "aws_fis_experiment_templates" {
  value = data.aws_fis_experiment_templates.pike
}

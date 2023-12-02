data "aws_emr_supported_instance_types" "example" {
  release_label = "emr-6.15.0"
}

output "label" {
  value = data.aws_emr_supported_instance_types.example
}

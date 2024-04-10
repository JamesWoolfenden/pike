data "aws_batch_job_definition" "pike" {
  arn = "arn:aws:batch:us-east-1:680235478471:job-definition/example"
}

output "aws_batch_job_definition" {
  value = data.aws_batch_job_definition.pike
}

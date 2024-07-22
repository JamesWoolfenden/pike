data "aws_timestreamwrite_database" "pike" {
  provider = aws.central
  name     = "pike"
}

output "aws_timestreamwrite_database" {
  value = data.aws_timestreamwrite_database.pike
}

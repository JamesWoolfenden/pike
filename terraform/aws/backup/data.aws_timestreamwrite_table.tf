data "aws_timestreamwrite_table" "pike" {
  provider      = aws.central
  database_name = "pike"
  name          = "pike"
}

output "aws_timestreamwrite_table" {
  value = data.aws_timestreamwrite_table.pike
}

data "aws_arn" "pike" {
  arn = "arn:aws:rds:eu-west-1:680235478471:db:mysql-db"
}

output "details" {
  value = data.aws_arn.pike
}

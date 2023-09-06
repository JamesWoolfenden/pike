data "aws_iam_access_keys" "pike" {
  user = "jameswoolfenden"
}

output "keys" {
  value = data.aws_iam_access_keys.pike
}

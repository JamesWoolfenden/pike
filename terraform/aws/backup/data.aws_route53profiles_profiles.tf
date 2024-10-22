data "aws_route53profiles_profiles" "pike" {
}

output "aws_route53profiles_profiles" {
  value = data.aws_route53profiles_profiles.pike
}

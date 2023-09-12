data "aws_lakeformation_resource" "pike" {
  arn = "arn:aws:s3:::tf-acc-test-9151654063908211878"
}

output "resource" {
  value = data.aws_lakeformation_resource.pike
}

data "aws_regions" "pike" {}

output "regions" {
  value = data.aws_regions.pike
}

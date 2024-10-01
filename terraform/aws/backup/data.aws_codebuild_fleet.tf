data "aws_codebuild_fleet" "pike" {
  name = "pike"
}

output "aws_codebuild_fleet" {
  value = data.aws_codebuild_fleet.pike
}

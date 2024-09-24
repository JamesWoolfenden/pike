data "aws_quicksight_analysis" "pike" {
  analysis_id = "example-id"
}

output "aws_quicksight_analysis" {
  value = data.aws_quicksight_analysis.pike
}

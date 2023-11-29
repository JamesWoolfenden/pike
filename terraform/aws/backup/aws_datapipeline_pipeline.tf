resource "aws_datapipeline_pipeline" "pike" {
  name        = "tf-pipeline-default"
  description = "pike"
  tags = {
    pike = "permission"
  }
}

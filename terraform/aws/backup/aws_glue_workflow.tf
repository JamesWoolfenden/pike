resource "aws_glue_workflow" "pike" {
  name        = "pike"
  description = ""
  tags = {
    pike   = "permission"
    delete = "me"
  }
}

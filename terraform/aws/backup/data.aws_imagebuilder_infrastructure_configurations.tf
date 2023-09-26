data "aws_imagebuilder_infrastructure_configurations" "pike" {
  filter {
    name   = "name"
    values = ["example"]
  }
}

resource "aws_imagebuilder_container_recipe" "pike" {
  name    = "example"
  version = "1.0.0"

  container_type = "DOCKER"
  parent_image   = "jameswoolfenden/pike"

  target_repository {
    repository_name = "arn:aws:ecr:eu-west-2:680235478471:repository/pike"
    service         = "ECR"
  }

  component {
    component_arn = aws_imagebuilder_component.example.arn

    parameter {
      name  = "Parameter1"
      value = "Value1"
    }

    parameter {
      name  = "Parameter2"
      value = "Value2"
    }
  }

  dockerfile_template_data = <<EOF
FROM {{{ imagebuilder:parentImage }}}
{{{ imagebuilder:environments }}}
{{{ imagebuilder:components }}}
EOF
}

resource "aws_imagebuilder_component" "example" {
  data = yamlencode({
    phases = [{
      name = "build"
      steps = [{
        action = "ExecuteBash"
        inputs = {
          commands = ["echo 'hello world'"]
        }
        name      = "example"
        onFailure = "Continue"
      }]
    }]
    schemaVersion = 1.0
  })
  name     = "example"
  platform = "Linux"
  version  = "1.0.0"
}

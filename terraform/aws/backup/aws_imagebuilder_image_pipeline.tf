resource "aws_imagebuilder_image_pipeline" "pike" {
  image_recipe_arn                 = aws_imagebuilder_image_recipe.pike.arn
  infrastructure_configuration_arn = aws_imagebuilder_infrastructure_configuration.pike.arn
  name                             = "example"

  schedule {
    schedule_expression = "cron(0 0 * * ? *)"
  }

  tags = {
    pike   = "permission"
    delete = "me"
  }
}

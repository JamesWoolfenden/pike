resource "aws_imagebuilder_image" "pike" {
  distribution_configuration_arn   = aws_imagebuilder_distribution_configuration.pike.arn
  image_recipe_arn                 = aws_imagebuilder_image_recipe.pike.arn
  infrastructure_configuration_arn = aws_imagebuilder_infrastructure_configuration.pike.arn

  tags = {
    pike   = "permission"
    delete = "me"
  }
}

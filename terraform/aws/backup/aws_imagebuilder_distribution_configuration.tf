resource "aws_imagebuilder_distribution_configuration" "pike" {
  name = "pike"

  distribution {
    ami_distribution_configuration {
      ami_tags = {
        CostCenter = "IT"
      }

      name = "pike-{{ imagebuilder:buildDate }}"

      launch_permission {
        user_ids = ["680235478471"]
      }
    }

    #    launch_template_configuration {
    #      launch_template_id = "lt-0aaa1bcde2ff3456"
    #    }

    region = "eu-west-2"
  }
  tags = {
    pike   = "permission"
    delete = "me"
  }
}

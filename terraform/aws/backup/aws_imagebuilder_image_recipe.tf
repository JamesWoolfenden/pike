resource "aws_imagebuilder_image_recipe" "pike" {
  block_device_mapping {
    device_name = "/dev/xvdb"

    ebs {
      delete_on_termination = true
      volume_size           = 100
      volume_type           = "gp2"
    }
  }

  component {
    component_arn = aws_imagebuilder_component.pike.arn

    #    parameter {
    #      name  = "Parameter1"
    #      value = "Value1"
    #    }
    #
    #    parameter {
    #      name  = "Parameter2"
    #      value = "Value2"
    #    }
  }

  name         = "pike"
  parent_image = "arn:${data.aws_partition.current.partition}:imagebuilder:${data.aws_region.current.name}:aws:image/amazon-linux-2-x86/x.x.x"
  version      = "1.0.0"

  tags = {
    pike = "permissions"
  }
}

data "aws_partition" "current" {}

data "aws_region" "current" {}

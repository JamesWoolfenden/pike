data "aws_region" "current" {}

data "aws_partition" "current" {}

resource "aws_iam_role" "example" {
  assume_role_policy = jsonencode({
    Version = "2012-10-17"
    Statement = [{
      Action = "sts:AssumeRole"
      Effect = "Allow"
      Principal = {
        Service = "imagebuilder.${data.aws_partition.current.dns_suffix}"
      }
    }]
  })
  name = "example"
}

resource "aws_iam_role_policy_attachment" "example" {
  policy_arn = "arn:${data.aws_partition.current.partition}:iam::aws:policy/service-role/EC2ImageBuilderLifecycleExecutionPolicy"
  role       = aws_iam_role.example.name
}

resource "aws_imagebuilder_lifecycle_policy" "example" {
  name           = "name"
  description    = "Example description"
  execution_role = aws_iam_role.example.arn
  resource_type  = "AMI_IMAGE"
  policy_detail {
    action {
      type = "DELETE"
    }
    filter {
      type            = "AGE"
      value           = 6
      retain_at_least = 10
      unit            = "YEARS"
    }
  }
  resource_selection {
    tag_map = {
      "key1" = "value1"
      "key2" = "value2"
    }
  }

  depends_on = [aws_iam_role_policy_attachment.example]
}

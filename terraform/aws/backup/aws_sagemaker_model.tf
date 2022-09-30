resource "aws_sagemaker_model" "pike" {
  name                     = "pike"
  execution_role_arn       = "arn:aws:iam::680235478471:role/sagemakerService"
  enable_network_isolation = true

  vpc_config {
    security_group_ids = ["sg-0403ba8bc0ba7876b"]
    subnets            = ["subnet-04338b6369d8288a5", "subnet-0d99e0f4e4f1ff84f"]
  }
  primary_container {
    image = data.aws_sagemaker_prebuilt_ecr_image.test.registry_path
  }
  tags = {
    pike = "permissions"
    #    delete="me"
  }
}

data "aws_sagemaker_prebuilt_ecr_image" "test" {
  repository_name = "kmeans"
}

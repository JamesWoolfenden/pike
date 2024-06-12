resource "aws_sagemaker_image_version" "pike" {
  image_name = aws_sagemaker_image.pike.id
  base_image = "012345678912.dkr.ecr.us-west-2.amazonaws.com/image:latest"
}

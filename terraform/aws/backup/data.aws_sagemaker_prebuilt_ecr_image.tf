data "aws_sagemaker_prebuilt_ecr_image" "pike" {
  repository_name = "sagemaker-scikit-learn"
  image_tag       = "2.2-1.0.11.0"
}

output "sage_image" {
  value = data.aws_sagemaker_prebuilt_ecr_image.pike
}

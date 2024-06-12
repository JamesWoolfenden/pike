resource "aws_sagemaker_image" "pike" {
  image_name = "pike"
  role_arn   = aws_iam_role.example.arn
}

resource "aws_iam_service_specific_credential" "pike" {
  service_name = "codecommit.amazonaws.com"
  user_name    = "jameswoolfenden"
}

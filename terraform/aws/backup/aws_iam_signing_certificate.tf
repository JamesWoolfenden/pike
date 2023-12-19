resource "aws_iam_signing_certificate" "pike" {
  user_name        = "jameswoolfenden"
  certificate_body = file("sample.pem")
}

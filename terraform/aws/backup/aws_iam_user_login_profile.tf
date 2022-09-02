resource "aws_iam_user_login_profile" "example" {
  user                    = "basic"
  pgp_key                 = "keybase:jameswoolfenden"
  password_length         = 10
  password_reset_required = true
}

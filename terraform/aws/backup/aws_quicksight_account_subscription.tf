resource "aws_quicksight_account_subscription" "pike" {
  account_name          = "quicksight-terraform"
  authentication_method = "IAM_AND_QUICKSIGHT"
  edition               = "ENTERPRISE"
  notification_email    = "notification@email.com"
}

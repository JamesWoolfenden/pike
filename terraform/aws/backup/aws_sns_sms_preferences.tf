resource "aws_sns_sms_preferences" "pike" {
  monthly_spend_limit = "1"
  default_sms_type    = "Transactional"
}

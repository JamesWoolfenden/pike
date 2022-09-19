resource "aws_iam_account_password_policy" "pike" {
  allow_users_to_change_password = true
  hard_expiry                    = false
  max_password_age               = 0
  minimum_password_length        = 6
  password_reuse_prevention      = 0
  require_lowercase_characters   = false
  require_numbers                = false
  require_symbols                = false
  require_uppercase_characters   = false
}

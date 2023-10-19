resource "aws_organizations_policy_attachment" "account" {
  policy_id = aws_organizations_policy.example.id
  target_id = "123456789012"
}

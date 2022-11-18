resource "aws_backup_selection" "pike" {
  iam_role_arn = "arn:aws:iam::680235478471:role/aws-service-role/backup.amazonaws.com/AWSServiceRoleForBackup"
  name         = "tf_example_backup_selection"
  plan_id      = aws_backup_plan.pike.id
  selection_tag {
    type  = "STRINGEQUALS"
    key   = "foo"
    value = "bar"
  }
}

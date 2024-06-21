resource "aws_db_instance_role_association" "pike" {
  db_instance_identifier = aws_db_instance.pike.identifier
  feature_name           = "" #"S3_INTEGRATION"
  role_arn               = aws_iam_role.example.arn
}

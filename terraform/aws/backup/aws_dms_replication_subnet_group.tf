resource "aws_dms_replication_subnet_group" "pike" {
  replication_subnet_group_description = "Example replication subnet group"
  replication_subnet_group_id          = "example-dms-replication-subnet-group-tf"


  subnet_ids = [
    "subnet-09ff91b5b0adb1fd4",
    "subnet-05e87623a2a5c41f0",
  ]

  tags = {
    Name = "example"
  }
  depends_on = [aws_iam_role_policy_attachment.example]
}



resource "aws_iam_role_policy_attachment" "example" {
  role       = aws_iam_role.dms-vpc-role.name
  policy_arn = "arn:aws:iam::aws:policy/service-role/AmazonDMSVPCManagementRole"
}

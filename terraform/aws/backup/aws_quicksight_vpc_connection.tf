resource "aws_quicksight_vpc_connection" "pike" {
  vpc_connection_id  = "example-connection-id"
  name               = "Example Connection"
  role_arn           = aws_iam_role.vpc_connection_role.arn
  security_group_ids = ["sg-00000000000000000"]
  subnet_ids = [
    "subnet-00000000000000000",
    "subnet-00000000000000001",
  ]
}

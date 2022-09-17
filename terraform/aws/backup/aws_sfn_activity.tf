resource "aws_sfn_activity" "pike" {
  name = "pike"

  tags = {
    pike = "permissions"
    // delete="me"
  }
}

resource "aws_codedeploy_app" "pike" {
  name             = "pike"
  compute_platform = "Server"
  tags = {
    pike = "permissions"
  }
}

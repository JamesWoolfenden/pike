resource "aws_appstream_user" "pike" {
  authentication_type = "USERPOOL"
  user_name           = "james@bridgecrew.io"
  first_name          = "james"
  last_name           = "woolfenden"
}

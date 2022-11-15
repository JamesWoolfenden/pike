resource "aws_cloudwatch_event_connection" "pike" {
  name               = "ngrok-connection"
  description        = "A connection description"
  authorization_type = "API_KEY"

  auth_parameters {
    api_key {
      key   = "x-signature"
      value = "1234"
    }
  }
}

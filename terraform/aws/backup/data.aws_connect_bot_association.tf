data "aws_connect_bot_association" "pike" {
  instance_id = "aaaaaaaa-bbbb-cccc-dddd-111111111111"
  lex_bot {
    name = "pike"
  }
}

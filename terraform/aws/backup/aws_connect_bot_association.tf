resource "aws_connect_bot_association" "pike" {

  instance_id = aws_connect_instance.pike.id
  lex_bot {
    lex_region = "eu-west-2"
    name       = "bookish"
  }
}

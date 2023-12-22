resource "aws_connect_routing_profile" "pike" {

  instance_id = aws_connect_instance.pike.id
  name        = "example"


  default_outbound_queue_id = aws_connect_queue.pike.id
  description               = "example description"

  media_concurrencies {
    channel     = "VOICE"
    concurrency = 1
  }

  queue_configs {
    channel  = "VOICE"
    delay    = 2
    priority = 1
    queue_id = "12345678-1234-1234-1234-123456789012"
  }


  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

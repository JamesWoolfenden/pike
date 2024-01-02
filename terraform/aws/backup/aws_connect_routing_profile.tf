resource "aws_connect_routing_profile" "pike" {

  instance_id = aws_connect_instance.pike.id
  name        = "example"


  default_outbound_queue_id = aws_connect_queue.pike.queue_id
  description               = "example description"

  media_concurrencies {
    channel     = "VOICE"
    concurrency = 1
  }

  queue_configs {
    channel  = "VOICE"
    delay    = 2
    priority = 1
    queue_id = aws_connect_queue.pike.queue_id
  }


  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

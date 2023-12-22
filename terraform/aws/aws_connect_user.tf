resource "aws_connect_user" "pike" {
  instance_id        = aws_connect_instance.pike.id
  name               = "example"
  password           = "Password123"
  routing_profile_id = aws_connect_routing_profile.pike.routing_profile_id

  security_profile_ids = [
    aws_connect_security_profile.pike.security_profile_id
  ]

  identity_info {
    first_name = "example"
    last_name  = "example2"
  }

  phone_config {
    after_contact_work_time_limit = 0
    phone_type                    = "SOFT_PHONE"
  }

  tags = {
    pike    = "permissions"
    another = "tag"
  }
}

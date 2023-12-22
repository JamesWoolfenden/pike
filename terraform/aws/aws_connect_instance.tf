resource "aws_connect_instance" "pike" {
  identity_management_type  = "CONNECT_MANAGED"
  inbound_calls_enabled     = false
  instance_alias            = "pike2"
  outbound_calls_enabled    = false
  contact_flow_logs_enabled = true
}

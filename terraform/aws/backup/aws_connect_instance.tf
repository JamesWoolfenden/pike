resource "aws_connect_instance" "pike" {
  identity_management_type       = "CONNECT_MANAGED"
  inbound_calls_enabled          = true
  instance_alias                 = "pike4"
  outbound_calls_enabled         = true
  contact_flow_logs_enabled      = true
  multi_party_conference_enabled = true
}

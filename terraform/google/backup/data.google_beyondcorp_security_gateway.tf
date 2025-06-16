data "google_beyondcorp_security_gateway" "pike" {
  security_gateway_id = "pike"

}

output "google_beyondcorp_security_gateway" {
  value = data.google_beyondcorp_security_gateway.pike
}

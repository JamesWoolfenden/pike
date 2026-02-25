resource "google_compute_network" "pike_vpn_network" {
  name                    = "pike-vpn-network"
  auto_create_subnetworks = false
}

resource "google_compute_vpn_gateway" "pike" {
  name    = "pike-vpn-gateway"
  network = google_compute_network.pike_vpn_network.id
}

resource "google_compute_address" "pike_vpn_ip" {
  name = "pike-vpn-static-ip"
}

resource "google_compute_forwarding_rule" "pike_vpn_esp" {
  name        = "pike-vpn-esp"
  ip_protocol = "ESP"
  ip_address  = google_compute_address.pike_vpn_ip.address
  target      = google_compute_vpn_gateway.pike.id
}

resource "google_compute_forwarding_rule" "pike_vpn_udp500" {
  name        = "pike-vpn-udp500"
  ip_protocol = "UDP"
  port_range  = "500"
  ip_address  = google_compute_address.pike_vpn_ip.address
  target      = google_compute_vpn_gateway.pike.id
}

resource "google_compute_forwarding_rule" "pike_vpn_udp4500" {
  name        = "pike-vpn-udp4500"
  ip_protocol = "UDP"
  port_range  = "4500"
  ip_address  = google_compute_address.pike_vpn_ip.address
  target      = google_compute_vpn_gateway.pike.id
}

resource "google_compute_vpn_tunnel" "pike" {
  name          = "pike-vpn-tunnel"
  peer_ip       = "8.8.8.8"
  shared_secret = "a-secret-message"

  target_vpn_gateway = google_compute_vpn_gateway.pike.id

  depends_on = [
    google_compute_forwarding_rule.pike_vpn_esp,
    google_compute_forwarding_rule.pike_vpn_udp500,
    google_compute_forwarding_rule.pike_vpn_udp4500,
  ]
}

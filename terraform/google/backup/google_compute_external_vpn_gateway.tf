resource "google_compute_external_vpn_gateway" "pike" {
  name            = "pike-external-vpn-gw"
  redundancy_type = "SINGLE_IP_INTERNALLY_REDUNDANT"
  description     = "Pike test external VPN gateway"

  interface {
    id         = 0
    ip_address = "8.8.8.8"
  }
}

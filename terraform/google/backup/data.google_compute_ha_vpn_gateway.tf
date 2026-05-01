data "google_compute_ha_vpn_gateway" "pike" {
  name = "pike"
}

output "google_compute_ha_vpn_gateway" {
  value = data.google_compute_ha_vpn_gateway.pike
}

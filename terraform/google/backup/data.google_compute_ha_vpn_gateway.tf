data "google_compute_ha_vpn_gateway" "pike" {
  name = "pike"
}

output "gateway" {
  value = data.google_compute_ha_vpn_gateway.pike
}

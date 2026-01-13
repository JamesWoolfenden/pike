data "google_compute_routers" "pike" {
}

output "google_compute_routers" {
  value = data.google_compute_routers.pike
}

data "google_compute_region_instance_group_manager" "pike" {
  name = "pike"
}

output "google_compute_region_instance_group_manager" {
  value = data.google_compute_region_instance_group_manager.pike
}

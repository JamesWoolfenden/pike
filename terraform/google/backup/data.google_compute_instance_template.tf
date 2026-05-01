data "google_compute_instance_template" "pike" {
  name = "generic-tpl-20200107"
}

output "google_compute_instance_template" {
  value = data.google_compute_instance_template.pike
}

data "google_compute_instance_template" "pike" {
  name = "generic-tpl-20200107"
}

output "template" {
  value = data.google_compute_instance_template.pike
}

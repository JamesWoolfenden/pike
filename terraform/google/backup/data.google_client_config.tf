data "google_client_config" "pike" {}

output "google_client_config" {
  sensitive = true
  value     = data.google_client_config.pike
}

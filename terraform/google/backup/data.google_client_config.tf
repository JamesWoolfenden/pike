data "google_client_config" "pike" {}

output "config" {
  sensitive = true
  value     = data.google_client_config.pike
}

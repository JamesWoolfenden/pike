data "google_secret_manager_regional_secret" "pike" {
  location  = "us-central1"
  secret_id = "secretname"
}

output "google_secret_manager_regional_secret" {
  value = data.google_secret_manager_regional_secret.pike
}

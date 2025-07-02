data "google_organizations" "pike" {
}

output "google_organizations" {
  value = data.google_organizations.pike
}

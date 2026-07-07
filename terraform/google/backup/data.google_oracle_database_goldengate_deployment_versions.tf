data "google_oracle_database_goldengate_deployment_versions" "pike" {
  location = "us-east4"
}

output "google_oracle_database_goldengate_deployment_versions" {
  value = data.google_oracle_database_goldengate_deployment_versions.pike
}

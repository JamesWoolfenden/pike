data "google_oracle_database_goldengate_deployment_environments" "pike" {
  location = "us-east4"
}

output "google_oracle_database_goldengate_deployment_environments" {
  value = data.google_oracle_database_goldengate_deployment_environments.pike
}

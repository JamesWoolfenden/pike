data "google_oracle_database_goldengate_deployment_types" "pike" {
  location = "us-east4"
}

output "google_oracle_database_goldengate_deployment_types" {
  value = data.google_oracle_database_goldengate_deployment_types.pike
}

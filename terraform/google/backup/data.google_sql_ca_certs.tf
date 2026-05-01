data "google_sql_ca_certs" "pike" {
  instance = "examplea"
}

output "google_sql_ca_certs" {
  value = data.google_sql_ca_certs.pike
}

data "google_sql_ca_certs" "pike" {
  instance = "examplea"
}

output "certs" {
  value = data.google_sql_ca_certs.pike
}

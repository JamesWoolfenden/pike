data "google_vpc_access_connector" "pike" {
  name = "pike"
}

output "google_vpc_access_connector" {
  value = data.google_vpc_access_connector.pike
}

data "google_vpc_access_connector" "pike" {
  name = "pike"
}

output "connect" {
  value = data.google_vpc_access_connector.pike
}

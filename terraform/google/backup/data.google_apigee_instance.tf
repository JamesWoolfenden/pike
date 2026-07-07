data "google_apigee_instance" "pike" {
  name   = "pike"
  org_id = "organizations/pike-477416"
}

output "google_apigee_instance" {
  value = data.google_apigee_instance.pike
}

data "google_vmwareengine_external_access_rule" "pike" {
  name   = "my-external-access-rule"
  parent = "projects/pike-gcp/locations/us-central1/privateClouds/my-cloud"
}

output "access_rule" {
  value = data.google_vmwareengine_external_access_rule.pike
}

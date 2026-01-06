resource "google_access_context_manager_service_perimeter_ingress_policy" "pike" {
  perimeter = google_access_context_manager_service_perimeter.storage-perimeter.name
  title     = "ingress policy title"
  ingress_from {
    identity_type = "ANY_IDENTITY"
    sources {
      access_level = "*"
    }
  }
  ingress_to {
    resources = ["*"]
    operations {
      service_name = "bigquery.googleapis.com"
      method_selectors {
        method = "*"
      }
    }
  }
  lifecycle {
    create_before_destroy = true
  }
}

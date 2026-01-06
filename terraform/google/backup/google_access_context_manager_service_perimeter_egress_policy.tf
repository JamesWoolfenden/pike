resource "google_access_context_manager_service_perimeter_egress_policy" "egress_policy" {
  perimeter = google_access_context_manager_service_perimeter.storage-perimeter.name
  title     = "egress policy title"
  egress_from {
    identity_type = "ANY_IDENTITY"
  }
  egress_to {
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

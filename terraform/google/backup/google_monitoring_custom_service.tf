resource "google_monitoring_custom_service" "pike" {
  service_id   = "custom-srv"
  display_name = "My Custom Service custom-srv"

  telemetry {
    resource_name = "//product.googleapis.com/foo/foo/services/test"
  }

  user_labels = {
    my_key       = "my_value"
    my_other_key = "my_other_value"
  }
}

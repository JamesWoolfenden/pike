resource "google_data_loss_prevention_discovery_config" "basic" {
  parent   = "projects/my-project-name/locations/us"
  location = "us"
  status   = "RUNNING"

  targets {
    big_query_target {
      filter {
        other_tables {}
      }
    }
  }
  inspect_templates = ["projects/12323132123/inspectTemplates/display_name"]
}

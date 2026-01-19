resource "google_healthcare_workspace" "pike" {
  settings {
    data_project_ids = ["123456"]
  }

  name    = "pike"
  dataset = "data-1234546"
}

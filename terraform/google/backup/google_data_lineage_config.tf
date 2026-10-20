resource "google_data_lineage_config" "pike" {
  location = "us-central1"
  parent   = "projects/pike-477416"
  ingestion {
    rule {
      integration_selector {
        integration = "DATAPROC"
      }
      lineage_enablement {
        enabled = true
      }
    }
  }
}

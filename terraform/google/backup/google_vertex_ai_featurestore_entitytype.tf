resource "google_vertex_ai_featurestore_entitytype" "pike" {
  provider = google-beta
  name     = "terraform"
  labels = {
    foo = "bar"
  }
  description  = "test description"
  featurestore = google_vertex_ai_featurestore.pike.name

  monitoring_config {
    snapshot_analysis {
      disabled                 = false
      monitoring_interval_days = 1
      staleness_days           = 21
    }
    numerical_threshold_config {
      value = 0.8
    }
    categorical_threshold_config {
      value = 10.0
    }
    import_features_analysis {
      state                      = "ENABLED"
      anomaly_detection_baseline = "PREVIOUS_IMPORT_FEATURES_STATS"
    }
  }
}

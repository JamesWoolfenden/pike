resource "google_vertex_ai_feature_online_store_featureview" "pike" {

  feature_online_store = google_vertex_ai_feature_online_store.pike.id
  region               = "us-central1"
  sync_config {
    cron = "0 0 * * *"
  }
  big_query_source {
    uri               = "bq://${google_bigquery_table.tf-test-table.project}.${google_bigquery_table.tf-test-table.dataset_id}.${google_bigquery_table.tf-test-table.table_id}"
    entity_id_columns = ["test_entity_column"]

  }
}

resource "google_bigquery_dataset" "tf-test-dataset" {

  dataset_id    = "example_feature_view"
  friendly_name = "test"
  description   = "This is a test description"
  location      = "US"
}

resource "google_bigquery_table" "tf-test-table" {
  deletion_protection = false

  dataset_id = google_bigquery_dataset.tf-test-dataset.dataset_id
  table_id   = "example_feature_view"
  schema     = <<EOF
  [
  {
    "name": "entity_id",
    "mode": "NULLABLE",
    "type": "STRING",
    "description": "Test default entity_id"
  },
    {
    "name": "test_entity_column",
    "mode": "NULLABLE",
    "type": "STRING",
    "description": "test secondary entity column"
  },
  {
    "name": "feature_timestamp",
    "mode": "NULLABLE",
    "type": "TIMESTAMP",
    "description": "Default timestamp value"
  }
]
EOF
}

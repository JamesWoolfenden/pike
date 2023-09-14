resource "google_bigquery_table" "pike" {
  #  dataset_id = google_bigquery_dataset.pike.dataset_id
  dataset_id = "foo"
  table_id   = "bar"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env  = "default"
    pike = "permissions"
  }

  schema              = <<EOF
[
  {
    "name": "permalink",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "The Permalink"
  },
  {
    "name": "state",
    "type": "STRING",
    "mode": "NULLABLE",
    "description": "State where the head office is located"
  }
]
EOF
  deletion_protection = false
}

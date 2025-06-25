resource "google_bigquery_dataset" "pike" {
  dataset_id                  = "foo"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env  = "default"
    pike = "permissions"
  }
}

resource "google_bigquery_table" "pike" {
  dataset_id = google_bigquery_dataset.pike.dataset_id
  table_id   = "bar"

  time_partitioning {
    type = "DAY"
  }

  labels = {
    env = "default"
  }

  schema = <<EOF
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

resource "google_bigtable_table" "table" {
  name          = "tf-table"
  instance_name = google_bigtable_instance.instance.name
  split_keys    = ["a", "b", "c", "d"]

  lifecycle {
    prevent_destroy = false
  }

  column_family {
    family = "family-first"
  }

  column_family {
    family = "family-second"
  }

  // change_stream_retention = "0"

}

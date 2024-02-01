resource "google_bigquery_table_iam_policy" "pike" {
  dataset_id  = google_bigquery_dataset.default.id
  policy_data = data.google_iam_policy.admin.policy_data
  table_id    = google_bigquery_table.default.id
}

data "google_iam_policy" "admin" {
  binding {
    role = "roles/bigquery.dataOwner"
    members = [
      "user:crwoolfenden@gmail.com",
    ]
  }
}

resource "google_bigquery_dataset" "default" {
  dataset_id                  = "foo"
  friendly_name               = "test"
  description                 = "This is a test description"
  location                    = "EU"
  default_table_expiration_ms = 3600000

  labels = {
    env = "default"
  }
}

resource "google_bigquery_table" "default" {
  dataset_id = google_bigquery_dataset.default.dataset_id
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

}

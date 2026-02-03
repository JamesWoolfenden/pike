resource "google_storage_bucket" "basic_bucket" {
  name                        = "bucket"
  location                    = "us-west1"
  uniform_bucket_level_access = true
  lifecycle {
    ignore_changes = [
      labels
    ]
  }

  project = "my-project-name"
}

resource "google_dataplex_lake" "basic_lake" {
  name     = "lake"
  location = "us-west1"
  project  = "my-project-name"
}


resource "google_dataplex_zone" "basic_zone" {
  name     = "zone"
  location = "us-west1"
  lake     = google_dataplex_lake.basic_lake.name
  type     = "RAW"

  discovery_spec {
    enabled = false
  }


  resource_spec {
    location_type = "SINGLE_REGION"
  }

  project = "my-project-name"
}


resource "google_dataplex_asset" "primary" {
  name     = "asset"
  location = "us-west1"

  lake          = google_dataplex_lake.basic_lake.name
  dataplex_zone = google_dataplex_zone.basic_zone.name

  discovery_spec {
    enabled = false
  }

  resource_spec {
    name = "projects/my-project-name/buckets/bucket"
    type = "STORAGE_BUCKET"
  }

  labels = {
    env      = "foo"
    my-asset = "exists"
  }


  project = "my-project-name"
  depends_on = [
    google_storage_bucket.basic_bucket
  ]
}

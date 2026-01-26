resource "google_app_engine_service_split_traffic" "pike" {
  service = google_app_engine_standard_app_version.liveapp_v2.service

  migrate_traffic = false
  split {
    shard_by = "IP"
    allocations = {
      (google_app_engine_standard_app_version.liveapp_v1.version_id) = 0.75
      (google_app_engine_standard_app_version.liveapp_v2.version_id) = 0.25
    }
  }
}

resource "google_app_engine_standard_app_version" "liveapp_v2" {
  version_id      = "v2"
  service         = "default"
  noop_on_destroy = true

  runtime = "nodejs20"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}


resource "google_app_engine_standard_app_version" "liveapp_v1" {
  version_id      = "v1"
  service         = "default"
  noop_on_destroy = true

  runtime = "nodejs20"
  entrypoint {
    shell = "node ./app.js"
  }
  deployment {
    zip {
      source_url = "https://storage.googleapis.com/${google_storage_bucket.bucket.name}/${google_storage_bucket_object.object.name}"
    }
  }
  env_variables = {
    port = "8080"
  }
}

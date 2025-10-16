resource "google_vertex_ai_index_endpoint_deployed_index" "basic_deployed_index" {
  deployed_index_id     = "deployed_index_id"
  display_name          = "vertex-deployed-index"
  region                = "us-central1"
  index                 = google_vertex_ai_index.index.id
  index_endpoint        = google_vertex_ai_index_endpoint.vertex_index_endpoint_deployed.id
  enable_access_logging = false
  reserved_ip_ranges    = ["vertex-ai-range"]

  deployed_index_auth_config {

    auth_provider {
      audiences       = ["123456-my-app"]
      allowed_issuers = ["${google_service_account.sa.email}"]
    }
  }
}

resource "google_vertex_ai_index" "index" {
  region              = "us-central1"
  display_name        = "test-index"
  description         = "index for test"
  index_update_method = "BATCH_UPDATE"
  labels = {
    foo = "bar"
  }

  metadata {
    contents_delta_uri = "gs://${google_storage_bucket.bucket.name}/contents"

    config {
      dimensions                  = 2
      approximate_neighbors_count = 150
      shard_size                  = "SHARD_SIZE_SMALL"
      distance_measure_type       = "DOT_PRODUCT_DISTANCE"

      algorithm_config {

        tree_ah_config {
          leaf_node_embedding_count    = 500
          leaf_nodes_to_search_percent = 7
        }
      }
    }
  }
}

resource "google_vertex_ai_index_endpoint" "vertex_index_endpoint_deployed" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  region       = "us-central1"
  network      = "projects/${data.google_project.project.number}/global/networks/${data.google_compute_network.vertex_network.name}"
  labels = {
    label-one = "value-one"
  }
}

resource "google_service_account" "sa" {
  account_id = "vertex-sa"
}

resource "google_storage_bucket" "bucket" {
  name                        = "bucket-name-jgw-test"
  location                    = "us-central1"
  uniform_bucket_level_access = true
}

# The sample data comes from the following link:
# https://cloud.google.com/vertex-ai/docs/matching-engine/filtering#specify-namespaces-tokens
resource "google_storage_bucket_object" "data" {
  name    = "contents/data.json"
  bucket  = google_storage_bucket.bucket.name
  content = <<EOF
{"id": "42", "embedding": [0.5, 1.0], "restricts": [{"namespace": "class", "allow": ["cat", "pet"]},{"namespace": "category", "allow": ["feline"]}]}
{"id": "43", "embedding": [0.6, 1.0], "restricts": [{"namespace": "class", "allow": ["dog", "pet"]},{"namespace": "category", "allow": ["canine"]}]}
EOF
}

data "google_compute_network" "vertex_network" {
  name = "default"
}

data "google_project" "project" {}

resource "google_vertex_ai_index_endpoint" "index_endpoint" {
  display_name = "sample-endpoint"
  description  = "A sample vertex endpoint"
  region       = "us-central1"
  labels = {
    label-one = "value-one"
  }
  network = "projects/${data.google_project.project.number}/global/networks/${google_compute_network.vertex_network.name}"
  depends_on = [
    google_service_networking_connection.vertex_vpc_connection
  ]
}

resource "google_service_networking_connection" "vertex_vpc_connection" {
  network                 = google_compute_network.vertex_network.id
  service                 = "servicenetworking.googleapis.com"
  reserved_peering_ranges = [google_compute_global_address.vertex_range.name]
}

resource "google_compute_global_address" "vertex_range" {
  name          = "address-name"
  purpose       = "VPC_PEERING"
  address_type  = "INTERNAL"
  prefix_length = 24
  network       = google_compute_network.vertex_network.id
}

resource "google_compute_network" "vertex_network" {
  name = "network-name"
}

resource "google_kms_crypto_key_iam_member" "crypto_key" {
  crypto_key_id = "projects/pike-gcp/locations/us-central1/keyRings/cluster-1vpUU/cryptoKeys/cluster-1vpUU"
  role          = "roles/cloudkms.cryptoKeyEncrypterDecrypter"
  member        = "serviceAccount:service-${data.google_project.project.number}@gcp-sa-aiplatform.iam.gserviceaccount.com"
}

data "google_project" "project" {}

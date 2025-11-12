provider "google" {
  project     = "pike-477416"
  region      = "europe-west2"
  credentials = var.basic_gcp_credentials
}

provider "google-beta" {
  project = "pike-477416"
  # region  = "europe-west2"
  region      = "us-central1"
  credentials = var.basic_gcp_credentials
}

provider "google" {
  alias       = "central"
  project     = "pike-477416"
  region      = "us-central1"
  credentials = var.basic_gcp_credentials
}

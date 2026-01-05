provider "google" {
  project     = "pike-477416"
  region      = "europe-west2"
  credentials = local.gcp_credentials_path
  scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/cloud-identity.policies.readonly",
  ]
}

provider "google-beta" {
  project = "pike-477416"
  # region  = "europe-west2"
  region      = "us-central1"
  credentials = local.gcp_credentials_path
  scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/cloud-identity.policies.readonly",
  ]
}

provider "google" {
  alias       = "central"
  project     = "pike-477416"
  region      = "us-central1"
  credentials = local.gcp_credentials_path
  scopes = [
    "https://www.googleapis.com/auth/cloud-platform",
    "https://www.googleapis.com/auth/cloud-identity.policies.readonly",
  ]
}

locals {
  gcp_credentials_path = fileexists("/Users/jwoolfenden/pike-gcp-super.json") ? "/Users/jwoolfenden/pike-gcp-super.json" : "c:/Users/jim_w/pike-super.json"
}

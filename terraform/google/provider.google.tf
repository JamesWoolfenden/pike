provider "google" {
  project     = "pike-gcp"
  region      = "europe-west2"
  credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  //credentials = "/Users/jwoolfenden/pike-service.json"
}

provider "google-beta" {
  project     = "pike-gcp"
  region      = "europe-west2"
  credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  //credentials = "/Users/jwoolfenden/pike-service.json"
}

provider "google" {
  alias       = "central"
  project     = "pike-gcp"
  region      = "us-central1"
  credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  //credentials = "/Users/jwoolfenden/pike-service.json"
}

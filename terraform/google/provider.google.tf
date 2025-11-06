provider "google" {
  project = "pike-477416"
  region  = "europe-west2"

  # credentials = "C:\\Users\\jim_w\\pike-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

provider "google-beta" {
  project = "pike-477416"
  # region  = "europe-west2"
  region = "us-central1"
  # credentials = "C:\\Users\\jim_w\\pike-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

provider "google" {
  alias   = "central"
  project = "pike-477416"
  region  = "us-central1"
  # credentials = "C:\\Users\\jim_w\\pike-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

provider "google" {
  project = "pike-412922"
  region  = "europe-west2"
  //credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

provider "google-beta" {
  project = "pike-412922"
  region  = "europe-west2"
  //credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

provider "google" {
  alias   = "central"
  project = "pike-412922"
  region  = "us-central1"
  // credentials = "C:\\Users\\jim_w\\pike-gcp-basic.json"
  credentials = "/Users/jwoolfenden/pike-gcp-basic.json"
}

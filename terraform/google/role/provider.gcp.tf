provider "google" {
  project = "pike-477416"
  # from GOOGLE_APPLICATION_CREDENTIALS
  # credentials = "c:/Users/jim_w/pike-super.json"
  credentials = local.gcp_credentials_path
  # credentials = "/Users/jwoolfenden/pike-gcp-super.json"
}


locals {
  gcp_credentials_path = fileexists("/Users/jwoolfenden/pike-gcp-super.json") ? "/Users/jwoolfenden/pike-gcp-super.json" : "c:/Users/jim_w/pike-super.json"
}

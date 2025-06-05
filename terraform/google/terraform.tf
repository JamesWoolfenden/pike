terraform {
  backend "gcs" {
    credentials = "/Users/jwoolfenden/pike-gcp-super.json"
    bucket      = "terraform-pike-bucket-tfstate"
    prefix      = "trial/state"
  }
}

terraform {
  backend "gcs" {
    credentials = "/Users/jwoolfenden/pike-gcp-super.json"
    # credentials = "c:/Users/jim_w/pike-412922-3277f720572e.json"
    bucket = "terraform-pike-bucket-tfstate"
    prefix = "iam/state"
  }
}

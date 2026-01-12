terraform {
  backend "gcs" {

    #  credentials = "c:/Users/jim_w/pike-super.json"
    # credentials = "/Users/jwoolfenden/pike-gcp-super.json"
    bucket = "terraform-pike-bucket-tfstate"
    prefix = "iam/state"
  }
}

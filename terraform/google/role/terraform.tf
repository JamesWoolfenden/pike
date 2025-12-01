terraform {
  backend "gcs" {
    credentials = "c:/Users/jim_w/pike-super.json"
    # credentials = "c:/Users/jim_w/pike-477416-3277f720572e.json"
    bucket = "terraform-pike-bucket-tfstate"
    prefix = "iam/state"
  }
}

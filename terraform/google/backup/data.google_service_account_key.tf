
resource "google_service_account" "myaccount" {
  account_id = "dev-foo-account"
}

resource "google_service_account_key" "mykey" {
  service_account_id = google_service_account.myaccount.name
}

data "google_service_account_key" "mykey" {
  name            = google_service_account_key.mykey.name
  public_key_type = "TYPE_X509_PEM_FILE"
}

resource "google_dataplex_data_product" "pike" {
  data_product_id = "pike"
  display_name    = "pike"
  location        = "us-central1"
  owner_emails    = ["james.woolfenden@gmail.com"]
}

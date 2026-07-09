resource "google_dataplex_data_product_iam_member" "pike" {
  location        = google_dataplex_data_product.pike.location
  data_product_id = google_dataplex_data_product.pike.data_product_id

  role   = "roles/viewer"
  member = "user:james.woolfenden@gmail.com"
}

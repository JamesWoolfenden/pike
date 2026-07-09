data "google_iam_policy" "google_dataplex_data_product" {
  binding {
    role = "roles/viewer"
    members = [
      "user:james.woolfenden@gmail.com",
    ]
  }
}

resource "google_dataplex_data_product_iam_policy" "pike" {
  location        = google_dataplex_data_product.pike.location
  data_product_id = google_dataplex_data_product.pike.data_product_id

  policy_data = data.google_iam_policy.google_dataplex_data_product.policy_data
}

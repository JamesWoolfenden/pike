data "google_dataplex_data_product_iam_policy" "pike" {
  location        = "us-central1"
  data_product_id = "pike"
}

output "google_dataplex_data_product_iam_policy" {
  value = data.google_dataplex_data_product_iam_policy.pike
}

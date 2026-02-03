resource "google_dataplex_data_product" "pike" {
  project         = "my-project-name"
  location        = "us-central1"
  data_product_id = "data-product-basic"
  display_name    = "terraform data product"

  owner_emails = ["terraform-test@google.com"]

  access_groups {
    id           = "analyst"
    group_id     = "analyst"
    display_name = "Data Analyst"
    principal {
      google_group = "tf-test-analysts-${random_id.suffix.id}@example.com"
    }
  }

  provider = google-beta
}

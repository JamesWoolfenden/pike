resource "google_data_catalog_taxonomy" "pike" {

  display_name           = "my_taxonomy"
  description            = "A collection of policy tags"
  activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
}

resource "google_bigquery_datapolicy_data_policy" "pike" {
  location         = "europe-west2"
  data_policy_id   = "data_policy"
  policy_tag       = google_data_catalog_policy_tag.pike.name
  data_policy_type = "COLUMN_LEVEL_SECURITY_POLICY"
}
#
# resource "google_data_catalog_policy_tag" "policy_tag" {
#   taxonomy     = google_data_catalog_taxonomy.taxonomy.id
#   display_name = "Low security"
#   description  = "A policy tag normally associated with low security items"
# }
#
# resource "google_data_catalog_taxonomy" "taxonomy" {
#   region                 = "us-central1"
#   display_name           = "taxonomy"
#   description            = "A collection of policy tags"
#   activated_policy_types = ["FINE_GRAINED_ACCESS_CONTROL"]
# }

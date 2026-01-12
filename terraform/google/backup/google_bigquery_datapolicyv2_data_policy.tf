resource "google_bigquery_datapolicyv2_data_policy" "pike" {
  location         = "us-central1"
  data_policy_type = "RAW_DATA_ACCESS_POLICY"
  data_policy_id   = "basic_data_policy"
}

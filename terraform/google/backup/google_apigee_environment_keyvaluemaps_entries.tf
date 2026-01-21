resource "google_apigee_environment_keyvaluemaps_entries" "pike" {
  name               = "pike"
  env_keyvaluemap_id = google_apigee_environment_keyvaluemaps.pike.env_id


  value = "testValue"
}

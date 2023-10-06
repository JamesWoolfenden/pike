data "google_gke_hub_scope_iam_policy" "pike" {
  project  = "pike-gcp"
  scope_id = "fred"
}

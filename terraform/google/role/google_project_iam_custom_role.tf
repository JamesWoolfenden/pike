resource "google_project_iam_custom_role" "pike" {
  project     = "pike-412922"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "storage.buckets.create",
    "storage.buckets.delete",
    "storage.buckets.get",
    "storage.buckets.update",

    //google_compute_region_health_check


    //google_compute_https_health_check


    //google_compute_http_health_check


    //google_compute_health_check


    //google_compute_backend_service


    //google_compute_region_backend_service


    //google_compute_backend_bucket


    //google_compute_region_url_map


    //google_compute_url_map


    //google_compute_target_http_proxy


    //google_compute_region_target_http_proxy











  ]
}

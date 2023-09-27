resource "google_compute_region_ssl_certificate" "pike" {
  //region      = "us-central1"
  description = "selfsigned"
  name_prefix = "pike-"
  private_key = file("./key.pem")
  certificate = file("./certificate.pem")

  lifecycle {
    create_before_destroy = true
  }
}

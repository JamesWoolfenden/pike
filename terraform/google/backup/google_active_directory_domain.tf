resource "google_active_directory_domain" "ad-domain" {
  domain_name         = "tfgen.org.com"
  locations           = ["us-central1"]
  reserved_ip_range   = "192.168.255.0/24"
  deletion_protection = false
}

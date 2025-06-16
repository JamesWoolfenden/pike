data "google_compute_security_policy" "pike" {
  name = "pike"

}

output "google_compute_security_policy" {
  value = data.google_compute_security_policy.pike
}

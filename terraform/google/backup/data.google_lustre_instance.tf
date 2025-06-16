data "google_lustre_instance" "pike" {
  instance_id = "pike"
  zone        = "us-central1-a"
}

output "google_lustre_instance" {
  value = data.google_lustre_instance.pike
}

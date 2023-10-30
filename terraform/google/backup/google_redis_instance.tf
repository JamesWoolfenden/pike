resource "google_redis_instance" "pike" {
  name           = "pike"
  memory_size_gb = 1
  labels = {
    pike = "permissions"
  }
}

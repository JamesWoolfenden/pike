data "google_memcache_instance" "pike" {
  name = "pike"
}

output "google_memcache_instance" {
  value = data.google_memcache_instance.pike
}

data "google_redis_cluster" "pike" {
}

output "google_redis_cluster" {
  value = data.google_redis_cluster.pike
}

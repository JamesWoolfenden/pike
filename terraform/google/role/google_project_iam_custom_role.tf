resource "google_project_iam_custom_role" "pike" {
  project     = "pike-gcp"
  role_id     = "terraform_pike"
  title       = "pike terraform user"
  description = "A user with least privileges"
  permissions = [
    "redis.instances.get",
    "redis.instances.create",
    "redis.operations.get",
    "redis.instances.update",
    "redis.instances.delete"
  ]
}
